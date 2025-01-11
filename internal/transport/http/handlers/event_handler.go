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
}

func NewEventsHandler(service *service.EventService, playerService *service.PlayerService) *EventsHandler {
    return &EventsHandler{service: service}
}

func (h *EventsHandler) GetEventsByPlayerIdHandler(c *gin.Context) {
	idStr := c.Param("playerId")
    playerId, err := uuid.Parse(idStr)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID format"})
        return
    }

    events, err := h.service.GetAllEventsByPlayerIdHandler(playerId)
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

    c.JSON(http.StatusCreated, gin.H{"data": newEvent,})
}

func (h *EventsHandler) UpdateEventsHandler(c *gin.Context) {
    var event model.Event
    if err := c.ShouldBindJSON(&event); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    updatedEvent, err := h.service.UpdateEvent(event)

    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": updatedEvent})
}

func (h *EventsHandler) DeleteEventHandler(c *gin.Context) {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID format"})
        return
    }

    if err := h.service.DeleteEvent(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}