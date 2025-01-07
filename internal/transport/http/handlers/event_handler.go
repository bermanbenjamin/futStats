package handler

import "github.com/bermanbenjamin/futStats/internal/service"

type EventsHandler struct {
	service *service.EventsService
}

func NewEventsHandler(service *service.EventsService) *EventsHandler {
    return &EventsHandler{service: eventsService}
}

