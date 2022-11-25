package infrastructures

import (
	"context"
	"study-event-go-asynq/domains/interfaces"

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
		return err
	}

	zap.S().Infow("Success to send a task", "task_id", taskInfo.ID, "task_state", taskInfo.State)

	return nil
}
