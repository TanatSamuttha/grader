package main

import (
	"log"
	"problem/config"
	"problem/handler"
	"problem/middleware"

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

	app.Use("/problem/new", middleware.VerifyToken);

	app.Use("/problem/new/public", middleware.CheckAdmin);

	app.Get("/problem", func (ctx fiber.Ctx) error {
		return ctx.SendString("Hello problems");
	});

	app.Post("/problem/new/public/meta", func (ctx fiber.Ctx) error {
		return handler.NewPublicProblem(ctx);
	});

	app.Post("/problem/new/public/files", middleware.CheckSameProblem, func (ctx fiber.Ctx) error {
		return handler.UploadPublicProblemFile(ctx);
	});

	app.Get("/problem/public", func (ctx fiber.Ctx) error {
		return handler.GetPublicProblems(ctx);
	});

	app.Listen(":3001");
}