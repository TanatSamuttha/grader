package middleware

import (
	"fmt"
	"problems/config"
	"problems/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func CheckAdmin(ctx fiber.Ctx) error {
	uid := ctx.Locals("uid");
	user, err := gorm.G[models.User](config.DB).Select("role").Where("uid = ? AND role = ?", uid, "Admin").First(ctx);
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}
	fmt.Println(user);

	return ctx.Next();
}