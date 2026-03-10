package handler

import (
	//"archive/zip"
	//"fmt"
	//"io"
	"net/http"
	//"os"
	//"strconv"
	"text/template"
)

var (
	filePath          = "data.txt"
	errorTemplatePath = "templates/error.html"
	templatePath      = "templates/index.html"
	zipFilePath       = "archive.zip"
)


func Handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		formHandler(w,r)
	}


	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		showError(w, "400 BAD REQUEST", http.StatusBadRequest)
		// return here will stop execution this function
		return
	}
	// Render the index.html template
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		showError(w, "404 TEMPLATE NOT FOUND", http.StatusNotFound)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
}

// Render the error.html template
func showError(w http.ResponseWriter, message string, statusCode int) {
	t, err := template.ParseFiles(errorTemplatePath)
	if err == nil {
		w.WriteHeader(statusCode)
		t.Execute(w, message)
	}
}
