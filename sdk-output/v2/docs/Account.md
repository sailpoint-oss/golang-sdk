# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Application** | Pointer to [**AccountApplication**](AccountApplication.md) |  | [optional] 
**Attributes** | Pointer to [**AccountAttributes**](AccountAttributes.md) |  | [optional] 
**Authoritative** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Identity** | Pointer to [**AccountIdentity**](AccountIdentity.md) |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Meta** | Pointer to [**AccountMeta**](AccountMeta.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NativeIdentity** | Pointer to **string** |  | [optional] 
**PendingAccessRequestIds** | Pointer to **[]string** |  | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**SupportsPasswordChange** | Pointer to **bool** |  | [optional] 
**SystemAccount** | Pointer to **bool** |  | [optional] 
**Uncorrelated** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Account) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Account) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Account) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Account) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetApplication

`func (o *Account) GetApplication() AccountApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Account) GetApplicationOk() (*AccountApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Account) SetApplication(v AccountApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Account) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetAttributes

`func (o *Account) GetAttributes() AccountAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Account) GetAttributesOk() (*AccountAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Account) SetAttributes(v AccountAttributes)`

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

### GetDisplayName

`func (o *Account) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Account) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Account) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Account) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFeatures

`func (o *Account) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Account) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Account) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Account) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

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

### GetIdentity

`func (o *Account) GetIdentity() AccountIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *Account) GetIdentityOk() (*AccountIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *Account) SetIdentity(v AccountIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *Account) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

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

### GetMeta

`func (o *Account) GetMeta() AccountMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Account) GetMetaOk() (*AccountMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Account) SetMeta(v AccountMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Account) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

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

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

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

### GetPendingAccessRequestIds

`func (o *Account) GetPendingAccessRequestIds() []string`

GetPendingAccessRequestIds returns the PendingAccessRequestIds field if non-nil, zero value otherwise.

### GetPendingAccessRequestIdsOk

`func (o *Account) GetPendingAccessRequestIdsOk() (*[]string, bool)`

GetPendingAccessRequestIdsOk returns a tuple with the PendingAccessRequestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAccessRequestIds

`func (o *Account) SetPendingAccessRequestIds(v []string)`

SetPendingAccessRequestIds sets PendingAccessRequestIds field to given value.

### HasPendingAccessRequestIds

`func (o *Account) HasPendingAccessRequestIds() bool`

HasPendingAccessRequestIds returns a boolean if a field has been set.

### GetSchema

`func (o *Account) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Account) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Account) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *Account) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSupportsPasswordChange

`func (o *Account) GetSupportsPasswordChange() bool`

GetSupportsPasswordChange returns the SupportsPasswordChange field if non-nil, zero value otherwise.

### GetSupportsPasswordChangeOk

`func (o *Account) GetSupportsPasswordChangeOk() (*bool, bool)`

GetSupportsPasswordChangeOk returns a tuple with the SupportsPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPasswordChange

`func (o *Account) SetSupportsPasswordChange(v bool)`

SetSupportsPasswordChange sets SupportsPasswordChange field to given value.

### HasSupportsPasswordChange

`func (o *Account) HasSupportsPasswordChange() bool`

HasSupportsPasswordChange returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


