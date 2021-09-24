package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":4545"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on localhost%s.", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// run: go run *
