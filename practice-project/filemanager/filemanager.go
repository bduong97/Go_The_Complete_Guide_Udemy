package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("reading the file content failed")
	}
	file.Close()
	return lines, nil
}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}
	encoder :=	json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
	}
	file.Close()
	return nil
}
