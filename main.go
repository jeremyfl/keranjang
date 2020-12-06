package main

import (
	"github.com/jeremylombogia/keranjang/controller"
	"github.com/jeremylombogia/keranjang/internal"
	"github.com/jeremylombogia/keranjang/repositories"
	"github.com/jeremylombogia/keranjang/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	redisClient := internal.NewDatabaseClient()

	repo := repositories.NewCartRepo(redisClient)
	service := service.NewCartService(repo)
	controller := controller.NewCartController(service)

	// Routes
	e.GET("/api/v1/carts", controller.Get)
	e.POST("/api/v1/carts", controller.Post)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
