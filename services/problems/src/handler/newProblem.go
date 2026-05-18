package handler

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"problems/config"
	"problems/models"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewProblem(ctx fiber.Ctx) error {
	var problemMeta models.ProblemDTO;
	if err := ctx.Bind().Body(&problemMeta); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	fmt.Println(problemMeta);

	author := ctx.Locals("uid").(string);
	problemID := uuid.New();

	problemPDF, err := ctx.FormFile("problem");
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesZip, err := ctx.FormFile("testcases");
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	testcasesSize, err := testCasesCount(testcasesZip);
	if err != nil {
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
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK);
}

func testCasesCount(file *multipart.FileHeader) (uint8, error) {
	src, err := file.Open();
	if err != nil {
		return 0, err;
	}
	defer src.Close();

	data := make([]byte, file.Size);

	_, err = src.Read(data);
	if err != nil {
		return 0, err;
	}

	reader, err := zip.NewReader(
		bytes.NewReader(data),
		file.Size,
	);
	if err != nil {
		return 0, err;
	}

	var inputCount uint8 = 0;
	var outputCount uint8 = 0;

	for _, file := range reader.File {
		if file.FileInfo().IsDir(){
			continue;
		}

		switch {
			case strings.HasPrefix(file.Name, "testcases/input/"):
				inputCount++;

			case strings.HasPrefix(file.Name, "testcases/output/"):
				outputCount++;
		}
	}

	if inputCount != outputCount {
		return 0, errors.New("input/output mismatch");
	}

	return inputCount, nil;
}