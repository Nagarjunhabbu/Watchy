package service

import (
	"context"
	"errors"
	"log"
	"watchy/internal/model"
	"watchy/internal/sql_data"
)

type WatchEventService interface {
	GetWatchEvent(ctx context.Context, userID string) ([]model.WatchEvent, error)
	CreateWatchEvent(ctx context.Context, WatchEvent model.WatchEvent) (model.WatchEvent, error)
}

type watchEventService struct {
	Data sql_data.WatchEventStorer
}

func NewWatchEventService(storer sql_data.WatchEventStorer) WatchEventService {
	return &watchEventService{Data: storer}
}

func (e watchEventService) GetWatchEvent(ctx context.Context, userID string) ([]model.WatchEvent, error) {
	if userID == "" {
		return []model.WatchEvent{}, errors.New("userID is empty")
	}
	events, err := e.Data.GetByUserID(ctx, userID)
	if err != nil {
		log.Println("error in getting WatchEvent", err)
		return []model.WatchEvent{}, err
	}
	if len(events) == 0 {
		log.Println("invalid userID", err)
		return []model.WatchEvent{}, errors.New("invalid userID")
	}

	return events, nil
}

func (e watchEventService) CreateWatchEvent(ctx context.Context, req model.WatchEvent) (model.WatchEvent, error) {

	if req.UserID == "" {
		log.Println("userID is empty")
		return model.WatchEvent{}, errors.New("userID is empty")
	}

	if req.VideoID == "" {
		log.Println("video ID is empty")
		return model.WatchEvent{}, errors.New("video ID is empty")

	}

	if req.EventTitle == "" {
		log.Println("event title is empty")
		return model.WatchEvent{}, errors.New("event title is empty")
	}

	event, err := e.Data.Create(ctx, req)
	if err != nil {
		log.Println("error in creating WatchEvent", err)
		return model.WatchEvent{}, err
	}

	return event, nil
}
