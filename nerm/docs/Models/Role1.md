---
id: nerm-role1
title: Role1
pagination_label: Role1
sidebar_label: Role1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Role1', 'NERMRole1'] 
slug: /tools/sdk/go/nerm/models/role1
tags: ['SDK', 'Software Development Kit', 'Role1', 'NERMRole1']
---

# Role1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRole1

`func NewRole1() *Role1`

NewRole1 instantiates a new Role1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRole1WithDefaults

`func NewRole1WithDefaults() *Role1`

NewRole1WithDefaults instantiates a new Role1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Role1) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Role1) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Role1) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Role1) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *Role1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Role1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Role1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Role1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Role1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Role1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroups

`func (o *Role1) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Role1) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Role1) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Role1) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


