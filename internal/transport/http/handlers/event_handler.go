package handler

import (
	"net/http"

	model "github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventsHandler struct {
	service *service.EventService
	playerService *service.PlayerService
}

func NewEventsHandler(service *service.EventService, playerService *service.PlayerService) *EventsHandler {
    return &EventsHandler{service: service, playerService: playerService}
}

func (h *EventsHandler) GetEventsByPlayerId(c *gin.Context) {
	idStr := c.Param("playerId")
    playerId, err := uuid.Parse(idStr)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID format"})
        return
    }

    events, err := h.service.GetAllEventsByPlayerId(playerId)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": events})
}

func (h *EventsHandler) CreateEventsHandler(c *gin.Context) {
	var event model.Event
    if err := c.ShouldBindJSON(&event); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    newEvent, err := h.service.CreateEvent(&event)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
        return
    }

	player, err := h.playerService.UpdatePlayerByEvent(event)

    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player's stats"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"event": newEvent, "player": player})
}
