package handler

import (
	"embed"
	"os"
)

var (
	FilePath    = "data.txt"
	ZipFilePath = "archive.zip"
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
