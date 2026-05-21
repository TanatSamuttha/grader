package handler

import "github.com/gofiber/fiber/v3"

func SubmissionHandler(ctx fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK);
}