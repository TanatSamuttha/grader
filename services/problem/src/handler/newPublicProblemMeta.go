package handler

import (
	"errors"
	"log"
	"problem/config"
	"problem/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewPublicProblem(ctx fiber.Ctx) error {
	var problemDTO models.ProblemDTO;
	if err := ctx.Bind().Body(&problemDTO); err != nil {
		log.Println(errors.New("Error get problem from body -> " + err.Error()));
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	author := ctx.Locals("uid").(string);
	problemID := uuid.New();

	problem := models.Problem{
		ProblemID: problemID.String(),
		Name: problemDTO.Name,
		AuthorUID: author,
		TimeLimit: problemDTO.TimeLimit,
		MemoryLimit: problemDTO.MemoryLimit,
	};

	if err := gorm.G[models.Problem](config.DB).Create(ctx, &problem); err != nil {
		log.Println(errors.New("Error create new problem mata data in database -> " + err.Error()));
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "CreatingProblemID",
		Value:    problemID.String(),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		MaxAge:   60 * 3,
	})

	return ctx.SendStatus(fiber.StatusOK);
}