package routes

import (
	"short-leaf/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App) {
	route := a.Group("/api/auth")

	route.Get("/", controllers.UserLogin)
}
