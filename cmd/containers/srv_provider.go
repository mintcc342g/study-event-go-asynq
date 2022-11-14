package containers

import (
	"study-event-go-asynq/applications"
)

// ApplicationContainer ...
type ApplicationContainer struct {
	AnnouncementSvc *applications.AnnouncementService
}

func newApplicationContainer(infraContainer *InfrastructureContainer) *ApplicationContainer {
	return &ApplicationContainer{
		AnnouncementSvc: applications.NewAnnouncementsService(infraContainer.AnnouncementRepo, infraContainer.EventRepo),
	}
}

// InitApplicationContainer ...
func InitApplicationContainer(infraContainer *InfrastructureContainer) *ApplicationContainer {
	return newApplicationContainer(infraContainer)
}
