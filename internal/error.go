package internal

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type CustomError struct {
	Type    error
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("[type] %s, [msg] %s", c.Type.Error(), c.Message)
}

func GenerateError(errorType error, msg string) CustomError {
	return CustomError{
		Type:    errorType,
		Message: msg,
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	var e CustomError
	if errors.As(err, &e) {
		if errors.Is(e.Type, ErrValidationError) {
			code = http.StatusBadRequest
		}

		if errors.Is(e.Type, ErrDataNotFound) {
			code = http.StatusNotFound
		}

		if errors.Is(e.Type, ErrInternalServer) {
			log.Println(err)
			return ctx.Status(http.StatusInternalServerError).JSON(JSONResponse{
				Message: "internal server error",
			})
		}

		return ctx.Status(code).JSON(JSONResponse{
			Message: e.Type.Error(),
			Error:   e.Message,
		})
	}

	log.Println(err)
	return ctx.Status(http.StatusInternalServerError).JSON(JSONResponse{
		Message: "internal server error",
	})
}
