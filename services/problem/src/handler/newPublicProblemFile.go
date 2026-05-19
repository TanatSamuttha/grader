package handler

import (
	"log"
	"problem/config"
	"problem/logic"
	"problem/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UploadPublicProblemFile(ctx fiber.Ctx) error {
	problemPDF, err := ctx.FormFile("problem");
	if err != nil {
		log.Println("Error get problem PDF from form -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesZip, err := ctx.FormFile("testcases");
	if err != nil {
		log.Println("Error get testcase Zip from form -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesSize, err := logic.TestCasesCount(testcasesZip);
	if err != nil {
		log.Println("Error count testcases from form -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	
	uid := ctx.Locals("uid");
	problemID := ctx.Cookies("CreatingProblemID");

	if err := logic.SaveProblem(problemID, problemPDF, testcasesZip); err != nil {
		log.Println("Error save problem to storage -> " + err.Error())
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	row, err := gorm.G[models.Problem](config.DB).Where("problem_id = ? AND author_uid = ?", problemID, uid).Update(ctx, "test_cases_size", testcasesSize);
	if err != nil {
		log.Println("Error update testcases size -> " + err.Error());
	}
	if row == 0 {
		log.Println("Error update testcases size -> cant find creating problem");
	}

	return ctx.SendStatus(fiber.StatusOK);
}