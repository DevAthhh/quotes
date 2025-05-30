package services

import (
	entity "github.com/DevAthhh/quotes/internal/enitity"
	"github.com/DevAthhh/quotes/internal/repository"
)

type QuoteService interface {
	GetQuoteByAuthor(author string) (*entity.Quote, error)
	DeleteQuoteByID(id string) error
	CreateQuote(quote *entity.Quote) error
	GetRandomQuote() *entity.Quote
	GetAllQuotes(author string) []entity.Quote
}

type quoteService struct {
	repo repository.QuoteRepository
}

func (qs *quoteService) GetQuoteByAuthor(author string) (*entity.Quote, error) {
	return qs.repo.GetQuoteByAuthor(author)
}

func (qs *quoteService) DeleteQuoteByID(id string) error {
	return qs.repo.DeleteQuoteByID(id)
}

func (qs *quoteService) CreateQuote(quote *entity.Quote) error {
	return qs.repo.CreateQuote(quote)
}

func (qs *quoteService) GetRandomQuote() *entity.Quote {
	return qs.repo.GetRandomQuote()
}

func (qs *quoteService) GetAllQuotes(author string) []entity.Quote {
	return qs.repo.GetAllQuotes(author)
}

func NewQuoteService(repo repository.QuoteRepository) QuoteService {
	return &quoteService{
		repo: repo,
	}
}
