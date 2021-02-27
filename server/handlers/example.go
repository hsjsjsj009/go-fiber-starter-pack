package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-starter-pack/usecases"
	"gorm.io/gorm"
	"net/http"
)

type ExampleHandler struct {
	*Handler
}

func(h *ExampleHandler) Example(ctx *fiber.Ctx) error {
	var (
		tx *gorm.DB
	)

	uc := usecases.NewExampleUC()
	err := h.DepContainer.InjectStruct(uc)
	if err != nil {
		return h.SendErrorResponse(ctx,err.Error(),http.StatusBadRequest)
	}

	err = h.DepContainer.InjectVariable(&tx)
	if err != nil {
		return h.SendErrorResponse(ctx,err.Error(),http.StatusInternalServerError)
	}

	tx = tx.Begin()

	uc.BeginTx(tx)

	data,err := uc.ExampleCase()
	if err != nil {
		tx.Rollback()
		return h.SendErrorResponse(ctx,err.Error(),http.StatusBadRequest)
	}

	tx.Commit()
	return h.SendSuccessResponse(ctx,data,nil)
}
