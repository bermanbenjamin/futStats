package middlewares

import (
	"net/http"
	"slices"

	"github.com/bermanbenjamin/futStats/cmd/api/constants"
	"github.com/bermanbenjamin/futStats/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func OwnerMiddleware(ctx *gin.Context, dependencies *config.Dependencies) {
	playerId := ctx.GetString("player_id")
	playerUuid, err := uuid.Parse(playerId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	leagueSlug := ctx.Param("leagueSlug")

	league, err := dependencies.LeagueService.GetLeagueBy(constants.QueryFilter("slug"), leagueSlug)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
		return
	}

	memberIds := make([]uuid.UUID, len(league.Members))
	for i, member := range league.Members {
		memberIds[i] = member.ID
	}

	if league.OwnerId != playerUuid && !slices.Contains(memberIds, playerUuid) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not a member of this league"})
		return
	}

	ctx.Next()
}
