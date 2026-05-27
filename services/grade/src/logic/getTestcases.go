package logic

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetTestcases(problemID string) ([]string, []string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://storage-server:3002/storage/get/testcases",
		nil,
	)
	if err != nil {
		return nil, nil, errors.New("Error create http request -> " + err.Error());
	}

	req.Header.Set("Storage-Key", os.Getenv("STORAGE_KEY"));
	req.Header.Set("problem_id", problemID);

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, errors.New("Error create http client -> " + err.Error());
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Error bad status code -> %d", resp.StatusCode)
	}

	zipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, errors.New("Error read zip -> " + err.Error());
	}

	reader, err := zip.NewReader(
		bytes.NewReader(zipBytes),
		int64(len(zipBytes)),
	)
	if err != nil {
		return nil, nil, errors.New("Error create reader -> " + err.Error());
	}

	inputs := []string{}
	outputs := []string{}

	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return nil, nil, errors.New("Error read loop -> " + err.Error());
		}

		contentBytes, err := io.ReadAll(rc)
		rc.Close()

		if err != nil {
			return nil, nil, errors.New("Error convert to bytes -> " + err.Error());
		}

		content := string(contentBytes)

		if strings.HasSuffix(file.Name, ".in") {
			inputs = append(inputs, content)
		} else if strings.HasSuffix(file.Name, ".out") {
			outputs = append(outputs, content)
		}
	}

	return inputs, outputs, nil
}