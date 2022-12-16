package internal

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Service interface {
	GenerateQuote(ctx context.Context, input GenerateQuote) (*Quote, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GenerateQuote(c *fiber.Ctx) error {
	var reqParams GenerateQuote
	err := c.QueryParser(&reqParams)
	if err != nil {
		return GenerateError(ErrInternalServer, fmt.Sprintf("error when parse query params: %s", err.Error()))
	}

	resp, err := h.service.GenerateQuote(c.Context(), reqParams)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(JSONResponse{
		Message: "OK",
		Data:    resp,
	})
}
