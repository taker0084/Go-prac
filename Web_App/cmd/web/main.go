package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portNumber = ":8080"

//main is the main application function
func main(){
	var app config.AppConfig

	tc,err :=render.CreateTemplateCache()
	if err != nil{
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	//make new instance of App
	repo := handlers.NewRepo(&app)
	//send AppConfig instance back to handlers
	handlers.NewHandlers(repo)

	//send AppConfig instance back to the render
	render.NewTemplates(&app)

	http.HandleFunc("/",handlers.Repo.Home)
	http.HandleFunc("/about",handlers.Repo.About)

	fmt.Printf("Starting application on port %s \n", portNumber)
	_=http.ListenAndServe(portNumber, nil)
}