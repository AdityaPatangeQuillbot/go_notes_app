package main

import (
	"fmt"
	"main/core"
	"main/handlers"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Path[1:]
	if name == "" {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/signup", handlers.UserSignup)
	http.ListenAndServe(fmt.Sprintf(":%s", core.DEFAULT_PORT), nil)
}
