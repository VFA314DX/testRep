package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./static")

	app.Get("/foo", func(c *fiber.Ctx) error {
		fmt.Println("Request1")
		return c.SendString("Gegege")
	})

	app.Get("/bar", func(c *fiber.Ctx) error {
		fmt.Println("Request2")
		return c.SendString("Memem")
	})

	app.Get("/too", func(c *fiber.Ctx) error {
		fmt.Println("Request3")
		return c.SendString("tototo")
	})

	app.Listen(":8080")
}
