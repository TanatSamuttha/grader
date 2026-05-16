package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"auth/config"
	"auth/logic"
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

	app.Post("/authen", func (ctx fiber.Ctx) error {
		var token models.TokenDTO;
		if err := ctx.Bind().Body(&token); err != nil {
			return err;
		}

		jwt, err := logic.GoogleAuthen(ctx, token.Token);

		if err != nil {
			return ctx.SendStatus(401);
		}

		var jwtToken models.TokenDTO;
		jwtToken.Token = jwt;

		return ctx.JSON(jwtToken);
	});

	app.Listen(":3000");
}