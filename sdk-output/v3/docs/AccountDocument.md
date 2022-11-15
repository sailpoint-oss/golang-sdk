# AccountDocument

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

### NewAccountDocument

`func NewAccountDocument(id string, name string, type_ DocumentType, ) *AccountDocument`

NewAccountDocument instantiates a new AccountDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDocumentWithDefaults

`func NewAccountDocumentWithDefaults() *AccountDocument`

NewAccountDocumentWithDefaults instantiates a new AccountDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccountDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountDocument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AccountDocument) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountDocument) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountDocument) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetAccountId

`func (o *AccountDocument) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountDocument) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountDocument) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountDocument) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSource

`func (o *AccountDocument) GetSource() Source1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccountDocument) GetSourceOk() (*Source1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccountDocument) SetSource(v Source1)`

SetSource sets Source field to given value.

### HasSource

`func (o *AccountDocument) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDisabled

`func (o *AccountDocument) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AccountDocument) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AccountDocument) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AccountDocument) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *AccountDocument) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AccountDocument) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AccountDocument) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *AccountDocument) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPrivileged

`func (o *AccountDocument) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *AccountDocument) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *AccountDocument) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *AccountDocument) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *AccountDocument) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *AccountDocument) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *AccountDocument) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *AccountDocument) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetPasswordLastSet

`func (o *AccountDocument) GetPasswordLastSet() time.Time`

GetPasswordLastSet returns the PasswordLastSet field if non-nil, zero value otherwise.

### GetPasswordLastSetOk

`func (o *AccountDocument) GetPasswordLastSetOk() (*time.Time, bool)`

GetPasswordLastSetOk returns a tuple with the PasswordLastSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLastSet

`func (o *AccountDocument) SetPasswordLastSet(v time.Time)`

SetPasswordLastSet sets PasswordLastSet field to given value.

### HasPasswordLastSet

`func (o *AccountDocument) HasPasswordLastSet() bool`

HasPasswordLastSet returns a boolean if a field has been set.

### SetPasswordLastSetNil

`func (o *AccountDocument) SetPasswordLastSetNil(b bool)`

 SetPasswordLastSetNil sets the value for PasswordLastSet to be an explicit nil

### UnsetPasswordLastSet
`func (o *AccountDocument) UnsetPasswordLastSet()`

UnsetPasswordLastSet ensures that no value is present for PasswordLastSet, not even an explicit nil
### GetEntitlementAttributes

`func (o *AccountDocument) GetEntitlementAttributes() map[string]interface{}`

GetEntitlementAttributes returns the EntitlementAttributes field if non-nil, zero value otherwise.

### GetEntitlementAttributesOk

`func (o *AccountDocument) GetEntitlementAttributesOk() (*map[string]interface{}, bool)`

GetEntitlementAttributesOk returns a tuple with the EntitlementAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementAttributes

`func (o *AccountDocument) SetEntitlementAttributes(v map[string]interface{})`

SetEntitlementAttributes sets EntitlementAttributes field to given value.

### HasEntitlementAttributes

`func (o *AccountDocument) HasEntitlementAttributes() bool`

HasEntitlementAttributes returns a boolean if a field has been set.

### SetEntitlementAttributesNil

`func (o *AccountDocument) SetEntitlementAttributesNil(b bool)`

 SetEntitlementAttributesNil sets the value for EntitlementAttributes to be an explicit nil

### UnsetEntitlementAttributes
`func (o *AccountDocument) UnsetEntitlementAttributes()`

UnsetEntitlementAttributes ensures that no value is present for EntitlementAttributes, not even an explicit nil
### GetCreated

`func (o *AccountDocument) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountDocument) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountDocument) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccountDocument) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *AccountDocument) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *AccountDocument) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *AccountDocument) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AccountDocument) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AccountDocument) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AccountDocument) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *AccountDocument) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *AccountDocument) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetAttributes

`func (o *AccountDocument) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AccountDocument) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AccountDocument) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AccountDocument) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIdentity

`func (o *AccountDocument) GetIdentity() DisplayReference`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AccountDocument) GetIdentityOk() (*DisplayReference, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AccountDocument) SetIdentity(v DisplayReference)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AccountDocument) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetAccess

`func (o *AccountDocument) GetAccess() []Entitlement1`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AccountDocument) GetAccessOk() (*[]Entitlement1, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AccountDocument) SetAccess(v []Entitlement1)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AccountDocument) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *AccountDocument) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *AccountDocument) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *AccountDocument) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *AccountDocument) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetUncorrelated

`func (o *AccountDocument) GetUncorrelated() bool`

GetUncorrelated returns the Uncorrelated field if non-nil, zero value otherwise.

### GetUncorrelatedOk

`func (o *AccountDocument) GetUncorrelatedOk() (*bool, bool)`

GetUncorrelatedOk returns a tuple with the Uncorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncorrelated

`func (o *AccountDocument) SetUncorrelated(v bool)`

SetUncorrelated sets Uncorrelated field to given value.

### HasUncorrelated

`func (o *AccountDocument) HasUncorrelated() bool`

HasUncorrelated returns a boolean if a field has been set.

### GetTags

`func (o *AccountDocument) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccountDocument) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccountDocument) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccountDocument) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


