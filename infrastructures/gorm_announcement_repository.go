package infrastructures

import (
	"context"
	"math/rand"
	"study-event-go-asynq/domains"
	"study-event-go-asynq/domains/interfaces"
)

// gormAnnouncementRepository ...
type gormAnnouncementRepository struct {
}

// NewGormAnnouncementRepository ...
func NewGormAnnouncementRepository() interfaces.AnnouncementRepository {
	return &gormAnnouncementRepository{}
}

func (g *gormAnnouncementRepository) Save(ctx context.Context, announcement *domains.Announcement) (*domains.Announcement, error) {
	// TODO: use a database

	announcement.ID = rand.Uint64()

	return announcement, nil
}
