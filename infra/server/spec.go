package server

import "context"

type Server interface {
	Get(path string, handler Handler)
	Post(path string, handler Handler)
	Listen(port string) error
}

type Handler func(ctx Context) error

type Context interface {
	Context() context.Context
	SendJSON(data interface{}) error
	SendString(data string) error
	GetIntParam(name string) (int, error)
	GetStringParam(name string) string
	GetQuery(name string) string
	GetQueryInt(name string) (int, error)
}
