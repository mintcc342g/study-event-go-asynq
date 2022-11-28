package infrastructures

import (
	"context"
	"math/rand"
	"study-event-go-asynq/conf"
	"study-event-go-asynq/domains"
	"study-event-go-asynq/domains/interfaces"
	"study-event-go-asynq/infrastructures/orm"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// gormAnnouncementRepository ...
type gormAnnouncementRepository struct {
	conn *gorm.DB
}

// NewGormAnnouncementRepository ...
func NewGormAnnouncementRepository(db *gorm.DB) interfaces.AnnouncementRepository {
	migrations := []interface{}{
		&orm.Announcement{},
	}
	if err := db.Set("gorm:table_options", conf.TableDefaultCharset).Migrator().AutoMigrate(migrations...); err != nil {
		zap.S().Panicw("fail to migrate automatically", "err", err)
	}

	return &gormAnnouncementRepository{
		conn: db,
	}
}

func (g *gormAnnouncementRepository) Save(ctx context.Context, announcement *domains.Announcement) (*domains.Announcement, error) {
	// TODO: use a database

	announcement.ID = rand.Uint64()

	return announcement, nil
}
