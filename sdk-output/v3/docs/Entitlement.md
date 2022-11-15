# Entitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
**Description** | Pointer to **string** | A description of the entitlement | [optional] 
**Attribute** | Pointer to **string** | The name of the entitlement attribute | [optional] 
**Value** | Pointer to **string** | The value of the entitlement | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Synced** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**DisplayName** | Pointer to **string** | The display name of the entitlement | [optional] 
**Source** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**IdentityCount** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEntitlement

`func NewEntitlement(id string, name string, type_ DocumentType, ) *Entitlement`

NewEntitlement instantiates a new Entitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementWithDefaults

`func NewEntitlementWithDefaults() *Entitlement`

NewEntitlementWithDefaults instantiates a new Entitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entitlement) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Entitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entitlement) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Entitlement) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entitlement) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entitlement) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetDescription

`func (o *Entitlement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Entitlement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Entitlement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Entitlement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttribute

`func (o *Entitlement) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Entitlement) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Entitlement) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Entitlement) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *Entitlement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Entitlement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Entitlement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Entitlement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetModified

`func (o *Entitlement) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Entitlement) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Entitlement) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Entitlement) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *Entitlement) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *Entitlement) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *Entitlement) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *Entitlement) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *Entitlement) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *Entitlement) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *Entitlement) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *Entitlement) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetDisplayName

`func (o *Entitlement) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Entitlement) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Entitlement) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Entitlement) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *Entitlement) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Entitlement) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Entitlement) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *Entitlement) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrivileged

`func (o *Entitlement) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *Entitlement) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *Entitlement) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *Entitlement) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetIdentityCount

`func (o *Entitlement) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *Entitlement) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *Entitlement) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *Entitlement) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetTags

`func (o *Entitlement) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Entitlement) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Entitlement) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Entitlement) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


