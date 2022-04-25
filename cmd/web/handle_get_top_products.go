package web

import (
	"github.com/machester4/products-leaderboard/infra/server"
	"github.com/machester4/products-leaderboard/internal/application/services"
)

func handleGetTopProducts(ctx server.Context, s services.ProductLeadboard) error {
	limit, err := ctx.GetQueryInt("limit")
	if err != nil {
		return err
	}

	products, err := s.GetTopProducts(ctx.Context(), limit)
	if err != nil {
		return err
	}

	return ctx.SendJSON(products)
}
