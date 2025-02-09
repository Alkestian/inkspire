// filepath: internal/api/books.go
package api

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "inkspire/pkg/models"
)

type BooksAPI struct {
    BooksAPIService *models.BooksAPIService
    books           map[string]*models.Book
}

func NewBooksAPI() *BooksAPI {
    config := models.NewConfiguration()
    client := models.NewAPIClient(config)

    return &BooksAPI{
        BooksAPIService: client.BooksAPI,
        books:           make(map[string]*models.Book),
    }
}

func (api *BooksAPI) AddBook(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    var newBook models.NewBook

    if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    resp, _, err := api.BooksAPIService.AddBook(ctx).NewBook(newBook).Execute()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (api *BooksAPI) GetBooks(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    // Extract query parameters (limit, offset, genre, status, etc.) from the request.
    limitStr := r.URL.Query().Get("limit")
    offsetStr := r.URL.Query().Get("offset")
    //genre := r.URL.Query()["genre"] // Multiple genres
    status := r.URL.Query().Get("status")
    author := r.URL.Query().Get("author")
    sort := r.URL.Query().Get("sort")

    // Convert limit and offset to integers (with error handling).
    var limit int32 = 10 // Default value
    if limitStr != "" {
        l, err := models.Atoi(limitStr)
        if err != nil {
            http.Error(w, "Invalid limit value", http.StatusBadRequest)
            return
        }
        limit = int32(l)
    }

    var offset int32 = 0 // Default value
    if offsetStr != "" {
        o, err := models.Atoi(offsetStr)
        if err != nil {
            http.Error(w, "Invalid offset value", http.StatusBadRequest)
            return
        }
        offset = int32(o)
    }

    // Call the generated API client's GetBooks method.
    resp, _, err := api.BooksAPIService.GetBooks(ctx).
        Limit(limit).
        Offset(offset).
        Status(status).
        Author(author).
        Sort(sort).
        Execute()

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the response as JSON.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func (api *BooksAPI) Routes() http.Handler {
    r := chi.NewRouter()
    r.Post("/", api.AddBook)
    r.Get("/", api.GetBooks) // Add the GetBooks route
    // Add other routes here
    return r
}