package web

import (
	"github.com/machester4/products-leaderboard/infra/server"
	"github.com/machester4/products-leaderboard/internal/application/services"
)

type web struct {
	server server.Server
}

func (w web) InitializeRoutes(s services.ProductLeadboard) {
	w.server.Get("/", handleHealthCheck)
	w.server.Get("/stats", func(ctx server.Context) error {
		return handleGetTopProducts(ctx, s)
	})
}

func New(s server.Server) web {
	return web{
		server: s,
	}
}
