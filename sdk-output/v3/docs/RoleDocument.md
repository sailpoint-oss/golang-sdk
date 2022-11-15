# RoleDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Type** | [**DocumentType**](DocumentType.md) |  | 
**Description** | Pointer to **string** | The description of the access item | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Synced** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Requestable** | Pointer to **bool** | Indicates if the access can be requested | [optional] 
**RequestCommentsRequired** | Pointer to **bool** | Indicates if comments are required when requesting access | [optional] 
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 
**AccessProfiles** | Pointer to [**[]Reference1**](Reference1.md) |  | [optional] 
**AccessProfileCount** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRoleDocument

`func NewRoleDocument(id string, name string, type_ DocumentType, ) *RoleDocument`

NewRoleDocument instantiates a new RoleDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDocumentWithDefaults

`func NewRoleDocumentWithDefaults() *RoleDocument`

NewRoleDocumentWithDefaults instantiates a new RoleDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleDocument) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RoleDocument) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleDocument) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleDocument) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RoleDocument) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleDocument) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleDocument) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RoleDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *RoleDocument) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RoleDocument) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RoleDocument) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RoleDocument) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *RoleDocument) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *RoleDocument) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *RoleDocument) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RoleDocument) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RoleDocument) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *RoleDocument) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *RoleDocument) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *RoleDocument) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *RoleDocument) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *RoleDocument) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *RoleDocument) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *RoleDocument) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *RoleDocument) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *RoleDocument) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetEnabled

`func (o *RoleDocument) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RoleDocument) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RoleDocument) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RoleDocument) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequestable

`func (o *RoleDocument) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *RoleDocument) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *RoleDocument) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *RoleDocument) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetRequestCommentsRequired

`func (o *RoleDocument) GetRequestCommentsRequired() bool`

GetRequestCommentsRequired returns the RequestCommentsRequired field if non-nil, zero value otherwise.

### GetRequestCommentsRequiredOk

`func (o *RoleDocument) GetRequestCommentsRequiredOk() (*bool, bool)`

GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCommentsRequired

`func (o *RoleDocument) SetRequestCommentsRequired(v bool)`

SetRequestCommentsRequired sets RequestCommentsRequired field to given value.

### HasRequestCommentsRequired

`func (o *RoleDocument) HasRequestCommentsRequired() bool`

HasRequestCommentsRequired returns a boolean if a field has been set.

### GetOwner

`func (o *RoleDocument) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RoleDocument) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RoleDocument) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *RoleDocument) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccessProfiles

`func (o *RoleDocument) GetAccessProfiles() []Reference1`

GetAccessProfiles returns the AccessProfiles field if non-nil, zero value otherwise.

### GetAccessProfilesOk

`func (o *RoleDocument) GetAccessProfilesOk() (*[]Reference1, bool)`

GetAccessProfilesOk returns a tuple with the AccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfiles

`func (o *RoleDocument) SetAccessProfiles(v []Reference1)`

SetAccessProfiles sets AccessProfiles field to given value.

### HasAccessProfiles

`func (o *RoleDocument) HasAccessProfiles() bool`

HasAccessProfiles returns a boolean if a field has been set.

### GetAccessProfileCount

`func (o *RoleDocument) GetAccessProfileCount() int32`

GetAccessProfileCount returns the AccessProfileCount field if non-nil, zero value otherwise.

### GetAccessProfileCountOk

`func (o *RoleDocument) GetAccessProfileCountOk() (*int32, bool)`

GetAccessProfileCountOk returns a tuple with the AccessProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileCount

`func (o *RoleDocument) SetAccessProfileCount(v int32)`

SetAccessProfileCount sets AccessProfileCount field to given value.

### HasAccessProfileCount

`func (o *RoleDocument) HasAccessProfileCount() bool`

HasAccessProfileCount returns a boolean if a field has been set.

### GetTags

`func (o *RoleDocument) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoleDocument) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoleDocument) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoleDocument) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


