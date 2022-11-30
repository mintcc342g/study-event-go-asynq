package workers

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type AsynqHandler func(context.Context, *asynq.Task) error

func TimeoutMiddleware(handler AsynqHandler) AsynqHandler {
	return func(ctx context.Context, t *asynq.Task) error {

		ch := make(chan error, 1)
		go func() {
			ch <- handler(ctx, t)
		}()

		select {
		case <-ctx.Done():
			err := ctx.Err()
			zap.S().Errorw("task is not completed within the time", "taskID", t.ResultWriter().TaskID(), "err", err.Error())
			return asynq.SkipRetry // archive
			// return nil // not archive

		case taskErr := <-ch:
			if taskErr != nil {
				zap.S().Errorw("announce error", "err", taskErr)
			}
			return taskErr
		}
	}
}
