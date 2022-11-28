package infrastructures

import (
	"context"
	"study-event-go-asynq/conf"
	"study-event-go-asynq/domains"
	"study-event-go-asynq/domains/interfaces"
	"study-event-go-asynq/infrastructures/orm"

	"github.com/jinzhu/copier"
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
	ormModel := &orm.Announcement{}
	if err := copier.Copy(ormModel, announcement); err != nil {
		zap.S().Errorw("fail to copy from entity", "err", err)
		return nil, err
	}

	if err := g.conn.Create(ormModel).Error; err != nil {
		zap.S().Errorw("fail to insert an announcement", "err", err)
		return nil, err
	}

	announcement.ID = ormModel.ID

	return announcement, nil
}

func (g *gormAnnouncementRepository) Read(ctx context.Context, id uint64) (*domains.Announcement, error) {
	ormModel := &orm.Announcement{}
	if err := g.conn.Model(ormModel).Where("id = ?", id).Find(ormModel).Error; err != nil {
		zap.S().Errorw("fail to get an announcement", "err", err)
		return nil, err
	}

	announcement := &domains.Announcement{}
	if err := copier.Copy(announcement, ormModel); err != nil {
		zap.S().Errorw("fail to copy from ormModel", "err", err)
		return nil, err
	}

	return announcement, nil
}

func (g *gormAnnouncementRepository) Update(ctx context.Context, announcement *domains.Announcement) (*domains.Announcement, error) {
	if err := g.conn.Model(&orm.Announcement{}).
		Where("id = ?", announcement.ID).
		Updates(map[string]interface{}{
			"updated_count": gorm.Expr("updated_count + ?", 1),
		}).Error; err != nil {
		zap.S().Errorw("fail to get an announcement", "err", err)
		return nil, err
	}

	return announcement, nil
}
