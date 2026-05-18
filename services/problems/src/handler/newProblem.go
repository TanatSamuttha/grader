package handler

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"problems/models"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func NewProblem(ctx fiber.Ctx) error {
	var problemMeta models.ProblemDTO;
	if err := ctx.Bind().Body(&problemMeta); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}
	fmt.Println(problemMeta);

	author := ctx.Locals("uid").(string);
	problemID := uuid.New();

	problemPDF := ctx.FormFile("problem");
	testcasesZip := ctx.FormFile("testcases");
	testcasesSize, err := testCasesCount(testcasesZip);
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	// problem := models.Problem{
	// 	ProblemID: problemID.String(),
	// 	Name: problemMeta.Name,
	// 	AuthorUID: author,
	// 	TestCasesSize: ,
	// };

	return ctx.SendStatus(200);
}

func testCasesCount(file *multipart.FileHeader) (int, error) {
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

	inputCount := 0;
	outputCount := 0;

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