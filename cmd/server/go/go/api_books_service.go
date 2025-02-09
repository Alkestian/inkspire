// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Inkspire
 *
 * This app allows for making lists of various types across different media and will eventually have fuzzy searching and rating functionality added to it. 
 *
 * API version: 0.0.1
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
	"reflect"
)

// BooksAPIService is a service that implements the logic for the BooksAPIServicer
// This service should implement the business logic for every endpoint for the BooksAPI API.
// Include any external packages or services that will be required by this service.
type BooksAPIService struct {
}

// NewBooksAPIService creates a default api service
func NewBooksAPIService() *BooksAPIService {
	return &BooksAPIService{}
}

// GetBooks - Get all books
func (s *BooksAPIService) GetBooks(ctx context.Context, limit int32, offset int32, genre []string, status string, author string, sort string) (ImplResponse, error) {
	// TODO - update GetBooks with the required logic for this service method.
	// Add api_books_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetBooks200Response{}) or use other options such as http.Ok ...
	// return Response(200, GetBooks200Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetBooks method not implemented")
}

// AddBook - Add a book
func (s *BooksAPIService) AddBook(ctx context.Context, newBook NewBook) (ImplResponse, error) {
	// TODO - update AddBook with the required logic for this service method.
	// Add api_books_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, CreatedResponse{}) or use other options such as http.Ok ...
	// return Response(201, CreatedResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddBook method not implemented")
}

// GetBook - Get a specific book
func (s *BooksAPIService) GetBook(ctx context.Context, bookId string) (ImplResponse, error) {
	// TODO - update GetBook with the required logic for this service method.
	// Add api_books_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetBooks200Response{}) or use other options such as http.Ok ...
	// return Response(200, GetBooks200Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetBook method not implemented")
}

// ReplaceBook - Replace a book
func (s *BooksAPIService) ReplaceBook(ctx context.Context, bookId string, book Book) (ImplResponse, error) {
	// TODO - update ReplaceBook with the required logic for this service method.
	// Add api_books_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetBooks200Response{}) or use other options such as http.Ok ...
	// return Response(200, GetBooks200Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ReplaceBook method not implemented")
}

// DeleteBook - Delete a book
func (s *BooksAPIService) DeleteBook(ctx context.Context, bookId string) (ImplResponse, error) {
	// TODO - update DeleteBook with the required logic for this service method.
	// Add api_books_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetBooks200Response{}) or use other options such as http.Ok ...
	// return Response(200, GetBooks200Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteBook method not implemented")
}

// UpdateBook - Update a book
func (s *BooksAPIService) UpdateBook(ctx context.Context, bookId string, book Book) (ImplResponse, error) {
	// TODO - update UpdateBook with the required logic for this service method.
	// Add api_books_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetBooks200Response{}) or use other options such as http.Ok ...
	// return Response(200, GetBooks200Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, Problem{}) or use other options such as http.Ok ...
	// return Response(400, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(401, Problem{}) or use other options such as http.Ok ...
	// return Response(401, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(403, Problem{}) or use other options such as http.Ok ...
	// return Response(403, Problem{}), nil

	// TODO: Uncomment the next line to return response Response(404, Problem{}) or use other options such as http.Ok ...
	// return Response(404, Problem{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateBook method not implemented")
}
