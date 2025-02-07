package services

import (
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type EventService struct {
	repo          *repository.EventsRepository
	playerService *PlayerService
}

func NewEventService(repo *repository.EventsRepository, playerService *PlayerService) *EventService {
	return &EventService{repo: repo, playerService: playerService}
}

func (s *EventService) GetEventById(id uuid.UUID) (*models.Event, error) {
	return s.repo.GetEventById(id)
}

func (s *EventService) GetAllEventsByPlayerId(playerId uuid.UUID) ([]models.Event, error) {
	events, err := s.repo.GetEventsByPlayerId(playerId)

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *EventService) CreateEvent(event *models.Event) (*models.Event, error) {
	player, err := s.playerService.UpdatePlayerByEvent(*event, true)

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

func (s *EventService) UpdateEvent(event models.Event) (*models.Event, error) {
	updatedEvent, err := s.repo.UpdateEvent(&event)
	if err != nil {
		return nil, err
	}

	return updatedEvent, nil
}

func (s *EventService) DeleteEvent(id uuid.UUID) error {
	event, err := s.GetEventById(id)
	if err != nil {
		return err
	}

	if event == nil {
		return nil
	}

	_, err = s.playerService.UpdatePlayerByEvent(*event, false)

	if err != nil {
		return err
	}

	return s.repo.DeleteEvent(id)
}
