package main

import (
	"fmt"
	"log"
	"short-leaf/app/config/database"
	"short-leaf/app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.DatabaseSetup()

	fmt.Println("Starting Servcer")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello From Go")
	})
	routes.AuthRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
