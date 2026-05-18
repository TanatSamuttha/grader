package handler

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v3"
)

func UploadProblem(ctx fiber.Ctx) error {
	problemPDF, err := ctx.FormFile("problem");
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	if err = checkType(problemPDF, "application/pdf"); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcaseZip, err := ctx.FormFile("testcase");
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	if err = checkType(testcaseZip, "application/zip"); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	err = os.MkdirAll("../safe/problems", os.ModePerm)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = os.MkdirAll("../safe/testcases", os.ModePerm)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	err = ctx.SaveFile(problemPDF, "../safe/problems/" + problemPDF.Filename)
	if err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	err = ctx.SaveFile(testcaseZip, "../safe/testcases/" + testcaseZip.Filename)
	if err != nil {
		fmt.Println(err);
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	return ctx.SendStatus(fiber.StatusOK);
}

func checkType(file *multipart.FileHeader, expect string) error {
	result := file.Header.Get("Content-Type");
	if result != expect {
		return errors.New("Invalid Type");
	}
	return nil;
}