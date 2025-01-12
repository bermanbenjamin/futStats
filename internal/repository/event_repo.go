package repository

import (
	"errors"
	"log"

	model "github.com/bermanbenjamin/futStats/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventsRepository struct {
	db *gorm.DB
}

func NewEventsRepository(db *gorm.DB) *EventsRepository {
    return &EventsRepository{db: db}
}

func (r *EventsRepository) GetEventById(id uuid.UUID) (*model.Event, error) {
    var event model.Event
    if err := r.db.Preload("Player").Preload("Assistent").Preload("Match").Where("id = ?", id).First(&event).Error; err!= nil {
        log.Printf("Error getting event with ID %v: %v", id, err)
        return nil, err
    }
    return &event, nil
}

func (r *EventsRepository) GetEventsByPlayerId(playerId uuid.UUID) ([]model.Event, error) {

	var events []model.Event = make([]model.Event, 0)
	if result := r.db.Preload("Player").Preload("Assistent").Preload("Match").Where("player_id = ?", playerId).Find(&events); result != nil {
		log.Printf("Error getting events for player with ID %v: %v", playerId, result.Error)
        return nil, errors.New("Could not find events for player with ID: " + playerId.String())
    }
	return events, nil
}

func (r *EventsRepository) CreateEvent(event *model.Event) (createdEvent *model.Event, err error) { 
    if err := r.db.Create(event).Error; err != nil {
        log.Println(`Error to create event in database: `, err)
        return nil, err
    }
    return event, nil
}

func (r *EventsRepository) DeleteEvent(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Event{}).Error; err != nil {
        log.Printf("Error deleting event with ID %v: %v", id, err)
        return err
    }
    return nil
}


func (r *EventsRepository) UpdateEvent(event *model.Event) (*model.Event, error) {
    if err := r.db.Model(&model.Event{}).Where("id = ?", event.ID).Updates(event).Error; err != nil {
        return nil, err 
    }
    return event, nil 
}
