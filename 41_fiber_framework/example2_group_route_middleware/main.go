package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(func(c fiber.Ctx) error {
		c.Set("X-Custom-Header", "Fiber-Is-Fast")
		return c.Next()
	})

	// Grouping
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "success", "data": "users"})
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
