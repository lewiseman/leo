package homepg

import "github.com/gofiber/fiber/v2"

func HomePg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendFile("./templates/index.html", true)
	}
}
