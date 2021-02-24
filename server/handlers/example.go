package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-starter-pack/db"
	"go-fiber-starter-pack/usecases"
	"net/http"
)

type ExampleHandler struct {
	*Handler
}

func(h *ExampleHandler) Example(ctx *fiber.Ctx) error {
	var (
		Tx db.SQLTx
	)

	uc := usecases.NewExampleUC()
	err := h.DepContainer.InjectStruct(uc)
	if err != nil {
		return h.SendErrorResponse(ctx,err.Error(),http.StatusBadRequest)
	}

	err = h.DepContainer.InjectVariable(&Tx)
	if err != nil {
		return h.SendErrorResponse(ctx,err.Error(),http.StatusInternalServerError)
	}

	uc.BeginTx(Tx)

	data,err := uc.ExampleCase()
	if err != nil {
		Tx.Rollback()
		return h.SendErrorResponse(ctx,err.Error(),http.StatusBadRequest)
	}

	Tx.Commit()
	return h.SendSuccessResponse(ctx,data,nil)
}
