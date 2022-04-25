package queue

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type pubsubQueue struct {
	client *pubsub.Client
}

var _ Queue = (*pubsubQueue)(nil)

func (p pubsubQueue) Consume(ctx context.Context, id string, handler Handler) {
	sub := p.client.Subscription(id)
	sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		err := handler(ctx, string(msg.Data))
		if err != nil {
			fmt.Printf("Error handling message: %v", err)
			msg.Nack()
			return
		}
		fmt.Println("Message handled successfully")
		msg.Ack()
	})
}

func NewPubsubQueue(client *pubsub.Client) *pubsubQueue {
	return &pubsubQueue{
		client: client,
	}
}
