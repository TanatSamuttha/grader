package logic

import (
	"archive/zip"
	"bytes"
	"errors"
	"mime/multipart"
	"strings"
)

func TestCasesCount(file *multipart.FileHeader) (uint8, error) {
	src, err := file.Open()
	if err != nil {
		return 0, errors.New("Error open file -> " + err.Error());
	}
	defer src.Close()

	data := make([]byte, file.Size)

	_, err = src.Read(data)
	if err != nil {
		return 0, errors.New("Error read src -> " + err.Error());
	}

	reader, err := zip.NewReader(
		bytes.NewReader(data),
		file.Size,
	)
	if err != nil {
		return 0, errors.New("Error create reader -> " + err.Error());
	}

	var inputCount uint8 = 0
	var outputCount uint8 = 0

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		switch {
		case strings.HasPrefix(file.Name, "testcases/input/"):
			inputCount++

		case strings.HasPrefix(file.Name, "testcases/output/"):
			outputCount++
		}
	}

	if inputCount != outputCount {
		return 0, errors.New("Error testcases count -> input/output mismatch");
	}

	return inputCount, nil
}