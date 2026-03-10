// Ascii-art-web
// We are going to use http library as the name implies, it's one that deals with the web.
// Golang was built from the ground up to be aware of the web to be able to deal with it in a meaningful fashion.
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

// We need to create a request handler. This is a func. that going to be used for every request that's made to our server.
func main() {
	// Handling the image
	fileServer := http.FileServer(http.Dir("./images"))

	// This tell HTTP that when a request is made we want to use our handler
	// "/" - Path name
	// http.HandleFunc is a part of the http package. I am going to give it a path name. ("/") That's the URL.
	// URL - uniform resource locator that I want to listen to.
	// http.ResponseWriter the first thing that you need to have because this is a handler function for the http.
	// :8080 common alternative HTTP port used for web traffic.
	const portNumber = ":8080"
	http.Handle("/images/", http.StripPrefix("/images", fileServer))
	// Routing
	http.HandleFunc("/", formHandler)
	// Routing
	http.HandleFunc("/ascii-art", resultHandler)
	http.HandleFunc("/download", downloadHandler)
	// Then we need to tell HTTP to listen and serve on port 8080
	fmt.Printf("Starting application on port %s\n", portNumber)
	// Web server that listens for requests. (Without it, the main func never executes)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("\nCannot start server")
	}
}

// Render template
// It doesn't return anything because it writes everything to the response writer.
// . dot means root level of my application (inside ascii-art-web folder)
// html is an argument of the template I want to pass as a string
//
//	func renderTemplate(w http.ResponseWriter, html string) {
//		parsedTemplate, _ := template.ParseFiles("./templates/" + html)
//		err :=parsedTemplate.Execute(w, nil)
//		if err != nil (
//			fmt.Println("error parsing template", err)
//			return
//		)
//	}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path != "/" {
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

func resultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		banner := r.FormValue("banner")
		text := r.FormValue("request")
		b, err := readFile(banner + ".txt")
		if err != nil {
			showError(w, "404 BANNER NOT FOUND", http.StatusNotFound)
			// return here will stop execution this function
			return
		}
		myMap := parseBanner(b)
		result := printMessageIntoString(text, myMap)
		err = writeToFile(filePath, []byte(result))
		if err != nil {
			showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
			return
		}
		t, err := template.ParseFiles("templates/result.html")
		if err != nil {
			showError(w, "404 TEMPLATE NOT FOUND", http.StatusNotFound)
			return
		}
		err = t.Execute(w, result)
		if err != nil {
			showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
			return
		}

		return
	}
	if r.Method == "GET" {
		showError(w, "400 BAD REQUEST", http.StatusBadRequest)
		return
	}
	showError(w, "400 BAD REQUEST", http.StatusBadRequest)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		format := r.FormValue("format")

		if format == "zip" {
			createZip(w)
			data, err := readFile(zipFilePath)
			if err != nil {
				showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
				return
			}
			contentLength := len([]byte(data))
			w.Header().Set("Content-Disposition", "attachment; filename="+zipFilePath)
			w.Header().Set("Content-Type", "application/zip")
			w.Header().Set("Content-Length", strconv.Itoa(contentLength))
			http.ServeFile(w, r, zipFilePath)
			return
		}

		if format == "text" {
			data, err := readFile(filePath)
			if err != nil {
				showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
				return
			}
			contentLength := len([]byte(data))
			w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", strconv.Itoa(contentLength))
			http.ServeFile(w, r, filePath)
			return
		}
		// If forrmat is not "text" nor "zip" then show error
		showError(w, "404 FORMAT NOT FOUND", http.StatusNotFound)

	}

	showError(w, "400 BAD REQUEST", http.StatusBadRequest)
}

// Render the error.html template
func showError(w http.ResponseWriter, message string, statusCode int) {
	t, err := template.ParseFiles(errorTemplatePath)
	if err == nil {
		w.WriteHeader(statusCode)
		t.Execute(w, message)
	}
}

func createZip(w http.ResponseWriter) {
	archive, err := os.Create(zipFilePath)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	f1, err := os.Open(filePath)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	defer f1.Close()

	w1, err := zipWriter.Create(filePath)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(w1, f1); err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	zipWriter.Close()
}
