package handler

import (
	"grade/config"
	"grade/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func SubmissionHandler(ctx fiber.Ctx) error {
	id := uuid.New().String();
	uid := ctx.Get("uid");
	problemID := ctx.Get("problem_id");

	var codeDTO models.CodeDTO;
	err := ctx.Bind().Body(&codeDTO);
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest);
	}

	code := codeDTO.Code;
	lang := codeDTO.Lang;
	
	job := models.Job{
		ID: id,
		UID: uid,
		ProblemID: problemID,
		Code: code,
		Lang: lang,
	};

	config.CallWorker(job);

	return ctx.SendStatus(fiber.StatusOK);
}