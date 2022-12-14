# MultiStatusObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **[]string** |  | [optional] 
**Fail** | Pointer to [**[]MultiStatusObjectFailInner**](MultiStatusObjectFailInner.md) |  | [optional] 

## Methods

### NewMultiStatusObject

`func NewMultiStatusObject() *MultiStatusObject`

NewMultiStatusObject instantiates a new MultiStatusObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiStatusObjectWithDefaults

`func NewMultiStatusObjectWithDefaults() *MultiStatusObject`

NewMultiStatusObjectWithDefaults instantiates a new MultiStatusObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MultiStatusObject) GetSuccess() []string`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MultiStatusObject) GetSuccessOk() (*[]string, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MultiStatusObject) SetSuccess(v []string)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *MultiStatusObject) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFail

`func (o *MultiStatusObject) GetFail() []MultiStatusObjectFailInner`

GetFail returns the Fail field if non-nil, zero value otherwise.

### GetFailOk

`func (o *MultiStatusObject) GetFailOk() (*[]MultiStatusObjectFailInner, bool)`

GetFailOk returns a tuple with the Fail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFail

`func (o *MultiStatusObject) SetFail(v []MultiStatusObjectFailInner)`

SetFail sets Fail field to given value.

### HasFail

`func (o *MultiStatusObject) HasFail() bool`

HasFail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


