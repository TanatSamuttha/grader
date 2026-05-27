package main

import (
	"grade/config"
	"grade/handler"
	"grade/logic"
	"grade/middleware"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/contrib/websocket"
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

	err = config.InitDockerClient();
	if err != nil {
		panic("Error init docker client -> " + err.Error());
	}

	log.Println(config.DockerClient);

	logic.SummonWorkers(2);

	app := fiber.New();

	app.Use("/grade", middleware.VerifyToken);

	app.Post("/grade/submit", handler.SubmissionHandler);

	app.Get("/grade/result", websocket.New(handler.ReturnResult));

	app.Listen(":3004");
}
