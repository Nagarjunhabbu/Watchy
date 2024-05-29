package service

import (
	"context"
	"errors"
	"testing"
	"watchy/internal/model"
)

func TestGetUserAPI(t *testing.T) {
	service := watchEventService{
		Data: MockSql{},
	}

	tc1 := struct {
		userID   string
		expected *model.WatchEvent
		err      error
	}{
		userID:   "",
		expected: nil,
		err:      errors.New("userID is empty"),
	}

	_, err := service.GetWatchEvent(context.Background(), tc1.userID)
	if err == nil {
		t.Errorf("failed test case, expected %v got nil", tc1.err)
	}

	tc2 := struct {
		userID   string
		expected []model.WatchEvent
		err      error
	}{
		userID: "user123",
		expected: []model.WatchEvent{
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
		},
		err: nil,
	}
	data, err := service.GetWatchEvent(context.Background(), tc2.userID)
	if err != nil {
		t.Errorf("failed test case, expected no error got %v  ", err)
	}
	if len(data) != len(tc2.expected) {
		t.Errorf("failed test case, expected %d  got %d  ", len(data), len(tc2.expected))

	}

	tc3 := struct {
		userID   string
		expected *model.WatchEvent
		err      error
	}{
		userID:   "user999",
		expected: nil,
		err:      errors.New("invalid userID"),
	}

	_, err = service.GetWatchEvent(context.Background(), tc3.userID)
	if err == nil {
		t.Errorf("failed test case, expected %v got nil", tc3.err)
	}

}

func TestCreateUserAPI(t *testing.T) {
	service := watchEventService{
		Data: MockSql{},
	}
	event := model.WatchEvent{
		ID:         2,
		UserID:     "user124",
		EventTitle: "",
		VideoID:    "v7231",
		Action:     "WATCH",
		Duration:   360,
	}

	tc1 := struct {
		event    model.WatchEvent
		expected *model.WatchEvent
		err      error
	}{
		event:    event,
		expected: nil,
		err:      errors.New("event title is empty"),
	}

	_, err := service.CreateWatchEvent(context.Background(), tc1.event)
	if err == nil {
		t.Errorf("failed test case, expected %v got nil", tc1.err)
	}

	event.EventTitle = "IPL Final -2024"
	tc2 := struct {
		event    model.WatchEvent
		expected *model.WatchEvent
		err      error
	}{
		event: event,
		expected: &model.WatchEvent{
			ID:         2,
			UserID:     "user124",
			EventTitle: "IPL final -2024",
			VideoID:    "v7231",
			Action:     "WATCH",
			Duration:   360,
		},
		err: nil,
	}
	data, err := service.CreateWatchEvent(context.Background(), tc2.event)
	if err != nil {
		t.Errorf("failed test case, expected no error got %v  ", err)
	}
	if data.ID != tc2.expected.ID {
		t.Errorf("failed test case, expected %v  got %v  ", tc2.expected.ID, data.ID)

	}
}
