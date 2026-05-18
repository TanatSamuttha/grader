package main

import (
	"fmt"
	"problems/config"
	"problems/handler"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env");
	if err != nil {
		fmt.Println(err);
	}

	err = config.InitDatabase();
	if err != nil {
		panic(err);
	}

	app := fiber.New();

	app.Get("/problems", func (ctx fiber.Ctx) error {
		return ctx.SendString("Hello problems");
	});

	app.Post("/problems/new/all", func (ctx fiber.Ctx) error {
		return handler.NewProblem(ctx);
	});

	app.Listen(":3001");
}