package repositories

import (
	"context"

	"github.com/machester4/products-leaderboard/internal/application/dto"
)

type ProductLeadboard interface {
	GetTopProducts(ctx context.Context, limit int) ([]dto.ProductLeadboard, error)
	IncrementScore(ctx context.Context, productID string) error
}
