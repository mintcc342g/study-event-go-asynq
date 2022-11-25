package interfaces

import "context"

type TaskRepository interface {
	SendTask(ctx context.Context, key string, payload []byte) error
}
