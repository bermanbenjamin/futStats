package repository

import (
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

func (r *EventsRepository) GetEventsByPlayerId(playerId uuid.UUID) []model.Event {

	var events []model.Event = make([]model.Event, 0)
	if result := r.db.Where("player_id =?", playerId).Find(&events); result.Error != nil {
		log.Println(`Error to find events returned from database for player with id: `, playerId) 
		panic(result.Error)
    }
	return events
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
