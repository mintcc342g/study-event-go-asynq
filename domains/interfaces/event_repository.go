package interfaces

import "context"

type EventRepository interface {
	SendTask(ctx context.Context, key string, payload []byte) error
}
