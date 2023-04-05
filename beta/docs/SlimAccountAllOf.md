# SlimAccountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSlimAccountAllOf

`func NewSlimAccountAllOf() *SlimAccountAllOf`

NewSlimAccountAllOf instantiates a new SlimAccountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlimAccountAllOfWithDefaults

`func NewSlimAccountAllOfWithDefaults() *SlimAccountAllOf`

NewSlimAccountAllOfWithDefaults instantiates a new SlimAccountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SlimAccountAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SlimAccountAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SlimAccountAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SlimAccountAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *SlimAccountAllOf) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *SlimAccountAllOf) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetNativeIdentity

`func (o *SlimAccountAllOf) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *SlimAccountAllOf) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *SlimAccountAllOf) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.

### HasNativeIdentity

`func (o *SlimAccountAllOf) HasNativeIdentity() bool`

HasNativeIdentity returns a boolean if a field has been set.

### GetDescription

`func (o *SlimAccountAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SlimAccountAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SlimAccountAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SlimAccountAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SlimAccountAllOf) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SlimAccountAllOf) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *SlimAccountAllOf) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SlimAccountAllOf) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SlimAccountAllOf) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SlimAccountAllOf) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *SlimAccountAllOf) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SlimAccountAllOf) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SlimAccountAllOf) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SlimAccountAllOf) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetManuallyCorrelated

`func (o *SlimAccountAllOf) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *SlimAccountAllOf) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *SlimAccountAllOf) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *SlimAccountAllOf) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetHasEntitlements

`func (o *SlimAccountAllOf) GetHasEntitlements() bool`

GetHasEntitlements returns the HasEntitlements field if non-nil, zero value otherwise.

### GetHasEntitlementsOk

`func (o *SlimAccountAllOf) GetHasEntitlementsOk() (*bool, bool)`

GetHasEntitlementsOk returns a tuple with the HasEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitlements

`func (o *SlimAccountAllOf) SetHasEntitlements(v bool)`

SetHasEntitlements sets HasEntitlements field to given value.

### HasHasEntitlements

`func (o *SlimAccountAllOf) HasHasEntitlements() bool`

HasHasEntitlements returns a boolean if a field has been set.

### GetSourceId

`func (o *SlimAccountAllOf) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SlimAccountAllOf) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SlimAccountAllOf) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SlimAccountAllOf) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *SlimAccountAllOf) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *SlimAccountAllOf) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *SlimAccountAllOf) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *SlimAccountAllOf) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetIdentityId

`func (o *SlimAccountAllOf) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *SlimAccountAllOf) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *SlimAccountAllOf) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *SlimAccountAllOf) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetAttributes

`func (o *SlimAccountAllOf) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SlimAccountAllOf) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SlimAccountAllOf) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SlimAccountAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


