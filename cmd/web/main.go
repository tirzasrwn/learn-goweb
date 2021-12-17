package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tirzasrwn/learn-goweb/cmd/pkg/config"
	"github.com/tirzasrwn/learn-goweb/cmd/pkg/handlers"
	"github.com/tirzasrwn/learn-goweb/cmd/pkg/render"
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

	fmt.Printf("Starting application on localhost%s.\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}

// run: $go run cmd/web/*.go
