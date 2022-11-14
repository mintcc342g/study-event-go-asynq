package interfaces

import (
	"context"
	"study-event-go-asynq/domains"
)

type AnnouncementRepository interface {
	Save(ctx context.Context, announcement *domains.Announcement) (*domains.Announcement, error)
}
