# ContextAttributeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** | The name of the attribute | [optional] 
**Value** | Pointer to [**ContextAttributeDtoValue**](ContextAttributeDtoValue.md) |  | [optional] 

## Methods

### NewContextAttributeDto

`func NewContextAttributeDto() *ContextAttributeDto`

NewContextAttributeDto instantiates a new ContextAttributeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextAttributeDtoWithDefaults

`func NewContextAttributeDtoWithDefaults() *ContextAttributeDto`

NewContextAttributeDtoWithDefaults instantiates a new ContextAttributeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *ContextAttributeDto) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *ContextAttributeDto) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *ContextAttributeDto) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *ContextAttributeDto) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *ContextAttributeDto) GetValue() ContextAttributeDtoValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContextAttributeDto) GetValueOk() (*ContextAttributeDtoValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContextAttributeDto) SetValue(v ContextAttributeDtoValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContextAttributeDto) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


