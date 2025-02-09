# GetBooks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Book**](Book.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewGetBooks200Response

`func NewGetBooks200Response() *GetBooks200Response`

NewGetBooks200Response instantiates a new GetBooks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBooks200ResponseWithDefaults

`func NewGetBooks200ResponseWithDefaults() *GetBooks200Response`

NewGetBooks200ResponseWithDefaults instantiates a new GetBooks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBooks200Response) GetData() []Book`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBooks200Response) GetDataOk() (*[]Book, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBooks200Response) SetData(v []Book)`

SetData sets Data field to given value.

### HasData

`func (o *GetBooks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetadata

`func (o *GetBooks200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetBooks200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetBooks200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetBooks200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


