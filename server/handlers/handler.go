package handlers

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hsjsjsj009/go-beans"
	"go-fiber-starter-pack/package/respond"
	"go-fiber-starter-pack/package/str"
	"net/http"
	"strings"
)

type Handler struct {
	DepContainer *beans.ProviderContainer
	Validator *validator.Validate
	Translator ut.Translator
}

//base send response
func (h Handler) SendResponse(ctx *fiber.Ctx, data interface{}, meta interface{}, err interface{}, code int) error {
	if code == 0 && err != nil {
		code = http.StatusUnprocessableEntity
		err = err.(error).Error()
	}

	if code != http.StatusOK && err != nil {
		return h.SendErrorResponse(ctx, err, code)
	}

	return h.SendSuccessResponse(ctx, data, meta)
}

//send response if status code 200
func (h Handler) SendSuccessResponse(ctx *fiber.Ctx, data interface{}, meta interface{}) error {
	response := respond.SuccessResponse(data, meta)

	return ctx.Status(http.StatusOK).JSON(response)
}

//send response if status code != 200
func (h Handler) SendErrorResponse(ctx *fiber.Ctx, err interface{}, code int) error {
	response := respond.ErrorResponse(err)

	return ctx.Status(code).JSON(response)
}

//extract error message from validator
func (h Handler) ExtractErrorValidationMessages(error validator.ValidationErrors) map[string][]string {
	errorMessage := map[string][]string{}
	errorTranslation := error.Translate(h.Translator)

	for _, err := range error {
		errKey := str.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
		)
	}

	return errorMessage
}

