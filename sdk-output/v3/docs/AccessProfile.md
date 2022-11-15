# AccessProfile

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
**Source** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Entitlements** | Pointer to [**[]BaseEntitlement**](BaseEntitlement.md) |  | [optional] 
**EntitlementCount** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccessProfile

`func NewAccessProfile(id string, name string, type_ DocumentType, ) *AccessProfile`

NewAccessProfile instantiates a new AccessProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessProfileWithDefaults

`func NewAccessProfileWithDefaults() *AccessProfile`

NewAccessProfileWithDefaults instantiates a new AccessProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessProfile) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccessProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessProfile) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AccessProfile) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessProfile) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessProfile) SetType(v DocumentType)`

SetType sets Type field to given value.


### GetDescription

`func (o *AccessProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreated

`func (o *AccessProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccessProfile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccessProfile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AccessProfile) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *AccessProfile) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *AccessProfile) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetModified

`func (o *AccessProfile) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AccessProfile) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AccessProfile) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AccessProfile) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *AccessProfile) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *AccessProfile) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *AccessProfile) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *AccessProfile) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *AccessProfile) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *AccessProfile) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *AccessProfile) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *AccessProfile) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetEnabled

`func (o *AccessProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccessProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccessProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccessProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequestable

`func (o *AccessProfile) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *AccessProfile) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *AccessProfile) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *AccessProfile) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetRequestCommentsRequired

`func (o *AccessProfile) GetRequestCommentsRequired() bool`

GetRequestCommentsRequired returns the RequestCommentsRequired field if non-nil, zero value otherwise.

### GetRequestCommentsRequiredOk

`func (o *AccessProfile) GetRequestCommentsRequiredOk() (*bool, bool)`

GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCommentsRequired

`func (o *AccessProfile) SetRequestCommentsRequired(v bool)`

SetRequestCommentsRequired sets RequestCommentsRequired field to given value.

### HasRequestCommentsRequired

`func (o *AccessProfile) HasRequestCommentsRequired() bool`

HasRequestCommentsRequired returns a boolean if a field has been set.

### GetOwner

`func (o *AccessProfile) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccessProfile) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccessProfile) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccessProfile) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSource

`func (o *AccessProfile) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccessProfile) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccessProfile) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *AccessProfile) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEntitlements

`func (o *AccessProfile) GetEntitlements() []BaseEntitlement`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *AccessProfile) GetEntitlementsOk() (*[]BaseEntitlement, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *AccessProfile) SetEntitlements(v []BaseEntitlement)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *AccessProfile) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetEntitlementCount

`func (o *AccessProfile) GetEntitlementCount() int32`

GetEntitlementCount returns the EntitlementCount field if non-nil, zero value otherwise.

### GetEntitlementCountOk

`func (o *AccessProfile) GetEntitlementCountOk() (*int32, bool)`

GetEntitlementCountOk returns a tuple with the EntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCount

`func (o *AccessProfile) SetEntitlementCount(v int32)`

SetEntitlementCount sets EntitlementCount field to given value.

### HasEntitlementCount

`func (o *AccessProfile) HasEntitlementCount() bool`

HasEntitlementCount returns a boolean if a field has been set.

### GetTags

`func (o *AccessProfile) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccessProfile) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccessProfile) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccessProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


