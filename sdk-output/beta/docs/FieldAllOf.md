# FieldAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the field | [optional] 
**DisplayType** | Pointer to **string** | Type of the field to display | [optional] 
**Required** | Pointer to **bool** | True if the field is required | [optional] 
**AllowedValuesList** | Pointer to **[]map[string]interface{}** | List of allowed values for the field | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value of the field | [optional] 

## Methods

### NewFieldAllOf

`func NewFieldAllOf() *FieldAllOf`

NewFieldAllOf instantiates a new FieldAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldAllOfWithDefaults

`func NewFieldAllOfWithDefaults() *FieldAllOf`

NewFieldAllOfWithDefaults instantiates a new FieldAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FieldAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FieldAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FieldAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FieldAllOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayType

`func (o *FieldAllOf) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *FieldAllOf) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *FieldAllOf) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *FieldAllOf) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetRequired

`func (o *FieldAllOf) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldAllOf) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldAllOf) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldAllOf) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetAllowedValuesList

`func (o *FieldAllOf) GetAllowedValuesList() []map[string]interface{}`

GetAllowedValuesList returns the AllowedValuesList field if non-nil, zero value otherwise.

### GetAllowedValuesListOk

`func (o *FieldAllOf) GetAllowedValuesListOk() (*[]map[string]interface{}, bool)`

GetAllowedValuesListOk returns a tuple with the AllowedValuesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValuesList

`func (o *FieldAllOf) SetAllowedValuesList(v []map[string]interface{})`

SetAllowedValuesList sets AllowedValuesList field to given value.

### HasAllowedValuesList

`func (o *FieldAllOf) HasAllowedValuesList() bool`

HasAllowedValuesList returns a boolean if a field has been set.

### GetValue

`func (o *FieldAllOf) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldAllOf) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldAllOf) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *FieldAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


