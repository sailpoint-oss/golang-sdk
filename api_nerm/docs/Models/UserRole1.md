---
id: nerm-user-role1
title: UserRole1
pagination_label: UserRole1
sidebar_label: UserRole1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserRole1', 'NERMUserRole1'] 
slug: /tools/sdk/go/nerm/models/user-role1
tags: ['SDK', 'Software Development Kit', 'UserRole1', 'NERMUserRole1']
---

# UserRole1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserRole1

`func NewUserRole1() *UserRole1`

NewUserRole1 instantiates a new UserRole1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRole1WithDefaults

`func NewUserRole1WithDefaults() *UserRole1`

NewUserRole1WithDefaults instantiates a new UserRole1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserRole1) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserRole1) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserRole1) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserRole1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRoleId

`func (o *UserRole1) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UserRole1) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UserRole1) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UserRole1) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


