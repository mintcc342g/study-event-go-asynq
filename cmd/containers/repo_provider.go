package containers

import (
	"study-event-go-asynq/domains/interfaces"
	"study-event-go-asynq/infrastructures"

	"github.com/hibiken/asynq"
)

// InfrastructureContainer ...
type InfrastructureContainer struct {
	AnnouncementRepo interfaces.AnnouncementRepository
	EventRepo        interfaces.EventRepository
}

func newInfrastructureContainer(client *asynq.Client) *InfrastructureContainer {
	return &InfrastructureContainer{
		AnnouncementRepo: infrastructures.NewGormAnnouncementRepository(),
		EventRepo:        infrastructures.NewEventRepository(client),
	}
}

// InitInfrastructureContainer ...
func InitInfrastructureContainer(client *asynq.Client) *InfrastructureContainer {
	return newInfrastructureContainer(client)
}
