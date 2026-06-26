---
id: nerm-user-managers
title: UserManagers
pagination_label: UserManagers
sidebar_label: UserManagers
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserManagers', 'NERMUserManagers'] 
slug: /tools/sdk/go/nerm/models/user-managers
tags: ['SDK', 'Software Development Kit', 'UserManagers', 'NERMUserManagers']
---

# UserManagers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserManagers** | Pointer to [**[]UserManager**](user-manager) |  | [optional] 

## Methods

### NewUserManagers

`func NewUserManagers() *UserManagers`

NewUserManagers instantiates a new UserManagers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserManagersWithDefaults

`func NewUserManagersWithDefaults() *UserManagers`

NewUserManagersWithDefaults instantiates a new UserManagers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserManagers

`func (o *UserManagers) GetUserManagers() []UserManager`

GetUserManagers returns the UserManagers field if non-nil, zero value otherwise.

### GetUserManagersOk

`func (o *UserManagers) GetUserManagersOk() (*[]UserManager, bool)`

GetUserManagersOk returns a tuple with the UserManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagers

`func (o *UserManagers) SetUserManagers(v []UserManager)`

SetUserManagers sets UserManagers field to given value.

### HasUserManagers

`func (o *UserManagers) HasUserManagers() bool`

HasUserManagers returns a boolean if a field has been set.


