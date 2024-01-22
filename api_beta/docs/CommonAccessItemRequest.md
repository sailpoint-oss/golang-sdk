# CommonAccessItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**CommonAccessItemAccess**](CommonAccessItemAccess.md) |  | [optional] 
**Status** | Pointer to [**CommonAccessItemState**](CommonAccessItemState.md) |  | [optional] 

## Methods

### NewCommonAccessItemRequest

`func NewCommonAccessItemRequest() *CommonAccessItemRequest`

NewCommonAccessItemRequest instantiates a new CommonAccessItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAccessItemRequestWithDefaults

`func NewCommonAccessItemRequestWithDefaults() *CommonAccessItemRequest`

NewCommonAccessItemRequestWithDefaults instantiates a new CommonAccessItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *CommonAccessItemRequest) GetAccess() CommonAccessItemAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CommonAccessItemRequest) GetAccessOk() (*CommonAccessItemAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CommonAccessItemRequest) SetAccess(v CommonAccessItemAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *CommonAccessItemRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetStatus

`func (o *CommonAccessItemRequest) GetStatus() CommonAccessItemState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonAccessItemRequest) GetStatusOk() (*CommonAccessItemState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonAccessItemRequest) SetStatus(v CommonAccessItemState)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonAccessItemRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


