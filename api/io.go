package main

import (
	"os"
)

var (
	filePath          = "data.txt"
	errorTemplatePath = "templates/error.html"
	templatePath      = "templates/index.html"
	zipFilePath       = "archive.zip"
)

func readFile(s string) (string, error) {
	data, err := os.ReadFile(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeToFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}
