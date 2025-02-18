package main

import (
	"inkspire/internal/api/books"
	"inkspire/internal/config"
	"inkspire/internal/gen"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    log.Print("Entering application")
    server := new()
    if server == nil {
        log.Fatal("error creating server")
        return
    }
    log.Fatal(server.ListenAndServe())
}

func new() *http.Server {
    cfg := config.GetConfig()

    r := chi.NewRouter()

    r.Use(middleware.Logger)

    bookService := books.NewBookService()

    handler := gen.HandlerFromMux(bookService, r)
    
    return &http.Server{
        Addr: cfg.Port,
        Handler: handler,
    }
}