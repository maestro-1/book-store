package handlers

import (
	"net/http"
	"go-book-sales-api/book_sales_api"
)

type BookHandler struct {
	repo *book_sales_api.Queries
}

func NewBookHandler(repo *book_sales_api.Queries) *BookHandler {
	return &BookHandler{
		repo: repo,
	}
}

func (s *BookHandler) SellBook(w http.ResponseWriter, r *http.Request) error {
	return nil
}

