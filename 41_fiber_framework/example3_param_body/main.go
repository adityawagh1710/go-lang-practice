package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	app := fiber.New()

	app.Post("/user/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		user := new(User)

		if err := c.Bind().Body(user); err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		return c.JSON(map[string]any{
			"id":   id,
			"user": user,
		})
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
