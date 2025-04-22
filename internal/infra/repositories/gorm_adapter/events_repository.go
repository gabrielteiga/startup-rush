package gorm_adapter

import (
	"github.com/gabrielteiga/startup-rush/database"
	"github.com/gabrielteiga/startup-rush/internal/domain/entities/event_entity"
	"gorm.io/gorm"
)

type EventsGORMRepository struct {
	DB *gorm.DB
}

func NewEventsGORMRepository(db *gorm.DB) *EventsGORMRepository {
	return &EventsGORMRepository{
		DB: db,
	}
}

func (er *EventsGORMRepository) Create(name string, score int) *event_entity.Event {
	event := &database.Events{
		Name:  name,
		Score: score,
	}

	if err := er.DB.Create(&event).Error; err != nil {
		return nil
	}

	return event_entity.NewEvent(event.ID, event.Name, event.Score)
}

func (er *EventsGORMRepository) List() ([]*event_entity.Event, error) {
	var events []database.Events

	err := er.DB.Select("id, name, score").Find(&events).Error
	if err != nil {
		return nil, err
	}

	var eventEntities []*event_entity.Event
	for _, eventModel := range events {
		event := event_entity.NewEvent(eventModel.ID, eventModel.Name, eventModel.Score)
		eventEntities = append(eventEntities, event)
	}

	return eventEntities, nil
}
