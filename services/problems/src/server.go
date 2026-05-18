package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New();

	app.Get("/problems", func (ctx fiber.Ctx) error {
		return ctx.SendString("Hello problems");
	});

	app.Post("/problems/new", func (ctx fiber.Ctx) error {
		name := ctx.Body.
	});

	app.Listen(":3001");
}