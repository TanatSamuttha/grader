package main

import (
	"user/handler"
	"user/middleware"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New();

	app.Get("/user/me", middleware.VerifyToken, handler.GetUserData);

	app.Listen("3005");
}