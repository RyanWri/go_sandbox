package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	// Respond with "Hello, test!" on root path, "/test"
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Test!")
	})

    app.Listen(":3000")
}