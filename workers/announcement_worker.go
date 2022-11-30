package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"study-event-go-asynq/domains/interfaces"
	"time"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type AnnouncementWorker struct {
	announcementRepo interfaces.AnnouncementRepository
}

func NewAnnouncementWorker(announcementRepo interfaces.AnnouncementRepository) *AnnouncementWorker {
	return &AnnouncementWorker{
		announcementRepo: announcementRepo,
	}
}

func (a *AnnouncementWorker) Announce(ctx context.Context, t *asynq.Task) error {
	task := struct { // TODO: use DTO?
		AnnouncementID uint64
	}{}
	if err := json.Unmarshal(t.Payload(), &task); err != nil {
		return err
	}

	announcement, err := a.announcementRepo.Read(ctx, task.AnnouncementID)
	if err != nil {
		zap.S().Errorw("fail to retrieve an announce", "err", err)
		return err
	}

	println("[ANNOUNCEMENT]", fmt.Sprintf("GOT A NEW ANNOUNCEMENT FROM %s.", announcement.From))
	time.Sleep(1 * time.Second)
	println("[ANNOUNCEMENT]", "THE MESSAGE IS ...")
	time.Sleep(1 * time.Second)
	println("[ANNOUNCEMENT]", "\""+announcement.Message+"\"")

	return nil
}

func (a *AnnouncementWorker) AnnounceWithTime(ctx context.Context, t *asynq.Task) error {
	task := struct { // TODO: use DTO?
		AnnouncementID uint64
	}{}
	if err := json.Unmarshal(t.Payload(), &task); err != nil {
		return err
	}

	announcement, err := a.announcementRepo.Read(ctx, task.AnnouncementID)
	if err != nil {
		zap.S().Errorw("fail to retrieve an announcement", "id", task.AnnouncementID, "err", err)
		return err
	}

	announcement, err = a.announcementRepo.Update(ctx, announcement)
	if err != nil {
		zap.S().Errorw("fail to update an announcement", "id", announcement.ID, "err", err)
		return err
	}

	println("[ANNOUNCEMENT]", fmt.Sprintf("GOT A NEW ANNOUNCEMENT FROM %s.", announcement.From))
	time.Sleep(1 * time.Second)
	println("[ANNOUNCEMENT]", "THE MESSAGE IS ...")
	time.Sleep(1 * time.Second)
	println("[ANNOUNCEMENT]", "\""+announcement.Message+"\"")

	return nil
}
