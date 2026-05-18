package handler

import (
	"problems/models"

	"github.com/gofiber/fiber/v3"
)

func NewProblem(ctx fiber.Ctx) error {
	var problem models.Problem;
	if err := ctx.Bind().Body(&problem); err != nil {
		return ctx.SendStatus(400);
	}

	return ctx.SendStatus(200);
}