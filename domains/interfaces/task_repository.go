package interfaces

import (
	"context"
	"time"
)

type TaskRepository interface {
	SendTask(ctx context.Context, key string, payload []byte) error
	SendTaskWithTimeout(ctx context.Context, key string, payload []byte, timeout time.Duration) error
	SendTaskWithDeadline(ctx context.Context, key string, payload []byte, time time.Time) error
}
