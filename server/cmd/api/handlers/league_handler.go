package handlers

import (
	"log"
	"net/http"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LeagueHandler struct {
	leagueService *services.LeagueService
	playerService *services.PlayerService
}

func NewLeagueHandler(leagueService *services.LeagueService, playerService *services.PlayerService) *LeagueHandler {
	return &LeagueHandler{leagueService: leagueService, playerService: playerService}
}

func (h *LeagueHandler) CreateLeague(c *gin.Context) {
	type CreateLeagueRequest struct {
		OwnerId uuid.UUID `json:"owner_id" binding:"required"`
		Name    string    `json:"name" binding:"required"`
	}

	var request CreateLeagueRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLeague, err := h.leagueService.CreateLeague(&models.League{
		OwnerId: request.OwnerId,
		Name:    request.Name,
	})

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdLeague)
}

func (h *LeagueHandler) GetLeagueBySlug(c *gin.Context) {
	leagueSlug := c.Param("leagueSlug")
	league, err := h.leagueService.GetLeagueBy(constants.SLUG, leagueSlug)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "League not found", "detail_message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, league)
}

func (h *LeagueHandler) UpdateLeague(c *gin.Context) {
	var league models.League
	if err := c.ShouldBindJSON(&league); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLeague, err := h.leagueService.UpdateLeague(&league)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLeague)
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
	type AddPlayerRequest struct {
		Email string `json:"email" binding:"required"`
	}

	var request AddPlayerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leagueSlug := c.Param("leagueSlug")
	league, err := h.leagueService.GetLeagueBy(constants.SLUG, leagueSlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
		return
	}

	player, err := h.playerService.GetPlayerBy(constants.EMAIL, request.Email)
	if err != nil || player == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	league.Members = append(league.Members, *player)

	updatedLeague, err := h.leagueService.UpdateLeague(league)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update league"})
		return
	}

	c.JSON(http.StatusOK, updatedLeague)
}
