---
id: nerm-user-profiles
title: UserProfiles
pagination_label: UserProfiles
sidebar_label: UserProfiles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserProfiles', 'NERMUserProfiles'] 
slug: /tools/sdk/go/nerm/models/user-profiles
tags: ['SDK', 'Software Development Kit', 'UserProfiles', 'NERMUserProfiles']
---

# UserProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserProfiles** | Pointer to [**[]UserProfile**](user-profile) |  | [optional] 

## Methods

### NewUserProfiles

`func NewUserProfiles() *UserProfiles`

NewUserProfiles instantiates a new UserProfiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfilesWithDefaults

`func NewUserProfilesWithDefaults() *UserProfiles`

NewUserProfilesWithDefaults instantiates a new UserProfiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserProfiles

`func (o *UserProfiles) GetUserProfiles() []UserProfile`

GetUserProfiles returns the UserProfiles field if non-nil, zero value otherwise.

### GetUserProfilesOk

`func (o *UserProfiles) GetUserProfilesOk() (*[]UserProfile, bool)`

GetUserProfilesOk returns a tuple with the UserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfiles

`func (o *UserProfiles) SetUserProfiles(v []UserProfile)`

SetUserProfiles sets UserProfiles field to given value.

### HasUserProfiles

`func (o *UserProfiles) HasUserProfiles() bool`

HasUserProfiles returns a boolean if a field has been set.


