package infrastructures

import (
	"context"
	"study-event-go-asynq/domains/interfaces"
	"time"

	"go.uber.org/zap"

	"github.com/hibiken/asynq"
)

// asynqTaskRepository ...
type asynqTaskRepository struct {
	client *asynq.Client
}

// NewTaskRepository ...
func NewTaskRepository(client *asynq.Client) interfaces.TaskRepository {
	return &asynqTaskRepository{
		client: client,
	}
}

func (a *asynqTaskRepository) SendTask(ctx context.Context, key string, payload []byte) error {
	taskInfo, err := a.client.EnqueueContext(ctx, asynq.NewTask(key, payload))
	if err != nil {
		zap.S().Errorw("Fail to send a task", "err", err)
		return err
	}

	zap.S().Infow("Success to send a task", "task_id", taskInfo.ID, "task_state", taskInfo.State)

	return nil
}

func (a *asynqTaskRepository) SendTaskWithTimeout(ctx context.Context, key string, payload []byte, timeout time.Duration) error {
	taskInfo, err := a.client.EnqueueContext(ctx, asynq.NewTask(key, payload), asynq.Timeout(timeout))
	if err != nil {
		zap.S().Errorw("Fail to send a task with timeout", "err", err)
		return err
	}

	zap.S().Infow("Success to send a task with timeout", "task_id", taskInfo.ID, "task_state", taskInfo.State)

	return nil
}
