# SlimAccount

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

## Methods

### NewSlimAccount

`func NewSlimAccount(name string, ) *SlimAccount`

NewSlimAccount instantiates a new SlimAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlimAccountWithDefaults

`func NewSlimAccountWithDefaults() *SlimAccount`

NewSlimAccountWithDefaults instantiates a new SlimAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlimAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlimAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlimAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SlimAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SlimAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlimAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlimAccount) SetName(v string)`

SetName sets Name field to given value.


### GetCreated

`func (o *SlimAccount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SlimAccount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SlimAccount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SlimAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *SlimAccount) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SlimAccount) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SlimAccount) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SlimAccount) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetUuid

`func (o *SlimAccount) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SlimAccount) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SlimAccount) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SlimAccount) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *SlimAccount) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *SlimAccount) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetNativeIdentity

`func (o *SlimAccount) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *SlimAccount) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *SlimAccount) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.

### HasNativeIdentity

`func (o *SlimAccount) HasNativeIdentity() bool`

HasNativeIdentity returns a boolean if a field has been set.

### GetDescription

`func (o *SlimAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlimAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlimAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlimAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SlimAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SlimAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *SlimAccount) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SlimAccount) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SlimAccount) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SlimAccount) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *SlimAccount) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SlimAccount) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SlimAccount) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SlimAccount) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *SlimAccount) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *SlimAccount) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *SlimAccount) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *SlimAccount) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetHasEntitlements

`func (o *SlimAccount) GetHasEntitlements() bool`

GetHasEntitlements returns the HasEntitlements field if non-nil, zero value otherwise.

### GetHasEntitlementsOk

`func (o *SlimAccount) GetHasEntitlementsOk() (*bool, bool)`

GetHasEntitlementsOk returns a tuple with the HasEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitlements

`func (o *SlimAccount) SetHasEntitlements(v bool)`

SetHasEntitlements sets HasEntitlements field to given value.

### HasHasEntitlements

`func (o *SlimAccount) HasHasEntitlements() bool`

HasHasEntitlements returns a boolean if a field has been set.

### GetSourceId

`func (o *SlimAccount) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SlimAccount) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SlimAccount) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SlimAccount) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *SlimAccount) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *SlimAccount) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *SlimAccount) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *SlimAccount) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetIdentityId

`func (o *SlimAccount) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *SlimAccount) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *SlimAccount) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *SlimAccount) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAttributes

`func (o *SlimAccount) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SlimAccount) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SlimAccount) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SlimAccount) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


