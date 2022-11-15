# Account1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
**AccountId** | Pointer to **string** | The ID of the account | [optional] 
**Source** | Pointer to [**Source1**](Source1.md) |  | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the account is disabled | [optional] 
**Locked** | Pointer to **bool** | Indicates if the account is locked | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**ManuallyCorrelated** | Pointer to **bool** | Indicates if the account has been manually correlated to an identity | [optional] 
**PasswordLastSet** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**EntitlementAttributes** | Pointer to **map[string]interface{}** | a map or dictionary of key/value pairs | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | a map or dictionary of key/value pairs | [optional] 
**Identity** | Pointer to [**DisplayReference**](DisplayReference.md) |  | [optional] 
**Access** | Pointer to [**[]Entitlement1**](Entitlement1.md) |  | [optional] 
**EntitlementCount** | Pointer to **int32** | The number of entitlements assigned to the account | [optional] 
**Uncorrelated** | Pointer to **bool** | Indicates if the account is not correlated to an identity | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccount1

`func NewAccount1(id string, name string, type_ DocumentType, ) *Account1`

NewAccount1 instantiates a new Account1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccount1WithDefaults

`func NewAccount1WithDefaults() *Account1`

NewAccount1WithDefaults instantiates a new Account1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Account1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account1) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Account1) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account1) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account1) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetAccountId

`func (o *Account1) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Account1) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Account1) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Account1) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSource

`func (o *Account1) GetSource() Source1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Account1) GetSourceOk() (*Source1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Account1) SetSource(v Source1)`

SetSource sets Source field to given value.

### HasSource

`func (o *Account1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDisabled

`func (o *Account1) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Account1) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Account1) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Account1) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *Account1) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Account1) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Account1) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Account1) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPrivileged

`func (o *Account1) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *Account1) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *Account1) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *Account1) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *Account1) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *Account1) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *Account1) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *Account1) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetPasswordLastSet

`func (o *Account1) GetPasswordLastSet() time.Time`

GetPasswordLastSet returns the PasswordLastSet field if non-nil, zero value otherwise.

### GetPasswordLastSetOk

`func (o *Account1) GetPasswordLastSetOk() (*time.Time, bool)`

GetPasswordLastSetOk returns a tuple with the PasswordLastSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastSet

`func (o *Account1) SetPasswordLastSet(v time.Time)`

SetPasswordLastSet sets PasswordLastSet field to given value.

### HasPasswordLastSet

`func (o *Account1) HasPasswordLastSet() bool`

HasPasswordLastSet returns a boolean if a field has been set.

### SetPasswordLastSetNil

`func (o *Account1) SetPasswordLastSetNil(b bool)`

 SetPasswordLastSetNil sets the value for PasswordLastSet to be an explicit nil

### UnsetPasswordLastSet
`func (o *Account1) UnsetPasswordLastSet()`

UnsetPasswordLastSet ensures that no value is present for PasswordLastSet, not even an explicit nil
### GetEntitlementAttributes

`func (o *Account1) GetEntitlementAttributes() map[string]interface{}`

GetEntitlementAttributes returns the EntitlementAttributes field if non-nil, zero value otherwise.

### GetEntitlementAttributesOk

`func (o *Account1) GetEntitlementAttributesOk() (*map[string]interface{}, bool)`

GetEntitlementAttributesOk returns a tuple with the EntitlementAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementAttributes

`func (o *Account1) SetEntitlementAttributes(v map[string]interface{})`

SetEntitlementAttributes sets EntitlementAttributes field to given value.

### HasEntitlementAttributes

`func (o *Account1) HasEntitlementAttributes() bool`

HasEntitlementAttributes returns a boolean if a field has been set.

### SetEntitlementAttributesNil

`func (o *Account1) SetEntitlementAttributesNil(b bool)`

 SetEntitlementAttributesNil sets the value for EntitlementAttributes to be an explicit nil

### UnsetEntitlementAttributes
`func (o *Account1) UnsetEntitlementAttributes()`

UnsetEntitlementAttributes ensures that no value is present for EntitlementAttributes, not even an explicit nil
### GetCreated

`func (o *Account1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Account1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Account1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Account1) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Account1) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Account1) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *Account1) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Account1) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Account1) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Account1) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *Account1) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *Account1) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetAttributes

`func (o *Account1) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Account1) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Account1) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Account1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIdentity

`func (o *Account1) GetIdentity() DisplayReference`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *Account1) GetIdentityOk() (*DisplayReference, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *Account1) SetIdentity(v DisplayReference)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *Account1) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetAccess

`func (o *Account1) GetAccess() []Entitlement1`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Account1) GetAccessOk() (*[]Entitlement1, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Account1) SetAccess(v []Entitlement1)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *Account1) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *Account1) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *Account1) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *Account1) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *Account1) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetUncorrelated

`func (o *Account1) GetUncorrelated() bool`

GetUncorrelated returns the Uncorrelated field if non-nil, zero value otherwise.

### GetUncorrelatedOk

`func (o *Account1) GetUncorrelatedOk() (*bool, bool)`

GetUncorrelatedOk returns a tuple with the Uncorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncorrelated

`func (o *Account1) SetUncorrelated(v bool)`

SetUncorrelated sets Uncorrelated field to given value.

### HasUncorrelated

`func (o *Account1) HasUncorrelated() bool`

HasUncorrelated returns a boolean if a field has been set.

### GetTags

`func (o *Account1) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Account1) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Account1) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Account1) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


