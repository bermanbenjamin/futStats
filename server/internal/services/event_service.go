package services

import (
	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/bermanbenjamin/futStats/internal/repository"
	"github.com/google/uuid"
)

type EventServiceInterface interface {
	GetEventById(id uuid.UUID) (*models.Event, error)
	GetAllEventsByPlayerId(playerId uuid.UUID) ([]models.Event, error)
	CreateEvent(event *models.Event) (*models.Event, error)
	UpdateEvent(event models.Event) (*models.Event, error)
	DeleteEvent(id uuid.UUID) error
}

type EventService struct {
	repo *repository.EventsRepository
}

func NewEventService(repo *repository.EventsRepository) *EventService {
	return &EventService{repo: repo}
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
	return s.repo.CreateEvent(event)
}

func (s *EventService) UpdateEvent(event models.Event) (*models.Event, error) {
	updatedEvent, err := s.repo.UpdateEvent(&event)
	if err != nil {
		return nil, err
	}
	return updatedEvent, nil
}

func (s *EventService) DeleteEvent(id uuid.UUID) error {
	return s.repo.DeleteEvent(id)
}
