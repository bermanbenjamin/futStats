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
	events, err := s.repo.GetEventsByPlayerId(playerId)

	if err != nil {
		return nil, err
	}
    
	return events, nil
}

func (s *EventService) CreateEvent(event *model.Event) (*model.Event, error) {
	player, err := s.playerService.UpdatePlayerByEvent(*event)

	if err != nil {
        return nil, err
    }

	event.Player = *player

    newEvent, err := s.repo.CreateEvent(event)
    if err != nil {
        return nil, err
    }

    return newEvent, nil
}

func (s *EventService) UpdateEvent(event model.Event) (*model.Event, error) {
	updatedEvent, err := s.repo.UpdateEvent(&event)
    if err != nil {
        return nil, err
    }

    player, err := s.playerService.UpdatePlayerByEvent(event)

    if err != nil {
        return nil, err
    }

    updatedEvent.Player = *player

    return updatedEvent, nil
}

func (s *EventService) DeleteEvent(id uuid.UUID) error {
    return s.repo.DeleteEvent(id)
}