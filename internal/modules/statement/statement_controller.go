package statement

import (
	"net/http"
	"strconv"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/shared/helper"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/clientes/:customerId/extrato", getStatement)
}

func getStatement(c echo.Context) error {
	customerIdStr := c.Param("customerId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	stt, err := GetStatement(customerId)
	if err != nil {
		return helper.HandleHttpError(c, err)
	}

	return c.JSON(http.StatusOK, stt)
}
