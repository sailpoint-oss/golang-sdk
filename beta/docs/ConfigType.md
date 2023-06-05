# ConfigType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalName** | Pointer to [**ConfigTypeEnum**](ConfigTypeEnum.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Human readable display name of the type to be shown on UI | [optional] 
**Description** | Pointer to **string** | Description of the type of work to be reassigned, displayed by the UI. | [optional] 

## Methods

### NewConfigType

`func NewConfigType() *ConfigType`

NewConfigType instantiates a new ConfigType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigTypeWithDefaults

`func NewConfigTypeWithDefaults() *ConfigType`

NewConfigTypeWithDefaults instantiates a new ConfigType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalName

`func (o *ConfigType) GetInternalName() ConfigTypeEnum`

GetInternalName returns the InternalName field if non-nil, zero value otherwise.

### GetInternalNameOk

`func (o *ConfigType) GetInternalNameOk() (*ConfigTypeEnum, bool)`

GetInternalNameOk returns a tuple with the InternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalName

`func (o *ConfigType) SetInternalName(v ConfigTypeEnum)`

SetInternalName sets InternalName field to given value.

### HasInternalName

`func (o *ConfigType) HasInternalName() bool`

HasInternalName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConfigType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConfigType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConfigType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConfigType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


