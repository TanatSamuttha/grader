package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"auth/config"
	"auth/models"
)

func main(){
	err := godotenv.Load();
	if err != nil {
		panic(err);
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

	app.Get("/", func (c fiber.Ctx) error {
		return c.SendString("Hello auth service");
	});

	app.Post("/authen", func (c fiber.Ctx) error {
		var token models.TokenDTO;
		if err := c.Bind().Body(&token); err != nil {
			return err;
		}

		jwt, err := Authen(token.Token);

		if err != nil {
			return c.SendStatus(401);
		}

		var jwtToken models.TokenDTO;
		jwtToken.Token = jwt;

		return c.JSON(jwtToken);
	});

	app.Listen(":3000");
}