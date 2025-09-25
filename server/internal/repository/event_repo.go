package repository

import (
	"errors"
	"log"

	"github.com/bermanbenjamin/futStats/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventsRepository struct {
	db *gorm.DB
}

func NewEventsRepository(db *gorm.DB) *EventsRepository {
	return &EventsRepository{db: db}
}

func (r *EventsRepository) GetEventById(id uuid.UUID) (*models.Event, error) {
	var event models.Event
	if err := r.db.Preload("Player").Preload("Assistent").Preload("Match").Where("id = ?", id).First(&event).Error; err != nil {
		log.Printf("Error getting event with ID %v: %v", id, err)
		return nil, err
	}
	return &event, nil
}

func (r *EventsRepository) GetEventsByPlayerId(playerId uuid.UUID) ([]models.Event, error) {

	var events []models.Event
	if result := r.db.Preload("Player").Preload("Assistent").Preload("Match").Where("player_id = ?", playerId).Find(&events); result.Error != nil {
		log.Printf("Error getting events for player with ID %v: %v", playerId, result.Error)
		return nil, errors.New("Could not find events for player with ID: " + playerId.String())
	}

	return events, nil
}

func (r *EventsRepository) CreateEvent(event *models.Event) (createdEvent *models.Event, err error) {
	if err := r.db.Create(event).Error; err != nil {
		log.Println(`Error to create event in database: `, err)
		return nil, err
	}
	return event, nil
}

func (r *EventsRepository) DeleteEvent(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Event{}).Error; err != nil {
		log.Printf("Error deleting event with ID %v: %v", id, err)
		return err
	}
	return nil
}

func (r *EventsRepository) UpdateEvent(event *models.Event) (*models.Event, error) {
	if err := r.db.Model(&models.Event{}).Where("id = ?", event.ID).Updates(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

// CreateEventWithTransaction creates an event within a transaction
func (r *EventsRepository) CreateEventWithTransaction(event *models.Event, fn func(tx *gorm.DB) error) (*models.Event, error) {
	var createdEvent *models.Event

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := fn(tx); err != nil {
			return err
		}

		// Create the event in the transaction
		if err := tx.Create(event).Error; err != nil {
			log.Println(`Error creating event in transaction: `, err)
			return err
		}

		createdEvent = event
		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdEvent, nil
}

// CreateEventInTransaction creates an event using an existing transaction
func (r *EventsRepository) CreateEventInTransaction(tx *gorm.DB, event *models.Event) (*models.Event, error) {
	if err := tx.Create(event).Error; err != nil {
		log.Println(`Error creating event in transaction: `, err)
		return nil, err
	}
	return event, nil
}

// DeleteEventWithTransaction deletes an event within a transaction
func (r *EventsRepository) DeleteEventWithTransaction(id uuid.UUID, fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

// CountEventsByPlayerAndType counts events for a specific player and event type
func (r *EventsRepository) CountEventsByPlayerAndType(playerId uuid.UUID, eventType string) (int, error) {
	var count int64
	err := r.db.Model(&models.Event{}).
		Where("player_id = ? AND type = ?", playerId, eventType).
		Count(&count).Error

	if err != nil {
		log.Printf("Error counting events for player %v and type %s: %v", playerId, eventType, err)
		return 0, err
	}

	return int(count), nil
}

// CountEventsByPlayer counts all events for a specific player
func (r *EventsRepository) CountEventsByPlayer(playerId uuid.UUID) (int, error) {
	var count int64
	err := r.db.Model(&models.Event{}).
		Where("player_id = ?", playerId).
		Count(&count).Error

	if err != nil {
		log.Printf("Error counting events for player %v: %v", playerId, err)
		return 0, err
	}

	return int(count), nil
}

// GetPlayerStatsSummary gets a summary of all event types for a player
func (r *EventsRepository) GetPlayerStatsSummary(playerId uuid.UUID) (map[string]int, error) {
	var results []struct {
		Type  string `json:"type"`
		Count int    `json:"count"`
	}

	err := r.db.Model(&models.Event{}).
		Select("type, COUNT(*) as count").
		Where("player_id = ?", playerId).
		Group("type").
		Scan(&results).Error

	if err != nil {
		log.Printf("Error getting stats summary for player %v: %v", playerId, err)
		return nil, err
	}

	stats := make(map[string]int)
	for _, result := range results {
		stats[result.Type] = result.Count
	}

	return stats, nil
}
