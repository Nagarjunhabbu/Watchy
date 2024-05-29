package service

import (
	"context"
	"errors"
	"watchy/internal/model"
)

type MockSql struct {
}

func (e MockSql) Get(ctx context.Context, id int) (model.WatchEvent, error) {
	if id == 0 {
		return model.WatchEvent{}, errors.New("id is empty")
	}

	r := model.WatchEvent{
		ID:         1,
		UserID:     "user123",
		EventTitle: "Movie Night",
		VideoID:    "v456",
		Action:     "WATCH",
		Duration:   120,
	}
	return r, nil
}

func (e MockSql) Create(ctx context.Context, event model.WatchEvent) (model.WatchEvent, error) {
	if event.UserID == "" {
		return model.WatchEvent{}, errors.New("userID is empty")
	}

	if event.VideoID == "" {
		return model.WatchEvent{}, errors.New("video ID is empty")

	}

	if event.EventTitle == "" {
		return model.WatchEvent{}, errors.New("event title is empty")
	}

	r := model.WatchEvent{
		ID:         2,
		UserID:     "user124",
		EventTitle: "IPL final -2024",
		VideoID:    "v7231",
		Action:     "WATCH",
		Duration:   360,
	}
	return r, nil
}

func (e MockSql) GetByUserID(ctx context.Context, userID string) ([]model.WatchEvent, error) {
	if userID == "" {
		return []model.WatchEvent{}, errors.New("userID is empty")
	}

	if userID == "user999" {
		return []model.WatchEvent{}, errors.New("invalid userID")
	}

	r := []model.WatchEvent{
		{
			ID:         1,
			UserID:     "user123",
			EventTitle: "Movie Night",
			VideoID:    "v456",
			Action:     "WATCH",
			Duration:   120,
		},
		{
			ID:         2,
			UserID:     "user123",
			EventTitle: "IPL Playoffs -2024",
			VideoID:    "v912321",
			Action:     "WATCH",
			Duration:   600,
		},
	}
	return r, nil
}
