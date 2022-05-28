package routes

import (
	"github.com/gofiber/fiber/v2"
	home "lewiseman.com/leo_app/pkg/controllers/home_pg"
)

func AppRoutes(app fiber.Router) {
	app.Static("/", "./public")
	// api := app.Group("/api")
	app.Get("/", home.HomePg())
}
