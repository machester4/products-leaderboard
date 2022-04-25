package repositories

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/machester4/products-leaderboard/internal/application/dto"
	"github.com/machester4/products-leaderboard/internal/application/repositories"
)

type productLeadBoardRedis struct {
	client *redis.Client
}

var _ repositories.ProductLeadboard = (*productLeadBoardRedis)(nil)

func (p productLeadBoardRedis) GetTopProducts(ctx context.Context, limit int) ([]dto.ProductLeadboard, error) {
	if limit == 0 {
		limit = 5
	}

	redisOut, err := p.client.ZRevRangeWithScores(ctx, "products", 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	return p.buildGetTopProductsOutput(redisOut), nil
}

func (p productLeadBoardRedis) buildGetTopProductsOutput(rOutput []redis.Z) []dto.ProductLeadboard {
	var products []dto.ProductLeadboard
	for _, product := range rOutput {
		products = append(products, dto.ProductLeadboard{
			ProductID: product.Member.(string),
			Score:     int(product.Score),
		})
	}

	return products
}

func (p productLeadBoardRedis) IncrementScore(ctx context.Context, productID string) error {
	return p.client.ZIncrBy(ctx, "products", 1, productID).Err()
}

func NewProductLeadBoardRedis(client *redis.Client) *productLeadBoardRedis {
	return &productLeadBoardRedis{client: client}
}
