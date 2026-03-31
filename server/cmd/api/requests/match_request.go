package requests

type CreateMatchRequest struct {
	LeagueId string `json:"league_id" binding:"required"`
	Date     string `json:"date" binding:"required"` // RFC3339 format e.g. "2024-01-15T15:00:00Z"
	SeasonId string `json:"season_id"`               // optional: associate match with a season
}
