package handler

import (
	"net/http"

	"github.com/DevAthhh/quotes/internal/controllers"
	"github.com/gorilla/mux"
)

func InitRoutes(ctrl *controllers.Controller) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/quotes", ctrl.CreateHandle).Methods("POST")

	r.HandleFunc("/quotes", ctrl.GetQuote).Methods("GET")

	r.HandleFunc("/quotes/random", ctrl.GetRandomQuote).Methods("GET")

	r.HandleFunc("/quotes/{id}", ctrl.DeleteQuote).Methods("DELETE")

	return r
}
