package server

import (
	"github.com/aeswibon/conferstream/cmd"
	"github.com/aeswibon/conferstream/internals/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

// Run starts the server
func Run() error {
	env := cmd.NewEnv()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(logger.New())
	app.Use(cors.New())
	app.Get("/", handlers.Welcome)
	return app.Listen(":" + env.Addr)
}