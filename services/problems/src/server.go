package main

import (
	"log"
	"problems/config"
	"problems/handler"
	"problems/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env");
	if err != nil {
		log.Println("Error load env -> " + err.Error());
	}

	err = config.InitDatabase();
	if err != nil {
		panic("Error init database -> " + err.Error());
	}

	app := fiber.New();

	app.Use("/problems/new", middleware.VerifyToken);

	app.Use("/problems/new/public", middleware.CheckAdmin);

	app.Get("/problems", func (ctx fiber.Ctx) error {
		return ctx.SendString("Hello problems");
	});

	app.Post("/problems/new/public/meta", func (ctx fiber.Ctx) error {
		return handler.NewPublicProblem(ctx);
	});

	app.Post("/problems/new/public/files", middleware.CheckSameProblem, func (ctx fiber.Ctx) error {
		return handler.UploadPublicProblemFile(ctx);
	});

	app.Listen(":3001");
}