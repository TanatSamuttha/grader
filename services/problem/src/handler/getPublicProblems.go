package handler

import (
	"log"
	"problem/config"
	"problem/models"

	"github.com/gofiber/fiber/v3"
)

func GetPublicProblems(ctx fiber.Ctx) error {
	var previews []models.ProblemPreviewDTO;
	err := config.DB.Model(&models.Problem{}).Select("problem_id", "name").Where("visibility = ?", "public").Find(&previews).Error;
	if err != nil {
		log.Println("Error query problems from database -> " + err.Error());
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}
	problemsList := models.ProblemsListDTO{Problems: previews};
	return ctx.JSON(problemsList);
}