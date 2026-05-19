package handler

import (
	"log"
	"os"
	"storage/logic"

	"github.com/gofiber/fiber/v3"
)

func UploadProblem(ctx fiber.Ctx) error {
	problemPDF, err := ctx.FormFile("problem");
	log.Println("Uploaded -> " + problemPDF.Filename);
	if err != nil {
		log.Println("Error get problem PDF -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	if err := logic.IsPDF(problemPDF); err != nil {
		log.Println("Error check type problem PDF -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesZip, err := ctx.FormFile("testcases");
	if err != nil {
		log.Println("Error get testcases zip -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	if err = logic.IsZip(testcasesZip); err != nil {
		log.Println("Error check type testcases zip -> " + err.Error());
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	err = os.MkdirAll("../safe/problems", os.ModePerm)
	if err != nil {
		log.Println("Error make safe/problems dir -> " + err.Error());
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = os.MkdirAll("../safe/testcases", os.ModePerm)
	if err != nil {
		log.Println("Error make safe/testcases dir -> " + err.Error());
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = ctx.SaveFile(problemPDF, "../safe/problems/" + problemPDF.Filename)
	if err != nil {
		log.Println("Error save problem PDF -> " + err.Error());
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	err = ctx.SaveFile(testcasesZip, "../safe/testcases/" + testcasesZip.Filename)
	if err != nil {
		log.Println("Error save testcases zip -> " + err.Error());
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	return ctx.SendStatus(fiber.StatusOK);
}