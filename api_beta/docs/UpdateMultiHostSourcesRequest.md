# UpdateMultiHostSourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | Pointer to [**[]JsonPatchOperation**](JsonPatchOperation.md) | Operations to be applied. | [optional] 

## Methods

### NewUpdateMultiHostSourcesRequest

`func NewUpdateMultiHostSourcesRequest() *UpdateMultiHostSourcesRequest`

NewUpdateMultiHostSourcesRequest instantiates a new UpdateMultiHostSourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMultiHostSourcesRequestWithDefaults

`func NewUpdateMultiHostSourcesRequestWithDefaults() *UpdateMultiHostSourcesRequest`

NewUpdateMultiHostSourcesRequestWithDefaults instantiates a new UpdateMultiHostSourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *UpdateMultiHostSourcesRequest) GetOperations() []JsonPatchOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *UpdateMultiHostSourcesRequest) GetOperationsOk() (*[]JsonPatchOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *UpdateMultiHostSourcesRequest) SetOperations(v []JsonPatchOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *UpdateMultiHostSourcesRequest) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


