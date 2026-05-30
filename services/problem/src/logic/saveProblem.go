package logic

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func SaveProblem(problemID string, problemPDF *multipart.FileHeader, testcasesZip *multipart.FileHeader) error {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// ---------- PDF ----------
	pdfFile, err := problemPDF.Open()
	if err != nil {
		return err
	}
	defer pdfFile.Close()

	pdfPart, err := writer.CreateFormFile("problem", problemID + ".pdf")
	if err != nil {
		return err
	}

	if _, err := io.Copy(pdfPart, pdfFile); err != nil {
		return err
	}

	// ---------- ZIP ----------
	zipFile, err := testcasesZip.Open()
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipPart, err := writer.CreateFormFile("testcases", problemID + ".zip")
	if err != nil {
		return err
	}

	if _, err := io.Copy(zipPart, zipFile); err != nil {
		return err
	}

	// finalize multipart
	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("STORAGE_SERVER_URL"),
		&body,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Storage-Key", os.Getenv("STORAGE_KEY"))

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return errors.New("storage server error: " + string(b))
	}

	return nil
}