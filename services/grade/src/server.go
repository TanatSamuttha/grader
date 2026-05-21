package main

import (
	"grade/config"
	"grade/handler"
	"grade/middleware"
	"log"

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

	config.SummonWorkers(2);

	app := fiber.New();

	app.Post("/grade", middleware.VerifyToken, handler.SubmissionHandler);

	app.Listen(":3001");
}
