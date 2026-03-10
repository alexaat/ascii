package handler

import (
	//"os"
)

var (
	FilePath          = "data.txt"
	ErrorTemplatePath = "templates/error.html"
	TemplatePath      = "templates/index.html"
	ZipFilePath       = "archive.zip"
)

// func ReadFile(s string) (string, error) {
// 	data, err := os.ReadFile(s)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(data), nil
// }

// func WriteToFile(fileName string, data []byte) error {
// 	return os.WriteFile(fileName, data, 0644)
// }
