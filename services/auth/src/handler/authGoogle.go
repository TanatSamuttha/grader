package handler

import (
	"auth/logic"
	"auth/models"
	"log"

	"github.com/gofiber/fiber/v3"
)

func AuthGoogle(ctx fiber.Ctx) error {
	var token models.TokenDTO;
	if err := ctx.Bind().Body(&token); err != nil {
		log.Println("Error get token from request -> " + err.Error());
		return err;
	}

	jwt, err := logic.GoogleAuthen(ctx, token.Token);
	if err != nil {
		log.Println("Error Google authentication -> " + err.Error());
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "Bearer",
		Value:    jwt,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		MaxAge:   60 * 60 * 24 * 3,
	})

	return ctx.SendStatus(fiber.StatusOK);
}