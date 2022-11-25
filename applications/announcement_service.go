package applications

import (
	"context"
	"study-event-go-asynq/applications/dto"
	"study-event-go-asynq/domains"
	"study-event-go-asynq/domains/interfaces"
)

// AnnouncementService ...
type AnnouncementService struct {
	announcementRepo interfaces.AnnouncementRepository
	taskRepo         interfaces.TaskRepository
}

// NewAnnouncementsService ...
func NewAnnouncementsService(announcementRepo interfaces.AnnouncementRepository, taskRepo interfaces.TaskRepository) *AnnouncementService {
	return &AnnouncementService{
		announcementRepo: announcementRepo,
		taskRepo:         taskRepo,
	}
}

// Schedule a message to announce.
func (a *AnnouncementService) Schedule(ctx context.Context, announcementDTO dto.Announcement) (*dto.Announcement, error) {
	announcement, err := domains.NewAnnouncement(ctx, announcementDTO)
	if err != nil {
		return nil, err
	}

	announcement, err = a.announcementRepo.Save(ctx, announcement)
	if err != nil {
		return nil, err
	}

	payload, err := announcement.NewEventPayload()
	if err != nil {
		return nil, err
	}

	if err = a.taskRepo.SendTask(ctx, announcement.TaskKey(), payload); err != nil {
		return nil, err
	}

	return &dto.Announcement{
		ID:      announcement.ID,
		From:    announcement.From,
		Message: announcement.Message,
	}, nil
}
