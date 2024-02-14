package controller

import (
	"net/http"
	"strconv"

	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/application"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/entity"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/infra/http/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateTransaction(c echo.Context) error {
	customerIdStr := c.Param("customerId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	body := new(entity.CreateTransactionRequest)
	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}
	if body.Type != "c" && body.Type != "d" {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	trx, err := application.CreateTransaction(customerId, body)
	if err != nil {
		return helper.HandleHttpError(c, err)
	}
	return c.JSON(http.StatusOK, trx)
}
