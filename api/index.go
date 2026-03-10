package handler

import (
	//"archive/zip"
	"fmt"
	//"io"
	"net/http"
	//"os"
	//"strconv"
	"embed"
	"html/template"
)

var (
	filePath          = "data.txt"
	errorTemplatePath = "templates/error.html"
	templatePath      = "templates/index.html"
	zipFilePath       = "archive.zip"
)

var templateFS embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		formHandler(w, r)
	}

	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templateFS, "../templates/index.html"))

	// Render the index.html template
	// t, err := template.ParseFiles(templatePath)
	// if err != nil {
	// 	showError(w, "404 TEMPLATE NOT FOUND: "+err.Error(), http.StatusNotFound)
	// 	return
	// }
	err := t.Execute(w, nil)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

// Render the error.html template
func showError(w http.ResponseWriter, message string, statusCode int) {
	fmt.Fprintf(w, "<h1>ERROR </h1>")
	fmt.Fprintf(w, message)
	// t, err := template.ParseFiles(errorTemplatePath)
	// if err == nil {
	// 	w.WriteHeader(statusCode)
	// 	t.Execute(w, message)
	// }
}
