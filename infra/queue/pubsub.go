package queue

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type pubsubQueue struct {
	client *pubsub.Client
}

var _ Queue = (*pubsubQueue)(nil)

func (p pubsubQueue) Consume(ctx context.Context, topic string, handler Handler) {
	sub := p.client.Subscription(topic)
	sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		handler(ctx, string(msg.Data))
		msg.Ack()
	})
}

func NewPubsubQueue(client *pubsub.Client) *pubsubQueue {
	return &pubsubQueue{
		client: client,
	}
}
