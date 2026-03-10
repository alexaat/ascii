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

	fmt.Printf("Starting application on port %s\n", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("\nCannot start server")
	}
	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
