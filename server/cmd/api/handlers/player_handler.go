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

type PlayerHandler struct {
	service *services.PlayerService
}

func NewPlayerHandler(service *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}

func (h *PlayerHandler) GetAllPlayers(ctx *gin.Context) {
	filterValue := ctx.GetHeader(constants.QUERY_FILTER)
	filterQuery := constants.QueryFilter(ctx.GetHeader("x-api-query-filter"))

	players, err := h.service.GetAllPlayers(filterQuery, filterValue)

	if err != nil {
		log.Printf("Error getting all players: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve players"})
		return
	}

	if players == nil {
		log.Println("No players found")
		ctx.JSON(http.StatusNoContent, gin.H{"data": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": players, "total": len(players)})
}

func (h *PlayerHandler) GetPlayerBy(ctx *gin.Context) {
	fieldType := constants.QueryFilter(ctx.GetHeader("x-api-query-field"))
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

func (h *PlayerHandler) CreatePlayer(ctx *gin.Context) {
	var player *models.Player
	if err := ctx.ShouldBindJSON(&player); err != nil {
		log.Printf("Error binding player data: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player data"})
		return
	}

	player, err := h.service.CreatePlayer(player)
	if err != nil {
		log.Printf("Error creating player: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
		return
	}

	log.Printf("Player created with ID: %v", player.ID)
	ctx.JSON(http.StatusCreated, player)
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
