# AttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The attribute name | [optional] 
**Op** | Pointer to **string** | The operation to perform | [optional] 
**Value** | Pointer to **string** | The value of the attribute | [optional] 

## Methods

### NewAttributeRequest

`func NewAttributeRequest() *AttributeRequest`

NewAttributeRequest instantiates a new AttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeRequestWithDefaults

`func NewAttributeRequestWithDefaults() *AttributeRequest`

NewAttributeRequestWithDefaults instantiates a new AttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttributeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOp

`func (o *AttributeRequest) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AttributeRequest) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AttributeRequest) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *AttributeRequest) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetValue

`func (o *AttributeRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AttributeRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AttributeRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AttributeRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


