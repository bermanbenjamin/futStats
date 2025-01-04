package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bermanbenjamin/futStats/api/models"
	"github.com/bermanbenjamin/futStats/api/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetPlayers(ctx *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	players, err := repositories.GetAllPlayers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get players"})
		return
	}

	ctx.JSON(http.StatusOK, players)
}

func GetPlayer(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	player, err := repositories.GetPlayerById(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to get player"})
		return
	}

	if player.ID == uuid.Nil {
		ctx.JSON(404, gin.H{"error": "Player does not exist"})
		return
	}

	ctx.JSON(http.StatusOK, player)
}

func CreatePlayer(ctx *gin.Context) {
	var player models.Player
	if err := ctx.ShouldBindJSON(&player); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := repositories.AddPlayer(player)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to add player"})
		return
	}

	ctx.JSON(http.StatusCreated, player)
}

func UpdatePlayer(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		ctx.JSON(400, gin.H{"error": "Missing id query parameter"})
		return
	}

	var player models.Player
	if err := ctx.ShouldBindJSON(&player); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	uuid, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	player.ID = uuid
	updatedPlayer, err := repositories.UpdatePlayer(player)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update player"})
		return
	}

	ctx.JSON(http.StatusOK, updatedPlayer)
}

func DeletePlayer(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	err = repositories.DeletePlayer(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete player"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Player deleted"})
}
