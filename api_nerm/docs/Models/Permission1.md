---
id: nerm-permission1
title: Permission1
pagination_label: Permission1
sidebar_label: Permission1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Permission1', 'NERMPermission1'] 
slug: /tools/sdk/go/nerm/models/permission1
tags: ['SDK', 'Software Development Kit', 'Permission1', 'NERMPermission1']
---

# Permission1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | Pointer to **string** | The id of the role | [optional] 
**SubjectId** | Pointer to **string** | The ID of the object that the permission is giving access to | [optional] 
**Value** | Pointer to **int32** | The permissions level of access | [optional] 
**Subject** | Pointer to **int32** | The type of permission | [optional] 

## Methods

### NewPermission1

`func NewPermission1() *Permission1`

NewPermission1 instantiates a new Permission1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermission1WithDefaults

`func NewPermission1WithDefaults() *Permission1`

NewPermission1WithDefaults instantiates a new Permission1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *Permission1) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *Permission1) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *Permission1) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *Permission1) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetSubjectId

`func (o *Permission1) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *Permission1) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *Permission1) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *Permission1) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetValue

`func (o *Permission1) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Permission1) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Permission1) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Permission1) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSubject

`func (o *Permission1) GetSubject() int32`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Permission1) GetSubjectOk() (*int32, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Permission1) SetSubject(v int32)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Permission1) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


