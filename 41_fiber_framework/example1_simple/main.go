package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World! from Fiber Framework")
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
