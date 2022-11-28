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

	ch := make(chan error, 1)
	go func() {
		ch <- a.announceForTime(ctx, task.AnnouncementID)
	}()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		zap.S().Errorw("task is not completed within the time", "taskID", t.ResultWriter().TaskID(), "announcementID", task.AnnouncementID, "err", err.Error())
		return err

	case taskErr := <-ch:
		if taskErr != nil {
			zap.S().Errorw("announce error", "err", taskErr)
		}
		return nil
	}
}

func (a *AnnouncementWorker) announceForTime(ctx context.Context, announcementID uint64) error {
	announcement, err := a.announcementRepo.Read(ctx, announcementID)
	if err != nil {
		zap.S().Errorw("fail to retrieve an announcement", "id", announcementID, "err", err)
		return err
	}

	announcement, err = a.announcementRepo.Update(ctx, announcement)
	if err != nil {
		zap.S().Errorw("fail to update an announcement", "id", announcement.ID, "err", err)
		return err
	}

	return nil
}
