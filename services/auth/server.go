package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"auth/config"
	"auth/logic"
	"auth/models"
)

func main(){
	err := godotenv.Load();
	if err != nil {
		fmt.Println(err);
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
	}));

	app.Get("/", func (c fiber.Ctx) error {
		return c.SendString("Hello auth service");
	});

	app.Post("/authen/google", func (ctx fiber.Ctx) error {
		var token models.TokenDTO;
		if err := ctx.Bind().Body(&token); err != nil {
			return err;
		}

		jwt, err := logic.GoogleAuthen(ctx, token.Token);

		if err != nil {
			fmt.Println(err);
			return ctx.SendStatus(401);
		}

		var jwtToken models.TokenDTO;
		jwtToken.Token = jwt;

		ctx.Cookie(&fiber.Cookie{
			Name: "Bearer",
			Value: jwt,
			HTTPOnly: true,
			Secure: false,
			SameSite: "Lax",
			MaxAge: 60 * 60 * 24 * 3,
		});

		ctx.Cookie(&fiber.Cookie{
			Name: "IsAuthenticated",
			Value: "true",
			HTTPOnly: false,
			Secure: false,
			SameSite: "None",
			MaxAge: 60 * 60 * 24 *3,
		});

		return ctx.SendStatus(200);
	});

	app.Listen(":3000");
}