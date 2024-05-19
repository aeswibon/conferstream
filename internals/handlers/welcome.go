package handlers

import "github.com/gofiber/fiber/v2"

// Welcome is a handler that returns a welcome message
func Welcome(c *fiber.Ctx) error {
	return c.Render("welcome", fiber.Map{
		"Title": "Welcome to Fiber!",
	}, "layouts/main")
}