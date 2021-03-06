package messaging

import (
	"context"

	"github.com/machester4/products-leaderboard/infra/queue"
	"github.com/machester4/products-leaderboard/internal/application/services"
)

type messaging struct {
	q queue.Queue
}

func (m messaging) ConsumeIncrementScore(s services.ProductLeadboard) {
	m.q.Consume(context.Background(), "voucher-metric-sub", func(ctx context.Context, msg interface{}) error {
		return s.IncrementScore(ctx, msg.(string))
	})
}

func New(q queue.Queue) *messaging {
	return &messaging{q: q}
}
