package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-starter-pack/server/handlers"
)

type ExampleRoutes struct {
	RouterGroup fiber.Router
	Handler *handlers.Handler
}

func (r *ExampleRoutes) RegisterRoutes()  {
	handler := handlers.ExampleHandler{
		Handler: r.Handler,
	}

	exampleGroups := r.RouterGroup.Group("/example")
	exampleGroups.Get("/main",handler.Example)
}
