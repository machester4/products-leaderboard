package queue

import "context"

type Queue interface {
	Consume(ctx context.Context, id string, handler Handler)
}

type Handler func(ctx context.Context, msg interface{}) error
