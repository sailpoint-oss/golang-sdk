# RoleDocumentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessProfiles** | Pointer to [**[]Reference1**](Reference1.md) |  | [optional] 
**AccessProfileCount** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRoleDocumentAllOf

`func NewRoleDocumentAllOf() *RoleDocumentAllOf`

NewRoleDocumentAllOf instantiates a new RoleDocumentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDocumentAllOfWithDefaults

`func NewRoleDocumentAllOfWithDefaults() *RoleDocumentAllOf`

NewRoleDocumentAllOfWithDefaults instantiates a new RoleDocumentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessProfiles

`func (o *RoleDocumentAllOf) GetAccessProfiles() []Reference1`

GetAccessProfiles returns the AccessProfiles field if non-nil, zero value otherwise.

### GetAccessProfilesOk

`func (o *RoleDocumentAllOf) GetAccessProfilesOk() (*[]Reference1, bool)`

GetAccessProfilesOk returns a tuple with the AccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfiles

`func (o *RoleDocumentAllOf) SetAccessProfiles(v []Reference1)`

SetAccessProfiles sets AccessProfiles field to given value.

### HasAccessProfiles

`func (o *RoleDocumentAllOf) HasAccessProfiles() bool`

HasAccessProfiles returns a boolean if a field has been set.

### GetAccessProfileCount

`func (o *RoleDocumentAllOf) GetAccessProfileCount() int32`

GetAccessProfileCount returns the AccessProfileCount field if non-nil, zero value otherwise.

### GetAccessProfileCountOk

`func (o *RoleDocumentAllOf) GetAccessProfileCountOk() (*int32, bool)`

GetAccessProfileCountOk returns a tuple with the AccessProfileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileCount

`func (o *RoleDocumentAllOf) SetAccessProfileCount(v int32)`

SetAccessProfileCount sets AccessProfileCount field to given value.

### HasAccessProfileCount

`func (o *RoleDocumentAllOf) HasAccessProfileCount() bool`

HasAccessProfileCount returns a boolean if a field has been set.

### GetTags

`func (o *RoleDocumentAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoleDocumentAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoleDocumentAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoleDocumentAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


