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
		panic("Error init database -> " + err.Error());
	}

	err = config.InitFirebase();
	if err != nil {
		panic("Error init firebase -> " + err.Error());
	}

	app := fiber.New();

	app.Get("/auth", func (ctx fiber.Ctx) error {
		return ctx.SendString("Hello auth service");
	});

	app.Get("/auth/me", middleware.VerifyToken, handler.GetUserData);

	app.Post("/auth/google", handler.AuthGoogle);

	app.Listen(":3000");
}