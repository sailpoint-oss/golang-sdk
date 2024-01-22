# ListAccounts200ResponseInner

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

### NewListAccounts200ResponseInner

`func NewListAccounts200ResponseInner(name string, ) *ListAccounts200ResponseInner`

NewListAccounts200ResponseInner instantiates a new ListAccounts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccounts200ResponseInnerWithDefaults

`func NewListAccounts200ResponseInnerWithDefaults() *ListAccounts200ResponseInner`

NewListAccounts200ResponseInnerWithDefaults instantiates a new ListAccounts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAccounts200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAccounts200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAccounts200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListAccounts200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListAccounts200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAccounts200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAccounts200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *ListAccounts200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListAccounts200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListAccounts200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListAccounts200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ListAccounts200ResponseInner) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ListAccounts200ResponseInner) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ListAccounts200ResponseInner) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ListAccounts200ResponseInner) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetUuid

`func (o *ListAccounts200ResponseInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListAccounts200ResponseInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListAccounts200ResponseInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListAccounts200ResponseInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ListAccounts200ResponseInner) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ListAccounts200ResponseInner) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetNativeIdentity

`func (o *ListAccounts200ResponseInner) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *ListAccounts200ResponseInner) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *ListAccounts200ResponseInner) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.

### HasNativeIdentity

`func (o *ListAccounts200ResponseInner) HasNativeIdentity() bool`

HasNativeIdentity returns a boolean if a field has been set.

### GetDescription

`func (o *ListAccounts200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListAccounts200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListAccounts200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListAccounts200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListAccounts200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListAccounts200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *ListAccounts200ResponseInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ListAccounts200ResponseInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ListAccounts200ResponseInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ListAccounts200ResponseInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *ListAccounts200ResponseInner) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ListAccounts200ResponseInner) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ListAccounts200ResponseInner) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ListAccounts200ResponseInner) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *ListAccounts200ResponseInner) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *ListAccounts200ResponseInner) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *ListAccounts200ResponseInner) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *ListAccounts200ResponseInner) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetHasEntitlements

`func (o *ListAccounts200ResponseInner) GetHasEntitlements() bool`

GetHasEntitlements returns the HasEntitlements field if non-nil, zero value otherwise.

### GetHasEntitlementsOk

`func (o *ListAccounts200ResponseInner) GetHasEntitlementsOk() (*bool, bool)`

GetHasEntitlementsOk returns a tuple with the HasEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitlements

`func (o *ListAccounts200ResponseInner) SetHasEntitlements(v bool)`

SetHasEntitlements sets HasEntitlements field to given value.

### HasHasEntitlements

`func (o *ListAccounts200ResponseInner) HasHasEntitlements() bool`

HasHasEntitlements returns a boolean if a field has been set.

### GetSourceId

`func (o *ListAccounts200ResponseInner) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ListAccounts200ResponseInner) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ListAccounts200ResponseInner) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ListAccounts200ResponseInner) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *ListAccounts200ResponseInner) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ListAccounts200ResponseInner) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ListAccounts200ResponseInner) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *ListAccounts200ResponseInner) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetIdentityId

`func (o *ListAccounts200ResponseInner) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *ListAccounts200ResponseInner) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *ListAccounts200ResponseInner) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *ListAccounts200ResponseInner) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAttributes

`func (o *ListAccounts200ResponseInner) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ListAccounts200ResponseInner) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ListAccounts200ResponseInner) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ListAccounts200ResponseInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAuthoritative

`func (o *ListAccounts200ResponseInner) GetAuthoritative() bool`

GetAuthoritative returns the Authoritative field if non-nil, zero value otherwise.

### GetAuthoritativeOk

`func (o *ListAccounts200ResponseInner) GetAuthoritativeOk() (*bool, bool)`

GetAuthoritativeOk returns a tuple with the Authoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritative

`func (o *ListAccounts200ResponseInner) SetAuthoritative(v bool)`

SetAuthoritative sets Authoritative field to given value.

### HasAuthoritative

`func (o *ListAccounts200ResponseInner) HasAuthoritative() bool`

HasAuthoritative returns a boolean if a field has been set.

### GetSystemAccount

`func (o *ListAccounts200ResponseInner) GetSystemAccount() bool`

GetSystemAccount returns the SystemAccount field if non-nil, zero value otherwise.

### GetSystemAccountOk

`func (o *ListAccounts200ResponseInner) GetSystemAccountOk() (*bool, bool)`

GetSystemAccountOk returns a tuple with the SystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccount

`func (o *ListAccounts200ResponseInner) SetSystemAccount(v bool)`

SetSystemAccount sets SystemAccount field to given value.

### HasSystemAccount

`func (o *ListAccounts200ResponseInner) HasSystemAccount() bool`

HasSystemAccount returns a boolean if a field has been set.

### GetUncorrelated

`func (o *ListAccounts200ResponseInner) GetUncorrelated() bool`

GetUncorrelated returns the Uncorrelated field if non-nil, zero value otherwise.

### GetUncorrelatedOk

`func (o *ListAccounts200ResponseInner) GetUncorrelatedOk() (*bool, bool)`

GetUncorrelatedOk returns a tuple with the Uncorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncorrelated

`func (o *ListAccounts200ResponseInner) SetUncorrelated(v bool)`

SetUncorrelated sets Uncorrelated field to given value.

### HasUncorrelated

`func (o *ListAccounts200ResponseInner) HasUncorrelated() bool`

HasUncorrelated returns a boolean if a field has been set.

### GetFeatures

`func (o *ListAccounts200ResponseInner) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ListAccounts200ResponseInner) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ListAccounts200ResponseInner) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ListAccounts200ResponseInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


