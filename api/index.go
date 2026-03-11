package handler

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/index.html
var templates embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFS(templates, "templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Go on Vercel",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

/*
package handler

import (
	//"archive/zip"
	"fmt"
	//"io"
	"net/http"
	//"os"
	//"strconv"
	"text/template"
)

var (
	filePath          = "data.txt"
	errorTemplatePath = "templates/error.html"
	templatePath      = "/ascii/templates/index.html"
	zipFilePath       = "archive.zip"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		formHandler(w, r)
	}

	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	// Render the index.html template
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		showError(w, "404 TEMPLATE NOT FOUND: "+err.Error(), http.StatusNotFound)
		return
	}
	err = t.ExecuteTemplate(w, "base", nil)
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
*/
