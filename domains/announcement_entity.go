package domains

import (
	"context"
	"encoding/json"
	"errors"
	"study-event-go-asynq/applications/dto"
	"time"
)

type Announcement struct {
	ID       uint64
	From     string
	Message  string
	Timeout  time.Duration
	Deadline time.Time
}

func NewAnnouncement(ctx context.Context, announcementDTO dto.Announcement) (*Announcement, error) {
	if announcementDTO.Message == "" {
		return nil, errors.New("invalid message")
	}

	if !announcementDTO.Deadline.IsZero() && announcementDTO.Seconds != 0 {
		return nil, errors.New("only one option for time can be set.")
	}

	return &Announcement{
		From:     announcementDTO.From,
		Message:  announcementDTO.Message,
		Timeout:  time.Duration(announcementDTO.Seconds) * time.Second,
		Deadline: announcementDTO.Deadline,
	}, nil
}

func (a *Announcement) NewEventPayload() ([]byte, error) {
	task := struct { // TODO: use DTO
		AnnouncementID uint64
	}{
		AnnouncementID: a.ID,
	}
	return json.Marshal(task)
}

func (a *Announcement) WithTimeout() bool {
	return a.Timeout != 0 // timeout can be set a negative number for a test
}

func (a *Announcement) WithDeadline() bool {
	return !a.Deadline.IsZero()
}
