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

	if r.URL.Path == "/" {
		formHandler(w,r)
	}


	


	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
