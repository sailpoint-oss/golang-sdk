---
id: nerm-user-manager
title: UserManager
pagination_label: UserManager
sidebar_label: UserManager
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserManager', 'NERMUserManager'] 
slug: /tools/sdk/go/nerm/models/user-manager
tags: ['SDK', 'Software Development Kit', 'UserManager', 'NERMUserManager']
---

# UserManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] 
**ManagerId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserManager

`func NewUserManager() *UserManager`

NewUserManager instantiates a new UserManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserManagerWithDefaults

`func NewUserManagerWithDefaults() *UserManager`

NewUserManagerWithDefaults instantiates a new UserManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserManager) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserManager) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserManager) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserManager) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *UserManager) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserManager) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserManager) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UserManager) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserId

`func (o *UserManager) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserManager) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserManager) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserManager) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetManagerId

`func (o *UserManager) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *UserManager) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *UserManager) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *UserManager) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.


