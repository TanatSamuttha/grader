package middleware

import (
	"log"
	"problems/config"
	"problems/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func CheckSameProblem(ctx fiber.Ctx) error {
	uid := ctx.Locals("uid");
	problemID := ctx.Cookies("CreatingProblemID");
	_, err := gorm.G[models.Problem](config.DB).Select("problem_id").Where("problem_id = ? AND author_uid = ?", problemID, uid).First(ctx);
	if err != nil {
		log.Println("Error check same problem -> " + err.Error());
		return ctx.SendStatus(fiber.StatusUnauthorized);
	}
	return ctx.Next();
}