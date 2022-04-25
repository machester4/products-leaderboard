package web

import (
	"github.com/machester4/products-leaderboard/infra/server"
	"github.com/machester4/products-leaderboard/internal/application/services"
)

func handleGetTopProducts(ctx server.Context, s services.ProductLeadboard) error {
	limit, _ := ctx.GetQueryInt("limit")

	products, err := s.GetTopProducts(ctx.Context(), limit)
	if err != nil {
		return err
	}

	return ctx.SendJSON(products)
}
