package bootstrap

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hsjsjsj009/go-beans"
)

type Bootstrap struct {
	DepContainer *beans.ProviderContainer
	Validator *validator.Validate
	Translator ut.Translator
	MainRouter fiber.Router
}
