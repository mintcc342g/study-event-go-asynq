package containers

import (
	"study-event-go-asynq/controllers"
)

// ControllerContainer ...
type ControllerContainer struct {
	AnnouncementCtrl *controllers.AnnouncementController
}

func newControllerContainer(svcContainer *ApplicationContainer, repoContainer *InfrastructureContainer) *ControllerContainer {
	return &ControllerContainer{
		AnnouncementCtrl: controllers.NewAnnouncementController(svcContainer.AnnouncementSvc),
	}
}

// InitControllerContainer ...
func InitControllerContainer(svcContainer *ApplicationContainer, repoContainer *InfrastructureContainer) *ControllerContainer {
	return newControllerContainer(svcContainer, repoContainer)
}
