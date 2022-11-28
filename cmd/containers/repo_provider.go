package containers

import (
	"study-event-go-asynq/domains/interfaces"
	"study-event-go-asynq/infrastructures"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

// InfrastructureContainer ...
type InfrastructureContainer struct {
	AnnouncementRepo interfaces.AnnouncementRepository
	TaskRepo         interfaces.TaskRepository
}

func newInfrastructureContainer(db *gorm.DB, client *asynq.Client) *InfrastructureContainer {
	return &InfrastructureContainer{
		AnnouncementRepo: infrastructures.NewGormAnnouncementRepository(db),
		TaskRepo:         infrastructures.NewTaskRepository(client),
	}
}

// InitInfrastructureContainer ...
func InitInfrastructureContainer(db *gorm.DB, client *asynq.Client) *InfrastructureContainer {
	return newInfrastructureContainer(db, client)
}
