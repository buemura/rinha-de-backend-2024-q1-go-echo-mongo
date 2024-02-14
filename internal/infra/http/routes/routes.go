package routes

import (
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/infra/http/controller"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/clientes/:customerId/extrato", controller.GetStatement)
	e.POST("/clientes/:customerId/transacoes", controller.CreateTransaction)
}
