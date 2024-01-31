# FullAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | System-generated unique ID of the Object | [optional] [readonly] 
**Name** | **string** | Name of the Object | 
**Created** | Pointer to **time.Time** | Creation date of the Object | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | Last modification date of the Object | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** | Unique ID from the owning source | [optional] 
**NativeIdentity** | Pointer to **string** | The native identifier of the account | [optional] 
**Description** | Pointer to **NullableString** | The description for the account | [optional] 
**Disabled** | Pointer to **bool** | Whether the account is disabled | [optional] 
**Locked** | Pointer to **bool** | Whether the account is locked | [optional] 
**ManuallyCorrelated** | Pointer to **bool** | Whether the account was manually correlated | [optional] 
**HasEntitlements** | Pointer to **bool** | Whether the account has any entitlements associated with it | [optional] 
**SourceId** | Pointer to **string** | The ID of the source for which this account belongs | [optional] 
**SourceName** | Pointer to **string** | The name of the source | [optional] 
**IdentityId** | Pointer to **string** | The ID of the identity for which this account is correlated to if not uncorrelated | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | A map containing attributes associated with the account | [optional] 
**Authoritative** | Pointer to **bool** | Whether this account belongs to an authoritative source | [optional] 
**SystemAccount** | Pointer to **bool** | Whether this account is for the IdentityNow source | [optional] 
**Uncorrelated** | Pointer to **bool** | True if this account is not correlated to an identity | [optional] 
**Features** | Pointer to **string** | A string list containing the owning source&#39;s features | [optional] 

## Methods

### NewFullAccount

`func NewFullAccount(name string, ) *FullAccount`

NewFullAccount instantiates a new FullAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullAccountWithDefaults

`func NewFullAccountWithDefaults() *FullAccount`

NewFullAccountWithDefaults instantiates a new FullAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FullAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullAccount) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *FullAccount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FullAccount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FullAccount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FullAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *FullAccount) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *FullAccount) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *FullAccount) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *FullAccount) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetUuid

`func (o *FullAccount) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FullAccount) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FullAccount) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FullAccount) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *FullAccount) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *FullAccount) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetNativeIdentity

`func (o *FullAccount) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *FullAccount) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *FullAccount) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.

### HasNativeIdentity

`func (o *FullAccount) HasNativeIdentity() bool`

HasNativeIdentity returns a boolean if a field has been set.

### GetDescription

`func (o *FullAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FullAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FullAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FullAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FullAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FullAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *FullAccount) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FullAccount) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FullAccount) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *FullAccount) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *FullAccount) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *FullAccount) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *FullAccount) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *FullAccount) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *FullAccount) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *FullAccount) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *FullAccount) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *FullAccount) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetHasEntitlements

`func (o *FullAccount) GetHasEntitlements() bool`

GetHasEntitlements returns the HasEntitlements field if non-nil, zero value otherwise.

### GetHasEntitlementsOk

`func (o *FullAccount) GetHasEntitlementsOk() (*bool, bool)`

GetHasEntitlementsOk returns a tuple with the HasEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitlements

`func (o *FullAccount) SetHasEntitlements(v bool)`

SetHasEntitlements sets HasEntitlements field to given value.

### HasHasEntitlements

`func (o *FullAccount) HasHasEntitlements() bool`

HasHasEntitlements returns a boolean if a field has been set.

### GetSourceId

`func (o *FullAccount) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *FullAccount) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *FullAccount) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *FullAccount) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *FullAccount) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *FullAccount) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *FullAccount) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *FullAccount) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetIdentityId

`func (o *FullAccount) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *FullAccount) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *FullAccount) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *FullAccount) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAttributes

`func (o *FullAccount) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FullAccount) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FullAccount) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FullAccount) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAuthoritative

`func (o *FullAccount) GetAuthoritative() bool`

GetAuthoritative returns the Authoritative field if non-nil, zero value otherwise.

### GetAuthoritativeOk

`func (o *FullAccount) GetAuthoritativeOk() (*bool, bool)`

GetAuthoritativeOk returns a tuple with the Authoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritative

`func (o *FullAccount) SetAuthoritative(v bool)`

SetAuthoritative sets Authoritative field to given value.

### HasAuthoritative

`func (o *FullAccount) HasAuthoritative() bool`

HasAuthoritative returns a boolean if a field has been set.

### GetSystemAccount

`func (o *FullAccount) GetSystemAccount() bool`

GetSystemAccount returns the SystemAccount field if non-nil, zero value otherwise.

### GetSystemAccountOk

`func (o *FullAccount) GetSystemAccountOk() (*bool, bool)`

GetSystemAccountOk returns a tuple with the SystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccount

`func (o *FullAccount) SetSystemAccount(v bool)`

SetSystemAccount sets SystemAccount field to given value.

### HasSystemAccount

`func (o *FullAccount) HasSystemAccount() bool`

HasSystemAccount returns a boolean if a field has been set.

### GetUncorrelated

`func (o *FullAccount) GetUncorrelated() bool`

GetUncorrelated returns the Uncorrelated field if non-nil, zero value otherwise.

### GetUncorrelatedOk

`func (o *FullAccount) GetUncorrelatedOk() (*bool, bool)`

GetUncorrelatedOk returns a tuple with the Uncorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncorrelated

`func (o *FullAccount) SetUncorrelated(v bool)`

SetUncorrelated sets Uncorrelated field to given value.

### HasUncorrelated

`func (o *FullAccount) HasUncorrelated() bool`

HasUncorrelated returns a boolean if a field has been set.

### GetFeatures

`func (o *FullAccount) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FullAccount) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FullAccount) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *FullAccount) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


