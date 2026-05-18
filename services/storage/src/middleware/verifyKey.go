package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

func VerifyKey(ctx fiber.Ctx) error {
	key := ctx.Get("Storage-Key");
	if key != os.Getenv("STORAGE_KEY") {
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}
	return ctx.Next();
}