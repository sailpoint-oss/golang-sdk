# BaseAccessAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the access item | [optional] 
**Created** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Synced** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Requestable** | Pointer to **bool** | Indicates if the access can be requested | [optional] 
**RequestCommentsRequired** | Pointer to **bool** | Indicates if comments are required when requesting access | [optional] 
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 

## Methods

### NewBaseAccessAllOf

`func NewBaseAccessAllOf() *BaseAccessAllOf`

NewBaseAccessAllOf instantiates a new BaseAccessAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAccessAllOfWithDefaults

`func NewBaseAccessAllOfWithDefaults() *BaseAccessAllOf`

NewBaseAccessAllOfWithDefaults instantiates a new BaseAccessAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BaseAccessAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseAccessAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseAccessAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseAccessAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *BaseAccessAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BaseAccessAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BaseAccessAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BaseAccessAllOf) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *BaseAccessAllOf) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *BaseAccessAllOf) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *BaseAccessAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *BaseAccessAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *BaseAccessAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *BaseAccessAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *BaseAccessAllOf) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *BaseAccessAllOf) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *BaseAccessAllOf) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *BaseAccessAllOf) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *BaseAccessAllOf) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *BaseAccessAllOf) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *BaseAccessAllOf) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *BaseAccessAllOf) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetEnabled

`func (o *BaseAccessAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseAccessAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseAccessAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseAccessAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequestable

`func (o *BaseAccessAllOf) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *BaseAccessAllOf) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *BaseAccessAllOf) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *BaseAccessAllOf) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetRequestCommentsRequired

`func (o *BaseAccessAllOf) GetRequestCommentsRequired() bool`

GetRequestCommentsRequired returns the RequestCommentsRequired field if non-nil, zero value otherwise.

### GetRequestCommentsRequiredOk

`func (o *BaseAccessAllOf) GetRequestCommentsRequiredOk() (*bool, bool)`

GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCommentsRequired

`func (o *BaseAccessAllOf) SetRequestCommentsRequired(v bool)`

SetRequestCommentsRequired sets RequestCommentsRequired field to given value.

### HasRequestCommentsRequired

`func (o *BaseAccessAllOf) HasRequestCommentsRequired() bool`

HasRequestCommentsRequired returns a boolean if a field has been set.

### GetOwner

`func (o *BaseAccessAllOf) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BaseAccessAllOf) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BaseAccessAllOf) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BaseAccessAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


