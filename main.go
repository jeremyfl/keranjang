package main

import (
	"log"

	"github.com/jeremylombogia/keranjang/controller"
	"github.com/jeremylombogia/keranjang/internal"
	"github.com/jeremylombogia/keranjang/repositories"
	"github.com/jeremylombogia/keranjang/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	redisClient := internal.NewDatabaseClient()

	repo := repositories.NewCartRedisRepo(redisClient)
	service := service.NewCartService(repo)
	controller := controller.NewCartController(service)

	// Routes
	e.GET("/api/v1/carts", controller.Get)
	e.POST("/api/v1/carts", controller.Post)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
