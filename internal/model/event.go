package model

type WatchEvent struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	UserID     string `json:"user_id"`
	EventTitle string `json:"event_title"`
	VideoID    string `json:"video_id"`
	Action     string `json:"action"`
	Duration   int    `json:"duration"`
}

func (w WatchEvent) TableName() string {
	return "events"
}
