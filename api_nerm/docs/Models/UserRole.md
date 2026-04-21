---
id: nerm-user-role
title: UserRole
pagination_label: UserRole
sidebar_label: UserRole
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserRole', 'NERMUserRole'] 
slug: /tools/sdk/go/nerm/models/user-role
tags: ['SDK', 'Software Development Kit', 'UserRole', 'NERMUserRole']
---

# UserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserRole

`func NewUserRole() *UserRole`

NewUserRole instantiates a new UserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleWithDefaults

`func NewUserRoleWithDefaults() *UserRole`

NewUserRoleWithDefaults instantiates a new UserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *UserRole) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserRole) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserRole) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UserRole) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserId

`func (o *UserRole) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserRole) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserRole) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserRole) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRoleId

`func (o *UserRole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UserRole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UserRole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UserRole) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


