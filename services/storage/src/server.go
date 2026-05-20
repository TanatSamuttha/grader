package main

import (
	"fmt"
	"storage/handler"
	"storage/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env");
	if err != nil {
		fmt.Println(err);
	}

	app := fiber.New();

	app.Use("/storage", middleware.VerifyKey);

	app.Post("/storage/upload/problem", func (ctx fiber.Ctx) error {
		return handler.UploadProblem(ctx);
	});

	app.Get("/storage/get/problem", func (ctx fiber.Ctx) error {
		return handler.GetProblem(ctx);
	});

	app.Listen(":3002");
}