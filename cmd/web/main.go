package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tirzasrwn/gocourse/cmd/pkg/config"
	"github.com/tirzasrwn/gocourse/cmd/pkg/handlers"
	"github.com/tirzasrwn/gocourse/cmd/pkg/render"
)

const portNumber = ":4545"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache.")
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on localhost%s.\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

// run: $go run cmd/web/*.go
