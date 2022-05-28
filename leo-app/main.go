package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"lewiseman.com/leo_app/pkg/routes"
)

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	routes.AppRoutes(app)
	app.Listen(":" + port)
}
