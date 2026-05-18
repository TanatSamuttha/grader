package main

import (
	"problems/config"
	"problems/handler"

	"github.com/gofiber/fiber/v3"
)

func main() {
	err := config.InitDatabase();
	if err != nil {
		panic(err);
	}

	app := fiber.New();

	app.Get("/problems", func (ctx fiber.Ctx) error {
		return ctx.SendString("Hello problems");
	});

	app.Post("/problems/new", func (ctx fiber.Ctx) error {
		return handler.NewProblem(ctx);
	});

	app.Listen(":3001");
}