package repository

import (
	"fmt"
	"log"

	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/models/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) *SeasonRepository {
	return &SeasonRepository{db: db}
}

func (r *SeasonRepository) CreateSeason(season *models.Season, leagueId uuid.UUID) (*models.Season, error) {
	if err := r.db.Create(season).Error; err != nil {
		log.Printf("Error creating season: %v", err)
		return nil, err
	}

	// Associate season with league via join table
	var league models.League
	if err := r.db.First(&league, leagueId).Error; err != nil {
		log.Printf("Error finding league %v: %v", leagueId, err)
		return nil, err
	}
	if err := r.db.Model(&league).Association("Seasons").Append(season); err != nil {
		log.Printf("Error associating season with league: %v", err)
		return nil, err
	}

	return season, nil
}

func (r *SeasonRepository) GetSeasonsByLeagueId(leagueId uuid.UUID) ([]models.Season, error) {
	var league models.League
	if err := r.db.Preload("Seasons").First(&league, leagueId).Error; err != nil {
		log.Printf("Error getting seasons for league %v: %v", leagueId, err)
		return nil, err
	}
	return league.Seasons, nil
}

func (r *SeasonRepository) GetSeasonById(id uuid.UUID) (*models.Season, error) {
	var season models.Season
	if err := r.db.First(&season, id).Error; err != nil {
		log.Printf("Error getting season %v: %v", id, err)
		return nil, err
	}
	return &season, nil
}

// GetSeasonStats returns all players that participated in a season with their aggregated stats.
func (r *SeasonRepository) GetSeasonStats(seasonId uuid.UUID) ([]*models.Player, error) {
	type statRow struct {
		PlayerId uuid.UUID
		Type     string
		Count    int
		Matches  int
	}

	var rows []statRow
	err := r.db.Raw(`
		SELECT
			e.player_id,
			e.type,
			COUNT(e.id) AS count,
			COUNT(DISTINCT e.match_id) AS matches
		FROM events e
		INNER JOIN matches m ON m.id = e.match_id
		WHERE m.season_id = ?
		  AND m.deleted_at IS NULL
		  AND e.deleted_at IS NULL
		GROUP BY e.player_id, e.type
	`, seasonId).Scan(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("GetSeasonStats: query error: %w", err)
	}

	// Collect unique player IDs and aggregate stats
	type playerStats struct {
		Goals       int
		Assists     int
		Disarms     int
		Dribbles    int
		YellowCards int
		RedCards    int
		Matches     int
	}
	statsMap := make(map[uuid.UUID]*playerStats)

	for _, row := range rows {
		if _, ok := statsMap[row.PlayerId]; !ok {
			statsMap[row.PlayerId] = &playerStats{}
		}
		s := statsMap[row.PlayerId]
		switch row.Type {
		case enums.Goal:
			s.Goals += row.Count
		case enums.Assist:
			s.Assists += row.Count
		case enums.Disarm:
			s.Disarms += row.Count
		case enums.Dribble:
			s.Dribbles += row.Count
		case enums.YellowCard:
			s.YellowCards += row.Count
		case enums.RedCard:
			s.RedCards += row.Count
		}
		if row.Matches > s.Matches {
			s.Matches = row.Matches
		}
	}

	if len(statsMap) == 0 {
		return []*models.Player{}, nil
	}

	// Collect player IDs
	playerIds := make([]uuid.UUID, 0, len(statsMap))
	for id := range statsMap {
		playerIds = append(playerIds, id)
	}

	// Fetch player records
	var players []*models.Player
	if err := r.db.Where("id IN ?", playerIds).Find(&players).Error; err != nil {
		return nil, fmt.Errorf("GetSeasonStats: fetch players error: %w", err)
	}

	// Populate computed stats
	for _, p := range players {
		if s, ok := statsMap[p.ID]; ok {
			p.Goals = s.Goals
			p.Assists = s.Assists
			p.Disarms = s.Disarms
			p.Dribbles = s.Dribbles
			p.YellowCards = s.YellowCards
			p.RedCards = s.RedCards
			p.Matches = s.Matches
		}
	}

	return players, nil
}

// FinishSeason sets the season status to "finished" and returns the updated season.
func (r *SeasonRepository) FinishSeason(seasonId uuid.UUID) (*models.Season, error) {
	var season models.Season
	if err := r.db.First(&season, seasonId).Error; err != nil {
		return nil, fmt.Errorf("FinishSeason: season not found: %w", err)
	}
	if err := r.db.Model(&season).Update("status", "finished").Error; err != nil {
		return nil, fmt.Errorf("FinishSeason: update error: %w", err)
	}
	season.Status = "finished"
	return &season, nil
}
