# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | System-generated unique ID of the Object | [optional] [readonly] 
**Name** | **string** | Name of the Object | 
**Created** | Pointer to **time.Time** | Creation date of the Object | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | Last modification date of the Object | [optional] [readonly] 
**SourceId** | Pointer to **string** |  | [optional] 
**IdentityId** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Authoritative** | Pointer to **bool** | Indicates if this account is from an authoritative source | [optional] 
**Description** | Pointer to **NullableString** | A description of the account | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the account is currently disabled | [optional] 
**Locked** | Pointer to **bool** | Indicates if the account is currently locked | [optional] 
**NativeIdentity** | Pointer to **string** |  | [optional] 
**SystemAccount** | Pointer to **bool** |  | [optional] 
**Uncorrelated** | Pointer to **bool** | Indicates if this account is not correlated to an identity | [optional] 
**Uuid** | Pointer to **string** | The unique ID of the account as determined by the account schema | [optional] 
**ManuallyCorrelated** | Pointer to **bool** | Indicates if the account has been manually correlated to an identity | [optional] 
**HasEntitlements** | Pointer to **bool** | Indicates if the account has entitlements | [optional] 

## Methods

### NewAccount

`func NewAccount(name string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *Account) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Account) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Account) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Account) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Account) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Account) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Account) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Account) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetSourceId

`func (o *Account) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Account) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Account) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Account) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetIdentityId

`func (o *Account) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Account) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Account) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *Account) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAttributes

`func (o *Account) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Account) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Account) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Account) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAuthoritative

`func (o *Account) GetAuthoritative() bool`

GetAuthoritative returns the Authoritative field if non-nil, zero value otherwise.

### GetAuthoritativeOk

`func (o *Account) GetAuthoritativeOk() (*bool, bool)`

GetAuthoritativeOk returns a tuple with the Authoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritative

`func (o *Account) SetAuthoritative(v bool)`

SetAuthoritative sets Authoritative field to given value.

### HasAuthoritative

`func (o *Account) HasAuthoritative() bool`

HasAuthoritative returns a boolean if a field has been set.

### GetDescription

`func (o *Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Account) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Account) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Account) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *Account) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Account) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Account) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Account) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *Account) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Account) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Account) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Account) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNativeIdentity

`func (o *Account) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *Account) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *Account) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.

### HasNativeIdentity

`func (o *Account) HasNativeIdentity() bool`

HasNativeIdentity returns a boolean if a field has been set.

### GetSystemAccount

`func (o *Account) GetSystemAccount() bool`

GetSystemAccount returns the SystemAccount field if non-nil, zero value otherwise.

### GetSystemAccountOk

`func (o *Account) GetSystemAccountOk() (*bool, bool)`

GetSystemAccountOk returns a tuple with the SystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccount

`func (o *Account) SetSystemAccount(v bool)`

SetSystemAccount sets SystemAccount field to given value.

### HasSystemAccount

`func (o *Account) HasSystemAccount() bool`

HasSystemAccount returns a boolean if a field has been set.

### GetUncorrelated

`func (o *Account) GetUncorrelated() bool`

GetUncorrelated returns the Uncorrelated field if non-nil, zero value otherwise.

### GetUncorrelatedOk

`func (o *Account) GetUncorrelatedOk() (*bool, bool)`

GetUncorrelatedOk returns a tuple with the Uncorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncorrelated

`func (o *Account) SetUncorrelated(v bool)`

SetUncorrelated sets Uncorrelated field to given value.

### HasUncorrelated

`func (o *Account) HasUncorrelated() bool`

HasUncorrelated returns a boolean if a field has been set.

### GetUuid

`func (o *Account) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Account) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Account) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Account) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *Account) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *Account) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *Account) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *Account) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetHasEntitlements

`func (o *Account) GetHasEntitlements() bool`

GetHasEntitlements returns the HasEntitlements field if non-nil, zero value otherwise.

### GetHasEntitlementsOk

`func (o *Account) GetHasEntitlementsOk() (*bool, bool)`

GetHasEntitlementsOk returns a tuple with the HasEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitlements

`func (o *Account) SetHasEntitlements(v bool)`

SetHasEntitlements sets HasEntitlements field to given value.

### HasHasEntitlements

`func (o *Account) HasHasEntitlements() bool`

HasHasEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


