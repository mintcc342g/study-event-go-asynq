package infrastructures

import (
	"context"
	"study-event-go-asynq/domains/interfaces"

	"go.uber.org/zap"

	"github.com/hibiken/asynq"
)

// asynqEventRepository ...
type asynqEventRepository struct {
	client *asynq.Client
}

// NewEventRepository ...
func NewEventRepository(client *asynq.Client) interfaces.EventRepository {
	return &asynqEventRepository{
		client: client,
	}
}

func (e *asynqEventRepository) SendTask(ctx context.Context, key string, payload []byte) error {
	taskInfo, err := e.client.EnqueueContext(ctx, asynq.NewTask(key, payload))
	if err != nil {
		return err
	}

	zap.S().Infow("Success to send a task", "task_id", taskInfo.ID, "task_state", taskInfo.State)

	return nil
}
