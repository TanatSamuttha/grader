package handler

import "github.com/gofiber/fiber/v3"

func GetTestcases(ctx fiber.Ctx) error {
	problemID := ctx.Get("problem_id");
	filePath := "../safe/testcases/" + problemID + ".zip";
	return ctx.SendFile(filePath);
}