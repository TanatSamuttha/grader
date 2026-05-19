package handler

import (
	"auth/config"
	"auth/models"
	"log"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetUserData(ctx fiber.Ctx) error {
	uid := ctx.Locals("uid").(string);
	user, err := gorm.G[models.User](config.DB).Where("uid = ?", uid).First(ctx);
	userDTO := models.UserDTO{
		Username: user.Username,
		PhotoURL: user.PhotoURL,
	};
	if err != nil {
		log.Println("Error query user from database -> " + err.Error());
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}
	return ctx.JSON(userDTO);
}