package main

import (
	"fmt"
	"short-leaf/app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Starting Servcer")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello From Go")
	})
	routes.AuthRoutes(app)
	app.Listen(":3000")
}
