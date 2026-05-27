package logic

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type testcase struct {
	Input  string
	Output string
}

func GetTestcases(problemID string) ([]string, []string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://storage-server:3002/storage/get/testcases",
		nil,
	)
	if err != nil {
		return nil, nil, errors.New("Error create http request -> " + err.Error())
	}

	req.Header.Set("Storage-Key", os.Getenv("STORAGE_KEY"))
	req.Header.Set("problem_id", problemID)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, errors.New("Error create http client -> " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Error bad status code -> %d", resp.StatusCode)
	}

	zipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, errors.New("Error read zip -> " + err.Error())
	}

	reader, err := zip.NewReader(
		bytes.NewReader(zipBytes),
		int64(len(zipBytes)),
	)
	if err != nil {
		return nil, nil, errors.New("Error create reader -> " + err.Error())
	}

	testcases := map[int]*testcase{}

	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return nil, nil, errors.New("Error read loop -> " + err.Error())
		}

		contentBytes, err := io.ReadAll(rc)
		rc.Close()

		if err != nil {
			return nil, nil, errors.New("Error convert to bytes -> " + err.Error())
		}

		content := string(contentBytes)

		base := filepath.Base(file.Name)

		ext := filepath.Ext(base)
		numberStr := strings.TrimSuffix(base, ext)

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			continue
		}

		if _, exists := testcases[number]; !exists {
			testcases[number] = &testcase{}
		}

		if ext == ".in" {
			testcases[number].Input = content
		} else if ext == ".out" {
			testcases[number].Output = content
		}
	}

	keys := make([]int, 0, len(testcases))

	for k := range testcases {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	inputs := []string{}
	outputs := []string{}

	for _, k := range keys {
		inputs = append(inputs, testcases[k].Input)
		outputs = append(outputs, testcases[k].Output)
	}

	return inputs, outputs, nil
}