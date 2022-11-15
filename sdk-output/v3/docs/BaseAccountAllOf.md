# BaseAccountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The ID of the account | [optional] 
**Source** | Pointer to [**Source1**](Source1.md) |  | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the account is disabled | [optional] 
**Locked** | Pointer to **bool** | Indicates if the account is locked | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**ManuallyCorrelated** | Pointer to **bool** | Indicates if the account has been manually correlated to an identity | [optional] 
**PasswordLastSet** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**EntitlementAttributes** | Pointer to **map[string]interface{}** | a map or dictionary of key/value pairs | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 

## Methods

### NewBaseAccountAllOf

`func NewBaseAccountAllOf() *BaseAccountAllOf`

NewBaseAccountAllOf instantiates a new BaseAccountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAccountAllOfWithDefaults

`func NewBaseAccountAllOfWithDefaults() *BaseAccountAllOf`

NewBaseAccountAllOfWithDefaults instantiates a new BaseAccountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BaseAccountAllOf) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BaseAccountAllOf) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BaseAccountAllOf) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BaseAccountAllOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSource

`func (o *BaseAccountAllOf) GetSource() Source1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BaseAccountAllOf) GetSourceOk() (*Source1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BaseAccountAllOf) SetSource(v Source1)`

SetSource sets Source field to given value.

### HasSource

`func (o *BaseAccountAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDisabled

`func (o *BaseAccountAllOf) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *BaseAccountAllOf) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *BaseAccountAllOf) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *BaseAccountAllOf) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *BaseAccountAllOf) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *BaseAccountAllOf) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *BaseAccountAllOf) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *BaseAccountAllOf) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPrivileged

`func (o *BaseAccountAllOf) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *BaseAccountAllOf) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *BaseAccountAllOf) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *BaseAccountAllOf) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *BaseAccountAllOf) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *BaseAccountAllOf) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *BaseAccountAllOf) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *BaseAccountAllOf) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetPasswordLastSet

`func (o *BaseAccountAllOf) GetPasswordLastSet() time.Time`

GetPasswordLastSet returns the PasswordLastSet field if non-nil, zero value otherwise.

### GetPasswordLastSetOk

`func (o *BaseAccountAllOf) GetPasswordLastSetOk() (*time.Time, bool)`

GetPasswordLastSetOk returns a tuple with the PasswordLastSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastSet

`func (o *BaseAccountAllOf) SetPasswordLastSet(v time.Time)`

SetPasswordLastSet sets PasswordLastSet field to given value.

### HasPasswordLastSet

`func (o *BaseAccountAllOf) HasPasswordLastSet() bool`

HasPasswordLastSet returns a boolean if a field has been set.

### SetPasswordLastSetNil

`func (o *BaseAccountAllOf) SetPasswordLastSetNil(b bool)`

 SetPasswordLastSetNil sets the value for PasswordLastSet to be an explicit nil

### UnsetPasswordLastSet
`func (o *BaseAccountAllOf) UnsetPasswordLastSet()`

UnsetPasswordLastSet ensures that no value is present for PasswordLastSet, not even an explicit nil
### GetEntitlementAttributes

`func (o *BaseAccountAllOf) GetEntitlementAttributes() map[string]interface{}`

GetEntitlementAttributes returns the EntitlementAttributes field if non-nil, zero value otherwise.

### GetEntitlementAttributesOk

`func (o *BaseAccountAllOf) GetEntitlementAttributesOk() (*map[string]interface{}, bool)`

GetEntitlementAttributesOk returns a tuple with the EntitlementAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementAttributes

`func (o *BaseAccountAllOf) SetEntitlementAttributes(v map[string]interface{})`

SetEntitlementAttributes sets EntitlementAttributes field to given value.

### HasEntitlementAttributes

`func (o *BaseAccountAllOf) HasEntitlementAttributes() bool`

HasEntitlementAttributes returns a boolean if a field has been set.

### SetEntitlementAttributesNil

`func (o *BaseAccountAllOf) SetEntitlementAttributesNil(b bool)`

 SetEntitlementAttributesNil sets the value for EntitlementAttributes to be an explicit nil

### UnsetEntitlementAttributes
`func (o *BaseAccountAllOf) UnsetEntitlementAttributes()`

UnsetEntitlementAttributes ensures that no value is present for EntitlementAttributes, not even an explicit nil
### GetCreated

`func (o *BaseAccountAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BaseAccountAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BaseAccountAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BaseAccountAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *BaseAccountAllOf) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *BaseAccountAllOf) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


