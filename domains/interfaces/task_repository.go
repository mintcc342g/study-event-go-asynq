package interfaces

import (
	"context"
	"time"
)

type TaskRepository interface {
	SendTask(ctx context.Context, key string, payload []byte) error
	SendTaskWithTimeout(ctx context.Context, key string, payload []byte, timeout time.Duration) error
}
