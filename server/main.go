package main

import (
	"fmt"
	"main/core"
	"main/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/signup", handlers.UserSignup)
	http.ListenAndServe(fmt.Sprintf(":%s", core.DEFAULT_PORT), nil)
}
