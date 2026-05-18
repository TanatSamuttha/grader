package handler

import (
	"fmt"
	"problems/config"
	"problems/logic"
	"problems/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewProblem(ctx fiber.Ctx) error {
	var problemMeta models.ProblemDTO;
	if err := ctx.Bind().Body(&problemMeta); err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	fmt.Println(problemMeta);

	author := ctx.Locals("uid").(string);
	problemID := uuid.New();

	problemPDF, err := ctx.FormFile("problem");
	if err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesZip, err := ctx.FormFile("testcases");
	if err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesSize, err := logic.TestCasesCount(testcasesZip);
	if err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	problem := models.Problem{
		ProblemID: problemID.String(),
		Name: problemMeta.Name,
		AuthorUID: author,
		TestCasesSize: testcasesSize,
		TimeLimit: problemMeta.TimeLimit,
		MemoryLimit: problemMeta.MemoryLimit,
	};

	err = gorm.G[models.Problem](config.DB).Create(ctx, &problem);
	if err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	err = logic.SaveProblem(problemPDF, testcasesZip);
	if err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	return ctx.SendStatus(fiber.StatusOK);
}