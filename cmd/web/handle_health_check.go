package web

import (
	"github.com/machester4/products-leaderboard/infra/server"
)

func handleHealthCheck(ctx server.Context) error {
	return ctx.SendString("OK")
}
