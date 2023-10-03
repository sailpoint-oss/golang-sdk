# EntitlementDocumentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modified** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**Synced** | Pointer to **NullableTime** | A date-time in ISO-8601 format | [optional] 
**DisplayName** | Pointer to **string** | The display name of the entitlement | [optional] 
**Source** | Pointer to [**Reference1**](Reference1.md) |  | [optional] 
**Privileged** | Pointer to **bool** |  | [optional] 
**IdentityCount** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEntitlementDocumentAllOf

`func NewEntitlementDocumentAllOf() *EntitlementDocumentAllOf`

NewEntitlementDocumentAllOf instantiates a new EntitlementDocumentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDocumentAllOfWithDefaults

`func NewEntitlementDocumentAllOfWithDefaults() *EntitlementDocumentAllOf`

NewEntitlementDocumentAllOfWithDefaults instantiates a new EntitlementDocumentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModified

`func (o *EntitlementDocumentAllOf) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *EntitlementDocumentAllOf) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *EntitlementDocumentAllOf) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *EntitlementDocumentAllOf) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModifiedNil

`func (o *EntitlementDocumentAllOf) SetModifiedNil(b bool)`

 SetModifiedNil sets the value for Modified to be an explicit nil

### UnsetModified
`func (o *EntitlementDocumentAllOf) UnsetModified()`

UnsetModified ensures that no value is present for Modified, not even an explicit nil
### GetSynced

`func (o *EntitlementDocumentAllOf) GetSynced() time.Time`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *EntitlementDocumentAllOf) GetSyncedOk() (*time.Time, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *EntitlementDocumentAllOf) SetSynced(v time.Time)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *EntitlementDocumentAllOf) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### SetSyncedNil

`func (o *EntitlementDocumentAllOf) SetSyncedNil(b bool)`

 SetSyncedNil sets the value for Synced to be an explicit nil

### UnsetSynced
`func (o *EntitlementDocumentAllOf) UnsetSynced()`

UnsetSynced ensures that no value is present for Synced, not even an explicit nil
### GetDisplayName

`func (o *EntitlementDocumentAllOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EntitlementDocumentAllOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EntitlementDocumentAllOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EntitlementDocumentAllOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *EntitlementDocumentAllOf) GetSource() Reference1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EntitlementDocumentAllOf) GetSourceOk() (*Reference1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EntitlementDocumentAllOf) SetSource(v Reference1)`

SetSource sets Source field to given value.

### HasSource

`func (o *EntitlementDocumentAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPrivileged

`func (o *EntitlementDocumentAllOf) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *EntitlementDocumentAllOf) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *EntitlementDocumentAllOf) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.

### HasPrivileged

`func (o *EntitlementDocumentAllOf) HasPrivileged() bool`

HasPrivileged returns a boolean if a field has been set.

### GetIdentityCount

`func (o *EntitlementDocumentAllOf) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *EntitlementDocumentAllOf) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *EntitlementDocumentAllOf) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *EntitlementDocumentAllOf) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetTags

`func (o *EntitlementDocumentAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntitlementDocumentAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntitlementDocumentAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntitlementDocumentAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


