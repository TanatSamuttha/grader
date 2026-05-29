package main

import (
	"log"
	"user/config"
	"user/handler"
	"user/middleware"

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

	app.Get("/user/me", middleware.VerifyToken, handler.GetUserData);

	app.Listen("3005");
}