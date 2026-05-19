package middleware

import "github.com/gofiber/fiber/v3"

func CheckAdmin(ctx fiber.Ctx) error {
	role := ctx.Locals("role");
	if role != "Admin" {
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}

	return ctx.Next();
}