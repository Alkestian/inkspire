package books

import (
	"inkspire/internal/gen"
	"net/http"
)

type BookService struct {
}

func (b BookService) GetBooks(w http.ResponseWriter, r *http.Request, params gen.GetBooksParams) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GetBooks endpoint hit"))
}

func (b BookService) AddBook(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement the logic to add a book
	w.WriteHeader(http.StatusNotImplemented)
}

func (b BookService) DeleteBook(w http.ResponseWriter, r *http.Request, bookId gen.BookId) {
	// TODO: Implement the logic to delete a book
	w.WriteHeader(http.StatusNotImplemented)
}

func (b BookService) GetBook(w http.ResponseWriter, r *http.Request, bookId gen.BookId) {
	// TODO: Implement the logic to get a book
	w.WriteHeader(http.StatusNotImplemented)
}

func (b BookService) UpdateBook(w http.ResponseWriter, r *http.Request, bookId gen.BookId) {
	// TODO: Implement the logic to update a book
	w.WriteHeader(http.StatusNotImplemented)
}

func (b BookService) ReplaceBook(w http.ResponseWriter, r *http.Request, bookId gen.BookId) {
	// TODO: Implement the logic to replace a book
	w.WriteHeader(http.StatusNotImplemented)
}

func NewBookService() BookService {
	return BookService{}
}