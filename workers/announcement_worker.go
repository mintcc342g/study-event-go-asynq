package workers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"study-event-go-asynq/domains"
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
	var announcement domains.Announcement
	if err := json.Unmarshal(t.Payload(), &announcement); err != nil {
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
	var announcement domains.Announcement
	if err := json.Unmarshal(t.Payload(), &announcement); err != nil {
		return err
	}

	ch := make(chan error, 1)
	go func() {
		ch <- a.announceForTime(ctx, announcement)
	}()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		zap.S().Errorw("task is not completed within the time", "announcement_id", announcement.ID, "timeout", announcement.Timeout, "deadline", announcement.Deadline, "err", err.Error())
		return err

	case taskErr := <-ch:
		if taskErr != nil {
			zap.S().Errorw("announce error", "err", taskErr)
		}
		return nil
	}
}

func (a *AnnouncementWorker) announceForTime(ctx context.Context, announcement domains.Announcement) error {
	println("[ANNOUNCEMENT]", fmt.Sprintf("GOT A NEW ANNOUNCEMENT FROM %s.", announcement.From))
	time.Sleep(1 * time.Second)
	println("[ANNOUNCEMENT]", "THE MESSAGE IS ...")
	time.Sleep(1 * time.Second)
	println("[ANNOUNCEMENT]", "\""+announcement.Message+"\"")

	return errors.New("done~") // for a test
}
