package handler

import "github.com/gofiber/fiber/v3"

func GetProblem(ctx fiber.Ctx) error {
	problemID := ctx.Get("problem_id");
	filePath := "../safe/problems/" + problemID + ".pdf";
	return ctx.SendFile(filePath);
}