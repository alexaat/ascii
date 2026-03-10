package handler

import (
	//"archive/zip"
	"fmt"
	//"io"
	"net/http"
	//"os"
	//"strconv"
	//"text/template"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("./images"))
	const portNumber = ":8080"
	http.Handle("/images/", http.StripPrefix("/images", fileServer))

	http.HandleFunc("/", formHandler)
	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
