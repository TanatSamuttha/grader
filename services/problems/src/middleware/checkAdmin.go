package middleware

import (
	"problems/config"
	"problems/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func CheckAdmin(ctx fiber.Ctx) error {
	uid := ctx.Locals("uid");
	_, err := gorm.G[models.User](config.DB).Select("role").Where("uid = ? AND role = ?", uid, "Admin").First(ctx);
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}

	return ctx.Next();
}