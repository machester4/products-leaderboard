package server

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app *fiber.App
}

var _ Server = (*fiberServer)(nil)

func (f fiberServer) Get(path string, handler Handler) {
	f.app.Get(path, func(c *fiber.Ctx) error {
		return handler(fiberAdaptedContext{c})
	})
}

func (f fiberServer) Post(path string, handler Handler) {
	f.app.Post(path, func(c *fiber.Ctx) error {
		return handler(fiberAdaptedContext{c})
	})
}

type fiberAdaptedContext struct {
	ctx *fiber.Ctx
}

func (c fiberAdaptedContext) SendJSON(data interface{}) error {
	return c.ctx.JSON(data)
}

func (c fiberAdaptedContext) SendString(data string) error {
	return c.ctx.SendString(data)
}

func (c fiberAdaptedContext) Context() context.Context {
	return c.ctx.Context()
}

func (c fiberAdaptedContext) GetIntParam(name string) (int, error) {
	return c.ctx.ParamsInt(name)
}

func (c fiberAdaptedContext) GetStringParam(name string) string {
	return c.ctx.Params(name)
}

func (c fiberAdaptedContext) GetQuery(name string) string {
	return c.ctx.Query(name)
}

func (c fiberAdaptedContext) GetQueryInt(name string) (int, error) {
	return strconv.Atoi(c.ctx.Query(name))
}

func (f fiberServer) Listen(port string) error {
	return f.app.Listen(port)
}

func NewFiberServer() Server {
	return &fiberServer{
		app: fiber.New(),
	}
}
