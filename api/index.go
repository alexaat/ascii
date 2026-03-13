package handler

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templates embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method == "GET" {
			formHandler(w)
		} else {
			showError(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		}
	case "/ascii-art":
		if r.Method == "POST" {
			resultHandler(w, r)
		} else {
			showError(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		}
	default:
		showError(w, "404 Unvalid URL", http.StatusNotFound)
	}
}

func formHandler(w http.ResponseWriter) {
	// Render the index.html template
	tmpl, err := template.ParseFS(templates, "templates/index.html")
	if err != nil {
		showError(w, "500 TEMPLATE NOT FOUND", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	// Render the result.html template
	banner := r.FormValue("banner")
	text := r.FormValue("request")
	b, err := readFile(banner + ".txt")
	if err != nil {
		showError(w, "505 BANNER NOT FOUND", http.StatusInternalServerError)
		return
	}
	myMap := parseBanner(b)
	result := printMessageIntoString(text, myMap)
	//err = writeToFile(FilePath, []byte(result))
	// if err != nil {
	// 	showError(w, "500 Cannot write to file", http.StatusInternalServerError)
	// 	return
	// }

	tmpl, err := template.ParseFS(templates, "templates/result.html")
	if err != nil {
		showError(w, "500 TEMPLATE NOT FOUND", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, result)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Render the error.html template
func showError(w http.ResponseWriter, message string, statusCode int) {
	tmpl, err := template.ParseFS(templates, "templates/error.html")
	if err == nil {
		w.WriteHeader(statusCode)
		err := tmpl.Execute(w, message)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "<h1>500 INTERNAL SERVER ERROR</h1>")
			fmt.Fprintf(w, "<h3>500 Template not found</h3>")
			fmt.Fprintf(w, err.Error())
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<h1>500 INTERNAL SERVER ERROR</h1>")
		fmt.Fprintf(w, "<h3>500 Template not found</h3>")
		fmt.Fprintf(w, err.Error())
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
