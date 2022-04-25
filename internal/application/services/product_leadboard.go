package services

import (
	"context"

	"github.com/machester4/products-leaderboard/internal/application/dto"
	"github.com/machester4/products-leaderboard/internal/application/repositories"
)

type productLeadboard struct {
	repo repositories.ProductLeadboard
}

var _ ProductLeadboard = (*productLeadboard)(nil)

func (p productLeadboard) GetTopProducts(ctx context.Context, limit int) ([]dto.ProductLeadboard, error) {
	if limit == 0 {
		limit = 5
	}

	return p.repo.GetTopProducts(ctx, limit)
}

func (p productLeadboard) IncrementScore(ctx context.Context, productID string) error {
	return p.repo.IncrementScore(ctx, productID)
}

func NewProductLeadboard(r repositories.ProductLeadboard) *productLeadboard {
	return &productLeadboard{repo: r}
}
