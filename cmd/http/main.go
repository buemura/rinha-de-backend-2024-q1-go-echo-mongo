package main

import (
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/config"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/infra/database"
	"github.com/buemura/rinha-de-backend-2024-q1-go-echo-mongo/internal/infra/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	e := echo.New()
	setupServerMiddlewares(e)
	host := ":" + config.PORT
	e.Start(host)
}

func setupServerMiddlewares(app *echo.Echo) {
	app.Use(middleware.Recover())
	app.Use(middleware.Secure())
	routes.SetupRoutes(app)
}
