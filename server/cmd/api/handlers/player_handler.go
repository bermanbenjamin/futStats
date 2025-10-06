package handlers

import (
	"log"
	"net/http"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/logger"
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type PlayerHandler struct {
	service *services.PlayerService
	logger  *logger.Logger
}

// NewPlayerHandler creates a new player handler with proper dependency injection
func NewPlayerHandler(service *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		service: service,
		logger:  logger.GetGlobal().WithComponent("player_handler"),
	}
}

// GetAllPlayers retrieves all players with optional filtering
func (h *PlayerHandler) GetAllPlayers(ctx *gin.Context) {
	filterValue := ctx.GetHeader(constants.QUERY_FILTER)
	filterQuery := constants.QueryFilter(ctx.GetHeader("x-api-query-filter"))

	h.logger.Info("Getting all players",
		zap.String("filter", string(filterQuery)),
		zap.String("value", filterValue))

	players, err := h.service.GetAllPlayers(filterQuery, filterValue)
	if err != nil {
		h.logger.Error("Failed to get players",
			zap.String("filter", string(filterQuery)),
			zap.String("value", filterValue),
			zap.Error(err))

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve players",
			"code":  "PLAYERS_FETCH_FAILED",
		})
		return
	}

	if len(players) == 0 {
		h.logger.Info("No players found",
			zap.String("filter", string(filterQuery)),
			zap.String("value", filterValue))

		ctx.JSON(http.StatusOK, gin.H{
			"data":  []interface{}{},
			"total": 0,
		})
		return
	}

	h.logger.Info("Successfully retrieved players",
		zap.Int("count", len(players)))

	ctx.JSON(http.StatusOK, gin.H{
		"data":  players,
		"total": len(players),
	})
}

func (h *PlayerHandler) GetPlayerBy(ctx *gin.Context) {
	fieldType := constants.QueryFilter(ctx.GetHeader(constants.QUERY_FILTER))
	if fieldType == "" {
		fieldType = constants.ID // Use ID as default
	}

	id := ctx.Param("id") // Assuming ID is passed as a URL parameter

	player, err := h.service.GetPlayerBy(fieldType, id)
	if err != nil {
		log.Printf("Error getting player with %s %v: %v", fieldType, id, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get player"})
		return
	}

	if player == nil {
		log.Printf("Player with %s %v not found", fieldType, id)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Player does not exist"})
		return
	}

	ctx.JSON(http.StatusOK, player)
}

// CreatePlayer creates a new player from request data
func (h *PlayerHandler) CreatePlayer(ctx *gin.Context) {
	var player *models.Player
	if err := ctx.ShouldBindJSON(&player); err != nil {
		h.logger.Error("Failed to bind player data from request",
			zap.Error(err))

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid player data",
			"code":  "INVALID_PLAYER_DATA",
		})
		return
	}

	h.logger.Info("Creating new player",
		zap.String("name", player.Name),
		zap.String("email", player.Email))

	createdPlayer, err := h.service.CreatePlayer(player)
	if err != nil {
		h.logger.Error("Failed to create player",
			zap.String("name", player.Name),
			zap.String("email", player.Email),
			zap.Error(err))

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create player",
			"code":  "PLAYER_CREATION_FAILED",
		})
		return
	}

	h.logger.Info("Player created successfully",
		zap.String("player_id", createdPlayer.ID.String()),
		zap.String("name", createdPlayer.Name))

	ctx.JSON(http.StatusCreated, createdPlayer)
}

func (h *PlayerHandler) UpdatePlayer(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Invalid player ID format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id format"})
		return
	}

	var player models.Player
	if err := ctx.ShouldBindJSON(&player); err != nil {
		log.Printf("Error binding player data: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player data"})
		return
	}

	player.ID = id
	updatedPlayer, err := h.service.UpdatePlayer(&player)
	if err != nil {
		log.Printf("Error updating player with ID %v: %v", id, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player"})
		return
	}

	log.Printf("Player with ID %v updated", id)
	ctx.JSON(http.StatusOK, updatedPlayer)
}

func (h *PlayerHandler) DeletePlayerById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Invalid player ID format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id format"})
		return
	}

	err = h.service.DeletePlayer(id)
	if err != nil {
		log.Printf("Error deleting player with ID %v: %v", id, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}

	log.Printf("Player with ID %v deleted", id)
	ctx.JSON(http.StatusOK, gin.H{"message": "Player deleted"})
}
