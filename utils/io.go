package handler

import (
	"embed"
	"os"
)

var (
	filePath          = "data.txt"
	errorTemplatePath = "templates/error.html"
	templatePath      = "templates/index.html"
	zipFilePath       = "archive.zip"
)

func ReadFile(s string, banners embed.FS) (string, error) {
	data, err := banners.ReadFile(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteToFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, 0644)
}
