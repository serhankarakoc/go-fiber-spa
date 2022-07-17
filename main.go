package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/web", "dist")

	app.Get("/web/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./dist/index.html")
	})

	log.Fatal(app.Listen(":3000"))
}
