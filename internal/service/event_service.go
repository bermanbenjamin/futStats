package service

import (
	model "github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type EventService struct { 
	repo *repository.EventsRepository
	playerService *PlayerService 
}

func NewEventService(repo *repository.EventsRepository, playerService *PlayerService) *EventService {
    return &EventService{repo: repo, playerService: playerService}
}

func (s *EventService) GetAllEventsByPlayerId(playerId uuid.UUID) ([]model.Event, error) {
	events := s.repo.GetEventsByPlayerId(playerId)
    
	return events, nil
}