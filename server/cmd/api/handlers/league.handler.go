package handlers

import (
	"log"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LeagueHandler struct {
	leagueService *services.LeagueService
}

func NewLeagueHandler(leagueService *services.LeagueService) *LeagueHandler {
	return &LeagueHandler{leagueService: leagueService}
}

func (h *LeagueHandler) CreateLeague(c *gin.Context) {
	type CreateLeagueRequest struct {
		OwnerId uuid.UUID `json:"owner_id" binding:"required"`
		Name    string    `json:"name" binding:"required"`
	}

	var request CreateLeagueRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdLeague, err := h.leagueService.CreateLeague(&models.League{
		OwnerId: request.OwnerId,
		Name:    request.Name,
	})

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, createdLeague)
}

func (h *LeagueHandler) GetLeagueBySlug(c *gin.Context) {
	leagueSlug := c.Param("leagueSlug")
	league, err := h.leagueService.GetLeagueBy(constants.SLUG, leagueSlug)

	if err != nil {
		c.JSON(404, gin.H{"error": "League not found"})
		return
	}

	c.JSON(200, league)

}

func (h *LeagueHandler) UpdateLeague(c *gin.Context) {
	var league models.League
	if err := c.ShouldBindJSON(&league); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedLeague, err := h.leagueService.UpdateLeague(&league)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, updatedLeague)
}

func (h *LeagueHandler) DeleteLeague(c *gin.Context) {
	leagueSlug := c.Param("leagueSlug")
	league, err := h.leagueService.GetLeagueBy(constants.SLUG, leagueSlug)

	if err != nil {
		c.JSON(404, gin.H{"error": "League not found"})
		return
	}

	err = h.leagueService.DeleteLeague(league)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{"message": "League deleted successfully"})
}

func (h *LeagueHandler) AddPlayerToLeague(c *gin.Context) {
	playerIdStr := c.Param("player_id")
	leagueIdStr := c.Param("league_id")

	playerId, err := uuid.Parse(playerIdStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid player ID format"})
		return
	}

	leagueId, err := uuid.Parse(leagueIdStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid league ID format"})
		return
	}

	league, err := h.leagueService.AddPlayerToLeague(playerId, leagueId)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, league)
}
