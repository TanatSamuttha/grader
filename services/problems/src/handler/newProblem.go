package handler

import (
	"fmt"
	"problems/models"

	"github.com/gofiber/fiber/v3"
)

func NewProblem(ctx fiber.Ctx) error {
	var problem models.ProblemDTO;
	if err := ctx.Bind().Body(&problem); err != nil {
		return ctx.SendStatus(400);
	}
	fmt.Println(problem);

	return ctx.SendStatus(200);
}