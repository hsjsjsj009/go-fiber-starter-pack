package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-starter-pack/server/bootstrap/routers"
	"go-fiber-starter-pack/server/handlers"
	"net/http"
)

func (b *Bootstrap) RegisterRouters() {
	handler := &handlers.Handler{
		DepContainer: b.DepContainer,
		Validator: b.Validator,
		Translator: b.Translator,
	}

	// Testing
	b.MainRouter.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("work")
	})

	apiV1 := b.MainRouter.Group("/v1")

	//example
	exampleRoutes := routers.ExampleRoutes{RouterGroup: apiV1,Handler: handler}
	exampleRoutes.RegisterRoutes()

}
