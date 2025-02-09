# \BooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBook**](BooksAPI.md#AddBook) | **Post** /books | Add a book
[**DeleteBook**](BooksAPI.md#DeleteBook) | **Delete** /books/{bookId} | Delete a book
[**GetBook**](BooksAPI.md#GetBook) | **Get** /books/{bookId} | Get a specific book
[**GetBooks**](BooksAPI.md#GetBooks) | **Get** /books | Get all books
[**ReplaceBook**](BooksAPI.md#ReplaceBook) | **Put** /books/{bookId} | Replace a book
[**UpdateBook**](BooksAPI.md#UpdateBook) | **Patch** /books/{bookId} | Update a book



## AddBook

> CreatedResponse AddBook(ctx).NewBook(newBook).Execute()

Add a book



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	newBook := *openapiclient.NewNewBook() // NewBook | Create a new book

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.AddBook(context.Background()).NewBook(newBook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.AddBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBook`: CreatedResponse
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.AddBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newBook** | [**NewBook**](NewBook.md) | Create a new book | 

### Return type

[**CreatedResponse**](CreatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBook

> GetBooks200Response DeleteBook(ctx, bookId).Execute()

Delete a book



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookId := "550e8400-e29b-41d4-a716-446655440000" // string | The ID of the book.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.DeleteBook(context.Background(), bookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.DeleteBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBook`: GetBooks200Response
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.DeleteBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **string** | The ID of the book. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBooks200Response**](GetBooks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBook

> GetBooks200Response GetBook(ctx, bookId).Execute()

Get a specific book



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookId := "550e8400-e29b-41d4-a716-446655440000" // string | The ID of the book.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBook(context.Background(), bookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBook`: GetBooks200Response
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **string** | The ID of the book. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetBooks200Response**](GetBooks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBooks

> GetBooks200Response GetBooks(ctx).Limit(limit).Offset(offset).Genre(genre).Status(status).Author(author).Sort(sort).Execute()

Get all books



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int32(10) // int32 | Limit the number of results returned. (optional) (default to 10)
	offset := int32(0) // int32 | Offset the start of the result set. (optional) (default to 0)
	genre := []string{"Inner_example"} // []string | Filter books by multiple genres. (optional)
	status := "read" // string | Filter books by reading status. (optional)
	author := "Brandon Sanderson" // string | Filter books by author name. (optional)
	sort := "title" // string | Sort books by a specific field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBooks(context.Background()).Limit(limit).Offset(offset).Genre(genre).Status(status).Author(author).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBooks`: GetBooks200Response
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of results returned. | [default to 10]
 **offset** | **int32** | Offset the start of the result set. | [default to 0]
 **genre** | **[]string** | Filter books by multiple genres. | 
 **status** | **string** | Filter books by reading status. | 
 **author** | **string** | Filter books by author name. | 
 **sort** | **string** | Sort books by a specific field. | 

### Return type

[**GetBooks200Response**](GetBooks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBook

> GetBooks200Response ReplaceBook(ctx, bookId).Book(book).Execute()

Replace a book



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookId := "550e8400-e29b-41d4-a716-446655440000" // string | The ID of the book.
	book := *openapiclient.NewBook("550e8400-e29b-41d4-a716-446655440000", "The Lusty Argonian Maid", "1a3e4567-e89b-12d3-a456-426614174000") // Book | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.ReplaceBook(context.Background(), bookId).Book(book).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.ReplaceBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceBook`: GetBooks200Response
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.ReplaceBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **string** | The ID of the book. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **book** | [**Book**](Book.md) |  | 

### Return type

[**GetBooks200Response**](GetBooks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBook

> GetBooks200Response UpdateBook(ctx, bookId).Book(book).Execute()

Update a book



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookId := "550e8400-e29b-41d4-a716-446655440000" // string | The ID of the book.
	book := *openapiclient.NewBook("550e8400-e29b-41d4-a716-446655440000", "The Lusty Argonian Maid", "1a3e4567-e89b-12d3-a456-426614174000") // Book | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.UpdateBook(context.Background(), bookId).Book(book).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.UpdateBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBook`: GetBooks200Response
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.UpdateBook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookId** | **string** | The ID of the book. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **book** | [**Book**](Book.md) |  | 

### Return type

[**GetBooks200Response**](GetBooks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

