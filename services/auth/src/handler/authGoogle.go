package handler

import (
	"auth/logic"
	"log"

	"github.com/gofiber/fiber/v3"
)

func AuthGoogle(ctx fiber.Ctx) error {
	token := ctx.Get("token");

	jwt, err := logic.GoogleAuthen(ctx, token);
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