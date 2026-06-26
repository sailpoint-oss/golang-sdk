---
id: nerm-users
title: Users
pagination_label: Users
sidebar_label: Users
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Users', 'NERMUsers'] 
slug: /tools/sdk/go/nerm/models/users
tags: ['SDK', 'Software Development Kit', 'Users', 'NERMUsers']
---

# Users

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]User**](user) |  | [optional] 

## Methods

### NewUsers

`func NewUsers() *Users`

NewUsers instantiates a new Users object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersWithDefaults

`func NewUsersWithDefaults() *Users`

NewUsersWithDefaults instantiates a new Users object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *Users) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Users) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Users) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Users) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


