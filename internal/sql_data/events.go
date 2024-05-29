package sql_data

import (
	"context"
	"fmt"
	"watchy/internal/model"

	"gorm.io/gorm"
)

type WatchEventStorer interface {
	Get(ctx context.Context, id int) (model.WatchEvent, error)
	Create(ctx context.Context, WatchEvent model.WatchEvent) (model.WatchEvent, error)
	GetByUserID(ctx context.Context, userID string) ([]model.WatchEvent, error)
}

type WatchEventStore struct {
	db *gorm.DB
}

func NewWatchEventStore(db *gorm.DB) WatchEventStorer {
	return &WatchEventStore{
		db: db,
	}
}

// function to get specific user watch events
func (e WatchEventStore) GetByUserID(ctx context.Context, userID string) ([]model.WatchEvent, error) {
	query := fmt.Sprintf("select * from events where user_id='%s'", userID)
	event := []model.WatchEvent{}
	result := e.db.Raw(query).Scan(&event)
	if result.Error != nil {
		return []model.WatchEvent{}, result.Error
	}
	return event, nil
}

// function to insert an user watch events
func (e WatchEventStore) Create(ctx context.Context, event model.WatchEvent) (model.WatchEvent, error) {
	sqlQuery := "INSERT INTO events (user_id, event_title,video_id,action,duration) VALUES (?, ?, ?, ?, ?)"

	// Execute the raw SQL query with parameters
	result := e.db.Exec(sqlQuery, event.UserID, event.EventTitle, event.VideoID, event.Action, event.Duration)
	if result.Error != nil {
		return model.WatchEvent{}, result.Error
	}
	var lastInsertedID uint
	e.db.Raw("SELECT LAST_INSERT_ID()").Scan(&lastInsertedID)

	return e.Get(ctx, int(lastInsertedID))
}

// function to get specific watchevent by ID
func (e WatchEventStore) Get(ctx context.Context, id int) (model.WatchEvent, error) {
	query := fmt.Sprintf("select * from events where id=%v", id)
	event := model.WatchEvent{}
	result := e.db.Raw(query).Scan(&event)
	if result.Error != nil {
		return model.WatchEvent{}, result.Error
	}
	return event, nil
}
