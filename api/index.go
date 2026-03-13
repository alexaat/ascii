package handler

import (
	utils "ascii/utils"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:embed templates/*.html banners/*.txt
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
	case "/download":
		if r.Method == "POST" {
			downloadHandler(w, r)
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
	bannerUrl := "banners/" + banner + ".txt"
	b, err := utils.ReadFile(bannerUrl, templates)
	if err != nil {
		showError(w, "505 BANNER NOT FOUND: "+bannerUrl, http.StatusInternalServerError)
		return
	}
	myMap := utils.ParseBanner(b)
	result := utils.PrintMessageIntoString(text, myMap)

	file, err := os.Create("/tmp/result.txt")
	if err != nil {
		showError(w, "500 Cannot write to file "+err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()
	file.WriteString(result)

	// err = utils.WriteToFile(utils.FilePath, []byte(result))
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

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	format := r.FormValue("format")
	if format == "zip" {
		fmt.Fprintf(w, "<h3>Download as zip</h3>")
		return
	}
	if format == "text" {
		filePath := "/tmp/result.txt"
		//fmt.Fprintf(w, "<h3>Download as text</h3>")
		w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
		w.Header().Set("Content-Type", "text/plain")
		http.ServeFile(w, r, filePath)
	}

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

// func createZip(w http.ResponseWriter) {
// 	archive, err := os.Create(ZipFilePath)
// 	if err != nil {
// 		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
// 		return
// 	}
// 	defer archive.Close()

// 	zipWriter := zip.NewWriter(archive)

// 	f1, err := os.Open(FilePath)
// 	if err != nil {
// 		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
// 		return
// 	}
// 	defer f1.Close()

// 	w1, err := zipWriter.Create(FilePath)
// 	if err != nil {
// 		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
// 		return
// 	}

// 	if _, err := io.Copy(w1, f1); err != nil {
// 		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
// 		return
// 	}

// 	zipWriter.Close()
// }

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
