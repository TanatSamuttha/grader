package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

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

	app.Get("/auth", func (ctx fiber.Ctx) error {
		return ctx.SendString("Hello auth service");
	});

	app.Get("/auth/me", logic.VerifyToken, func (ctx fiber.Ctx) error {
		uid := ctx.Locals("uid").(string);
		user, err := gorm.G[models.User](config.DB).Where("uid = ?", uid).First(ctx);
		userDTO := models.UserDTO{
			Username: user.Username,
			PhotoURL: user.PhotoURL,
		};
		if err != nil {
			fmt.Println(err);
			return ctx.SendStatus(401);
		}
		return ctx.JSON(userDTO);
	})

	app.Post("/auth/google", func (ctx fiber.Ctx) error {
		var token models.TokenDTO;
		if err := ctx.Bind().Body(&token); err != nil {
			return err;
		}

		jwt, err := logic.GoogleAuthen(ctx, token.Token);

		if err != nil {
			fmt.Println(err);
			return ctx.SendStatus(401);
		}

		ctx.Cookie(&fiber.Cookie{
			Name: "Bearer",
			Value: jwt,
			HTTPOnly: true,
			Secure: false,
			SameSite: "Lax",
			MaxAge: 60 * 60 * 24 * 3,
		});

		return ctx.SendStatus(200);
	});

	app.Listen(":3000");
}