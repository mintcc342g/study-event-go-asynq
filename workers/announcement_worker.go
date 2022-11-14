package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"study-event-go-asynq/domains"
	"study-event-go-asynq/domains/interfaces"
	"time"

	"github.com/hibiken/asynq"
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
