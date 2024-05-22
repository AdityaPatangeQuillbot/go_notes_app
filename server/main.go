package main

import (
	"fmt"
	"net/http"

	"github.com/AdityaPatangeQuillBot/go_notes_app/server/core"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if name == "" {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", core.DEFAULT_PORT), nil)
}
