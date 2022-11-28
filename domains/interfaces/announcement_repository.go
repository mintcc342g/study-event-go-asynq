package interfaces

import (
	"context"
	"study-event-go-asynq/domains"
)

type AnnouncementRepository interface {
	Read(ctx context.Context, id uint64) (*domains.Announcement, error)
	Save(ctx context.Context, announcement *domains.Announcement) (*domains.Announcement, error)
	Update(ctx context.Context, announcement *domains.Announcement) (*domains.Announcement, error)
}
