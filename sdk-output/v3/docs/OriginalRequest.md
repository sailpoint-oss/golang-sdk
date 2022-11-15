# OriginalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | the account id | [optional] 
**AttributeRequests** | Pointer to [**[]AttributeRequest**](AttributeRequest.md) |  | [optional] 
**Op** | Pointer to **string** | the operation that was used | [optional] 
**Source** | Pointer to [**Source1**](Source1.md) |  | [optional] 

## Methods

### NewOriginalRequest

`func NewOriginalRequest() *OriginalRequest`

NewOriginalRequest instantiates a new OriginalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalRequestWithDefaults

`func NewOriginalRequestWithDefaults() *OriginalRequest`

NewOriginalRequestWithDefaults instantiates a new OriginalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *OriginalRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OriginalRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OriginalRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *OriginalRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAttributeRequests

`func (o *OriginalRequest) GetAttributeRequests() []AttributeRequest`

GetAttributeRequests returns the AttributeRequests field if non-nil, zero value otherwise.

### GetAttributeRequestsOk

`func (o *OriginalRequest) GetAttributeRequestsOk() (*[]AttributeRequest, bool)`

GetAttributeRequestsOk returns a tuple with the AttributeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRequests

`func (o *OriginalRequest) SetAttributeRequests(v []AttributeRequest)`

SetAttributeRequests sets AttributeRequests field to given value.

### HasAttributeRequests

`func (o *OriginalRequest) HasAttributeRequests() bool`

HasAttributeRequests returns a boolean if a field has been set.

### GetOp

`func (o *OriginalRequest) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *OriginalRequest) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *OriginalRequest) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *OriginalRequest) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetSource

`func (o *OriginalRequest) GetSource() Source1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OriginalRequest) GetSourceOk() (*Source1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OriginalRequest) SetSource(v Source1)`

SetSource sets Source field to given value.

### HasSource

`func (o *OriginalRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


