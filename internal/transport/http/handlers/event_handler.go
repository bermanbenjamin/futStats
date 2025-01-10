package handler

import (
	"net/http"

	"github.com/bermanbenjamin/futStats/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventsHandler struct {
	service *service.EventService
}

func NewEventsHandler(service *service.EventService) *EventsHandler {
    return &EventsHandler{service: service}
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




