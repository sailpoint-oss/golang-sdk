# EntitlementDocument

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

### NewEntitlementDocument

`func NewEntitlementDocument(id string, name string, type_ DocumentType, ) *EntitlementDocument`

NewEntitlementDocument instantiates a new EntitlementDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDocumentWithDefaults

`func NewEntitlementDocumentWithDefaults() *EntitlementDocument`

NewEntitlementDocumentWithDefaults instantiates a new EntitlementDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntitlementDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementDocument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EntitlementDocument) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementDocument) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementDocument) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetDescription

`func (o *EntitlementDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttribute

`func (o *EntitlementDocument) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *EntitlementDocument) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *EntitlementDocument) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *EntitlementDocument) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetValue

`func (o *EntitlementDocument) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntitlementDocument) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntitlementDocument) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntitlementDocument) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetModified

`func (o *EntitlementDocument) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *EntitlementDocument) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *EntitlementDocument) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *EntitlementDocument) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *EntitlementDocument) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *EntitlementDocument) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *EntitlementDocument) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *EntitlementDocument) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *EntitlementDocument) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *EntitlementDocument) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *EntitlementDocument) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *EntitlementDocument) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetDisplayName

`func (o *EntitlementDocument) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EntitlementDocument) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EntitlementDocument) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EntitlementDocument) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *EntitlementDocument) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EntitlementDocument) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EntitlementDocument) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *EntitlementDocument) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrivileged

`func (o *EntitlementDocument) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *EntitlementDocument) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *EntitlementDocument) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *EntitlementDocument) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetIdentityCount

`func (o *EntitlementDocument) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *EntitlementDocument) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *EntitlementDocument) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *EntitlementDocument) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetTags

`func (o *EntitlementDocument) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntitlementDocument) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntitlementDocument) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntitlementDocument) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


