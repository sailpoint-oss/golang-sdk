# BaseAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | The description of the access item | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Synced** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Requestable** | Pointer to **bool** | Indicates if the access can be requested | [optional] 
**RequestCommentsRequired** | Pointer to **bool** | Indicates if comments are required when requesting access | [optional] 
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 

## Methods

### NewBaseAccess

`func NewBaseAccess() *BaseAccess`

NewBaseAccess instantiates a new BaseAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAccessWithDefaults

`func NewBaseAccessWithDefaults() *BaseAccess`

NewBaseAccessWithDefaults instantiates a new BaseAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseAccess) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseAccess) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseAccess) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseAccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseAccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseAccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseAccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BaseAccess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseAccess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseAccess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseAccess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *BaseAccess) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BaseAccess) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BaseAccess) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BaseAccess) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *BaseAccess) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *BaseAccess) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *BaseAccess) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BaseAccess) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BaseAccess) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *BaseAccess) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *BaseAccess) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *BaseAccess) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *BaseAccess) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *BaseAccess) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *BaseAccess) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *BaseAccess) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *BaseAccess) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *BaseAccess) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetEnabled

`func (o *BaseAccess) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseAccess) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseAccess) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseAccess) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequestable

`func (o *BaseAccess) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *BaseAccess) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *BaseAccess) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *BaseAccess) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetRequestCommentsRequired

`func (o *BaseAccess) GetRequestCommentsRequired() bool`

GetRequestCommentsRequired returns the RequestCommentsRequired field if non-nil, zero value otherwise.

### GetRequestCommentsRequiredOk

`func (o *BaseAccess) GetRequestCommentsRequiredOk() (*bool, bool)`

GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCommentsRequired

`func (o *BaseAccess) SetRequestCommentsRequired(v bool)`

SetRequestCommentsRequired sets RequestCommentsRequired field to given value.

### HasRequestCommentsRequired

`func (o *BaseAccess) HasRequestCommentsRequired() bool`

HasRequestCommentsRequired returns a boolean if a field has been set.

### GetOwner

`func (o *BaseAccess) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BaseAccess) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BaseAccess) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BaseAccess) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


