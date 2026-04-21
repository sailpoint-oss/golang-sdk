---
id: nerm-system-role
title: SystemRole
pagination_label: SystemRole
sidebar_label: SystemRole
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SystemRole', 'NERMSystemRole'] 
slug: /tools/sdk/go/nerm/models/system-role
tags: ['SDK', 'Software Development Kit', 'SystemRole', 'NERMSystemRole']
---

# SystemRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the object | [optional] [readonly] 
**Uid** | Pointer to **string** | The user identifier for the object | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the role | [optional] 

## Methods

### NewSystemRole

`func NewSystemRole() *SystemRole`

NewSystemRole instantiates a new SystemRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemRoleWithDefaults

`func NewSystemRoleWithDefaults() *SystemRole`

NewSystemRoleWithDefaults instantiates a new SystemRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SystemRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *SystemRole) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SystemRole) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SystemRole) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SystemRole) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *SystemRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemRole) HasName() bool`

HasName returns a boolean if a field has been set.


