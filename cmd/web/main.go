package main

import (
	"fmt"
	"net/http"

	"github.com/tirzasrwn/gocourse/cmd/pkg/handlers"
)

const portNumber = ":4545"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on localhost%s.\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

// run: $go run cmd/web/*.go
