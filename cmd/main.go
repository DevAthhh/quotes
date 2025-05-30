package main

import (
	"log"

	"github.com/DevAthhh/quotes/internal/controllers"
	"github.com/DevAthhh/quotes/internal/http"
	"github.com/DevAthhh/quotes/internal/http/handler"
	"github.com/DevAthhh/quotes/internal/repository"
	"github.com/DevAthhh/quotes/internal/services"
)

func main() {
	repo := repository.NewQuoteRepository()
	service := services.NewQuoteService(repo)
	ctrl := controllers.NewController(service)

	routes := handler.InitRoutes(ctrl)

	server := http.NewServer("8080", routes)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
