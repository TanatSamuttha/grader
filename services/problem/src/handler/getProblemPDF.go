package handler

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
)

func GetProblemPDF(ctx fiber.Ctx) error {
	req, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("STORAGE_SERVER_URL"),
		nil,
	);

	if err != nil {
		log.Println("Error create request -> " + err.Error())
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	// pass header to storage server
	req.Header.Set("Storage-Key", os.Getenv("STORAGE_KEY"),);
	req.Header.Set("problem_id", ctx.Query("id"),);

	client := &http.Client{};

	resp, err := client.Do(req);
	if err != nil {
		log.Println("Error request storage server -> " + err.Error());
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}
	defer resp.Body.Close();

	if resp.StatusCode != http.StatusOK {
		log.Println("Error storage server returned status -> " + resp.Status);
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	// set response headers
	ctx.Set(
		"Content-Type",
		resp.Header.Get("Content-Type"),
	);

	// stream directly to frontend
	_, err = io.Copy(ctx.Response().BodyWriter(), resp.Body)
	if err != nil {
		log.Println("Error stream pdf to frontend -> " + err.Error());
		return ctx.SendStatus(fiber.StatusInternalServerError);
	}

	return nil
}