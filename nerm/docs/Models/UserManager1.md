---
id: nerm-user-manager1
title: UserManager1
pagination_label: UserManager1
sidebar_label: UserManager1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserManager1', 'NERMUserManager1'] 
slug: /tools/sdk/go/nerm/models/user-manager1
tags: ['SDK', 'Software Development Kit', 'UserManager1', 'NERMUserManager1']
---

# UserManager1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**ManagerId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserManager1

`func NewUserManager1() *UserManager1`

NewUserManager1 instantiates a new UserManager1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserManager1WithDefaults

`func NewUserManager1WithDefaults() *UserManager1`

NewUserManager1WithDefaults instantiates a new UserManager1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserManager1) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserManager1) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserManager1) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserManager1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetManagerId

`func (o *UserManager1) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *UserManager1) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *UserManager1) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *UserManager1) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.


