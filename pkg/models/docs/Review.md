# Review

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**BookId** | **string** |  | 
**Rating** | Pointer to **int32** |  | [optional] 
**ReviewText** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewReview

`func NewReview(id string, bookId string, createdAt time.Time, ) *Review`

NewReview instantiates a new Review object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewWithDefaults

`func NewReviewWithDefaults() *Review`

NewReviewWithDefaults instantiates a new Review object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Review) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Review) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Review) SetId(v string)`

SetId sets Id field to given value.


### GetBookId

`func (o *Review) GetBookId() string`

GetBookId returns the BookId field if non-nil, zero value otherwise.

### GetBookIdOk

`func (o *Review) GetBookIdOk() (*string, bool)`

GetBookIdOk returns a tuple with the BookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookId

`func (o *Review) SetBookId(v string)`

SetBookId sets BookId field to given value.


### GetRating

`func (o *Review) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Review) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Review) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Review) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetReviewText

`func (o *Review) GetReviewText() string`

GetReviewText returns the ReviewText field if non-nil, zero value otherwise.

### GetReviewTextOk

`func (o *Review) GetReviewTextOk() (*string, bool)`

GetReviewTextOk returns a tuple with the ReviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewText

`func (o *Review) SetReviewText(v string)`

SetReviewText sets ReviewText field to given value.

### HasReviewText

`func (o *Review) HasReviewText() bool`

HasReviewText returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Review) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Review) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Review) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


