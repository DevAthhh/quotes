package controllers

import (
	"encoding/json"
	"net/http"

	entity "github.com/DevAthhh/quotes/internal/enitity"
	"github.com/DevAthhh/quotes/internal/services"
	"github.com/gorilla/mux"
)

type Controller struct {
	quoteService services.QuoteService
}

func (c *Controller) CreateHandle(w http.ResponseWriter, r *http.Request) {
	var quote entity.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.quoteService.CreateQuote(&quote); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"status\": \"quote has been created\"}"))

}

func (c *Controller) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	randomQuote := c.quoteService.GetRandomQuote()
	byteRandomQuote, err := json.Marshal(randomQuote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteRandomQuote)
}

func (c *Controller) GetQuote(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	author := query.Get("author")

	results := c.quoteService.GetAllQuotes(author)

	byteQuotes, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteQuotes)
}

func (c *Controller) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.quoteService.DeleteQuoteByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"status\": \"quote has been deleted\"}"))
}

func NewController(quoteService services.QuoteService) *Controller {
	return &Controller{
		quoteService: quoteService,
	}
}
