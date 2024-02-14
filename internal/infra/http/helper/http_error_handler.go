package helper

import (
	"errors"
	"net/http"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/entity"
	"github.com/labstack/echo/v4"
)

func HandleHttpError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, entity.ErrCustomerNotFound):
		return c.NoContent(http.StatusNotFound)
	case errors.Is(err, entity.ErrCustomerNoLimit):
		return c.NoContent(http.StatusUnprocessableEntity)
	default:
		return c.NoContent(http.StatusInternalServerError)
	}
}
