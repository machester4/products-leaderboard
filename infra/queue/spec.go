package queue

import "context"

type Queue interface {
	Consume(ctx context.Context, topic string, handler Handler)
}

type Handler func(ctx context.Context, msg interface{})
