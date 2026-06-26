---
id: nerm-user-profile
title: UserProfile
pagination_label: UserProfile
sidebar_label: UserProfile
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserProfile', 'NERMUserProfile'] 
slug: /tools/sdk/go/nerm/models/user-profile
tags: ['SDK', 'Software Development Kit', 'UserProfile', 'NERMUserProfile']
---

# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**NeAttributeId** | Pointer to **string** |  | [optional] 
**RelationshipType** | Pointer to **string** |  | [optional] 

## Methods

### NewUserProfile

`func NewUserProfile() *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *UserProfile) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserProfile) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserProfile) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UserProfile) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserId

`func (o *UserProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetProfileId

`func (o *UserProfile) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *UserProfile) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *UserProfile) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *UserProfile) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetNeAttributeId

`func (o *UserProfile) GetNeAttributeId() string`

GetNeAttributeId returns the NeAttributeId field if non-nil, zero value otherwise.

### GetNeAttributeIdOk

`func (o *UserProfile) GetNeAttributeIdOk() (*string, bool)`

GetNeAttributeIdOk returns a tuple with the NeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeId

`func (o *UserProfile) SetNeAttributeId(v string)`

SetNeAttributeId sets NeAttributeId field to given value.

### HasNeAttributeId

`func (o *UserProfile) HasNeAttributeId() bool`

HasNeAttributeId returns a boolean if a field has been set.

### GetRelationshipType

`func (o *UserProfile) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *UserProfile) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *UserProfile) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.

### HasRelationshipType

`func (o *UserProfile) HasRelationshipType() bool`

HasRelationshipType returns a boolean if a field has been set.


