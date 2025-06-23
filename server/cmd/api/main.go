package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kido1611/image-server/internal/delivery/http"
)

func main() {
	app := fiber.New(fiber.Config{})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
	}))

	router := http.NewRouter(app)
	router.Setup()

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
