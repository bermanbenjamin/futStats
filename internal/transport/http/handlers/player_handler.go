package handler

import (
	"log"
	"net/http"

	model "github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PlayerHandler struct {
	service *service.PlayerService
}

func NewPlayerHandler(service *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}

func (h *PlayerHandler) GetAllPlayers(ctx *gin.Context) {
	filterQuery := ctx.GetHeader("x-api-query-filter")
	filterValue := ctx.GetHeader("x-api-query-value")
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

	ctx.JSON(http.StatusOK, gin.H{"data": players, "total": len(players)} )
}

func (h *PlayerHandler) GetPlayer(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Printf("Invalid player ID format: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id format"})
		return
	}

	player, err := h.service.GetPlayer(id)
	if err != nil {
		log.Printf("Error getting player with ID %v: %v", id, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get player"})
		return
	}

	if player == nil {
		log.Printf("Player with ID %v not found", id)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Player does not exist"})
		return
	}

	ctx.JSON(http.StatusOK, player)
}

func (h *PlayerHandler) CreatePlayer(ctx *gin.Context) {
	var player *model.Player
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

	var player model.Player
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
