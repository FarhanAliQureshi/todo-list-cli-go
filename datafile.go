package main

import (
	"io"
	"os"
)

func ReadDataFile(filename string) (string, error) {
	// If file doesn't exist then create an empty file
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytesData, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytesData), nil
}

func WriteDataFile(filename, data string) error {
	bytesData := []byte(data)
	err := os.WriteFile(filename, bytesData, 0666)
	if err != nil {
		return err
	}

	return nil
}
