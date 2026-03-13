package handler

import (
	"archive/zip"
	utils "ascii/utils"
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

//go:embed templates/*.html banners/*.txt
var assets embed.FS

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
	tmpl, err := template.ParseFS(assets, "templates/index.html")
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
	b, err := utils.ReadFile(bannerUrl, assets)
	if err != nil {
		showError(w, "505 BANNER NOT FOUND: "+bannerUrl, http.StatusInternalServerError)
		return
	}
	myMap := utils.ParseBanner(b)
	result := utils.PrintMessageIntoString(text, myMap)

	file, err := os.Create(utils.FilePath)
	if err != nil {
		showError(w, "500 CANNOT WRITE TO FILE", http.StatusInternalServerError)
	}
	defer file.Close()
	file.WriteString(result)

	createZip(w)

	tmpl, err := template.ParseFS(assets, "templates/result.html")
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
	switch format {
	case "zip":
		info, err := os.Stat(utils.ZipFilePath)
		if err != nil {
			size := info.Size()
			w.Header().Set("Content-Length", strconv.Itoa(int(size)))
		}
		w.Header().Set("Content-Disposition", "attachment; filename="+utils.ZipFilePath)
		w.Header().Set("Content-Type", "application/zip")

		http.ServeFile(w, r, utils.ZipFilePath)
	case "text":
		info, err := os.Stat(utils.FilePath)
		if err != nil {
			size := info.Size()
			w.Header().Set("Content-Length", strconv.Itoa(int(size)))
		}
		w.Header().Set("Content-Disposition", "attachment; filename="+utils.FilePath)
		w.Header().Set("Content-Type", "text/plain")
		http.ServeFile(w, r, utils.FilePath)
	}
	w.WriteHeader(http.StatusOK)

}

func createZip(w http.ResponseWriter) {

	archive, err := os.Create(utils.ZipFilePath)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	f1, err := os.Open(utils.FilePath)
	if err != nil {
		showError(w, "500 INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	defer f1.Close()

	w1, err := zipWriter.Create(utils.FilePath)
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

// Render the error.html template
func showError(w http.ResponseWriter, message string, statusCode int) {
	tmpl, err := template.ParseFS(assets, "templates/error.html")
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
