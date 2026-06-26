---
id: nerm-user-roles
title: UserRoles
pagination_label: UserRoles
sidebar_label: UserRoles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserRoles', 'NERMUserRoles'] 
slug: /tools/sdk/go/nerm/models/user-roles
tags: ['SDK', 'Software Development Kit', 'UserRoles', 'NERMUserRoles']
---

# UserRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRoles** | Pointer to [**[]UserRole**](user-role) |  | [optional] 

## Methods

### NewUserRoles

`func NewUserRoles() *UserRoles`

NewUserRoles instantiates a new UserRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRolesWithDefaults

`func NewUserRolesWithDefaults() *UserRoles`

NewUserRolesWithDefaults instantiates a new UserRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRoles

`func (o *UserRoles) GetUserRoles() []UserRole`

GetUserRoles returns the UserRoles field if non-nil, zero value otherwise.

### GetUserRolesOk

`func (o *UserRoles) GetUserRolesOk() (*[]UserRole, bool)`

GetUserRolesOk returns a tuple with the UserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoles

`func (o *UserRoles) SetUserRoles(v []UserRole)`

SetUserRoles sets UserRoles field to given value.

### HasUserRoles

`func (o *UserRoles) HasUserRoles() bool`

HasUserRoles returns a boolean if a field has been set.


