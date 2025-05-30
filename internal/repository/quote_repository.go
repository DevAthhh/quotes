package repository

import (
	"errors"
	"math/rand"
	"strconv"

	entity "github.com/DevAthhh/quotes/internal/enitity"
)

type QuoteRepository interface {
	GetQuoteByAuthor(author string) (*entity.Quote, error)
	DeleteQuoteByID(id string) error
	CreateQuote(quote *entity.Quote) error
	GetRandomQuote() *entity.Quote
	GetAllQuotes(author string) []entity.Quote
}

type quoteRepository struct {
	quotes []entity.Quote
}

func (qr *quoteRepository) GetQuoteByAuthor(author string) (*entity.Quote, error) {
	for _, quote := range qr.quotes {
		if quote.Author == author {
			return &quote, nil
		}
	}
	return nil, errors.New("author not found")
}

func (qr *quoteRepository) DeleteQuoteByID(idQuote string) error {
	intID, err := strconv.Atoi(idQuote)
	if err != nil {
		return err
	}

	for id, quote := range qr.quotes {
		if quote.ID != intID {
			continue
		}
		tmp := qr.quotes[id+1:]
		qr.quotes = qr.quotes[:id]
		qr.quotes = append(qr.quotes, tmp...)
		return nil
	}

	return errors.New("quote not found")
}

func (qr *quoteRepository) CreateQuote(quote *entity.Quote) error {
	if len(qr.quotes) < 1 {
		quote.ID = 1
		qr.quotes = append(qr.quotes, *quote)
	}

	lastID := qr.quotes[len(qr.quotes)-1].ID
	quote.ID = lastID + 1
	qr.quotes = append(qr.quotes, *quote)

	return nil
}

func (qr *quoteRepository) GetRandomQuote() *entity.Quote {
	randomID := rand.Intn(len(qr.quotes))
	return &qr.quotes[randomID]
}

func (qr *quoteRepository) GetAllQuotes(author string) []entity.Quote {
	results := make([]entity.Quote, 0, len(qr.quotes))

	if author != "" {
		for _, quote := range qr.quotes {
			if quote.Author != author {
				continue
			}

			results = append(results, quote)
		}
	} else {
		results = append(results, qr.quotes...)
	}

	return results
}

func NewQuoteRepository() QuoteRepository {
	return &quoteRepository{
		quotes: make([]entity.Quote, 0),
	}
}
