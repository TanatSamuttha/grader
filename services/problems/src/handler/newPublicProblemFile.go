package handler

import (
	"fmt"
	"problems/config"
	"problems/logic"
	"problems/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UploadPublicProblemFile(ctx fiber.Ctx) error {
	problemPDF, err := ctx.FormFile("problem");
	if err != nil {
		fmt.Println("Error get problem PDF from form -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesZip, err := ctx.FormFile("testcases");
	if err != nil {
		fmt.Println("Error get testcase Zip from form -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesSize, err := logic.TestCasesCount(testcasesZip);
	if err != nil {
		fmt.Println("Error count testcases from form -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	
	uid := ctx.Locals("uid");
	problemID := ctx.Cookies("CreatingProblemID");

	if err := logic.SaveProblem(problemID, problemPDF, testcasesZip); err != nil {
		fmt.Println("Error save problem to storage -> " + err.Error())
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	row, err := gorm.G[models.Problem](config.DB).Where("problem_id = ? AND author_uid = ?", problemID, uid).Update(ctx, "test_cases_size", testcasesSize);
	if err != nil {
		fmt.Println("Error update testcases size -> " + err.Error());
	}
	if row == 0 {
		fmt.Println("Error update testcases size -> cant find creating problem");
	}

	return ctx.SendStatus(fiber.StatusOK);
}