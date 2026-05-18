package handler

import (
	"fmt"
	"problems/config"
	"problems/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewPublicProblem(ctx fiber.Ctx) error {
	var problemDTO models.ProblemDTO;
	if err := ctx.Bind().Body(&problemDTO); err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	fmt.Println(problemDTO);

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
		fmt.Println(err);
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