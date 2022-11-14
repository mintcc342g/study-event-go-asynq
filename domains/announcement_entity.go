package domains

import (
	"context"
	"encoding/json"
	"errors"
	"study-event-go-asynq/applications/dto"
	"study-event-go-asynq/consts"
)

type Announcement struct {
	ID      uint64
	From    string
	Message string
}

func NewAnnouncement(ctx context.Context, announcementDTO dto.Announcement) (*Announcement, error) {
	if announcementDTO.Message == "" {
		return nil, errors.New("invalid message")
	}

	return &Announcement{
		From:    announcementDTO.From,
		Message: announcementDTO.Message,
	}, nil
}

func (a *Announcement) NewEventPayload() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Announcement) TaskKey() string {
	return consts.AnnouncementTaskKey
}
