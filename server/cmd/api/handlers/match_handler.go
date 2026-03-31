package handlers

import (
	"net/http"
	"time"

	"github.com/bermanbenjamin/futStats/cmd/api/requests"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MatchHandler struct {
	service *services.MatchService
}

func NewMatchHandler(service *services.MatchService) *MatchHandler {
	return &MatchHandler{service: service}
}

func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var req requests.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leagueId, err := uuid.Parse(req.LeagueId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league_id format"})
		return
	}

	date, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use RFC3339 e.g. 2024-01-15T15:00:00Z"})
		return
	}

	newMatch := &models.Match{
		LeagueId: leagueId,
		Date:     date,
	}
	if req.SeasonId != "" {
		seasonId, err := uuid.Parse(req.SeasonId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season_id format"})
			return
		}
		newMatch.SeasonId = &seasonId
	}
	match, err := h.service.CreateMatch(newMatch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": match})
}

func (h *MatchHandler) GetMatchesByLeague(c *gin.Context) {
	leagueSlug := c.Param("leagueSlug")
	matches, err := h.service.GetMatchesByLeagueSlug(leagueSlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": matches})
}

func (h *MatchHandler) GetMatchById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID format"})
		return
	}

	match, err := h.service.GetMatchById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": match})
}

func (h *MatchHandler) DeleteMatch(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID format"})
		return
	}

	if err := h.service.DeleteMatch(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete match"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match deleted successfully"})
}
