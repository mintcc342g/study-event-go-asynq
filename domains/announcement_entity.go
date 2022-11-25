package domains

import (
	"context"
	"encoding/json"
	"errors"
	"study-event-go-asynq/applications/dto"
	"time"
)

type Announcement struct {
	ID      uint64
	From    string
	Message string
	Timeout time.Duration
}

func NewAnnouncement(ctx context.Context, announcementDTO dto.Announcement) (*Announcement, error) {
	if announcementDTO.Message == "" {
		return nil, errors.New("invalid message")
	}

	return &Announcement{
		From:    announcementDTO.From,
		Message: announcementDTO.Message,
		Timeout: time.Duration(announcementDTO.Seconds * int64(time.Second)),
	}, nil
}

func (a *Announcement) NewEventPayload() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Announcement) WithTimeout() bool {
	return a.Timeout != 0 // timeout can be set a negative number for test
}
