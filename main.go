package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// serve Single Page application on "/web"
	// assume static file at dist folder
	app.Static("/web", "dist")

	app.Get("/web/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./dist/index.html")
	})

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
