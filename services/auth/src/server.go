package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"auth/config"
	"auth/handler"
	"auth/middleware"
)

func main(){
	err := godotenv.Load("../.env");
	if err != nil {
		log.Println("Error load env -> " + err.Error());
	}

	err = config.InitDatabase();
	if err != nil {
		panic(err);
	}

	err = config.InitFirebase();
	if err != nil {
		panic(err);
	}

	app := fiber.New();

	app.Get("/auth", func (ctx fiber.Ctx) error {
		return ctx.SendString("Hello auth service");
	});

	app.Get("/auth/me", middleware.VerifyToken, func (ctx fiber.Ctx) error {
		return handler.GetUserData(ctx);
	})

	app.Post("/auth/google", func (ctx fiber.Ctx) error {
		return handler.AuthGoogle(ctx);
	});

	app.Listen(":3000");
}