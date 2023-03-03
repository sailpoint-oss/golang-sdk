# MfaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If MFA method is enabled. | [optional] 
**Host** | Pointer to **string** | The server host name or IP address of the MFA provider. | [optional] 
**AccessKey** | Pointer to **string** | The secret key for authenticating requests to the MFA provider. | [optional] 
**IdentityAttribute** | Pointer to **string** | Optional. The name of the attribute for mapping IdentityNow identity to the MFA provider. | [optional] 

## Methods

### NewMfaConfig

`func NewMfaConfig() *MfaConfig`

NewMfaConfig instantiates a new MfaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMfaConfigWithDefaults

`func NewMfaConfigWithDefaults() *MfaConfig`

NewMfaConfigWithDefaults instantiates a new MfaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MfaConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MfaConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MfaConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MfaConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *MfaConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MfaConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MfaConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MfaConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAccessKey

`func (o *MfaConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *MfaConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *MfaConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *MfaConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetIdentityAttribute

`func (o *MfaConfig) GetIdentityAttribute() string`

GetIdentityAttribute returns the IdentityAttribute field if non-nil, zero value otherwise.

### GetIdentityAttributeOk

`func (o *MfaConfig) GetIdentityAttributeOk() (*string, bool)`

GetIdentityAttributeOk returns a tuple with the IdentityAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAttribute

`func (o *MfaConfig) SetIdentityAttribute(v string)`

SetIdentityAttribute sets IdentityAttribute field to given value.

### HasIdentityAttribute

`func (o *MfaConfig) HasIdentityAttribute() bool`

HasIdentityAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


