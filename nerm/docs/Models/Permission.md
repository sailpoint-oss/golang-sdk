---
id: nerm-permission
title: Permission
pagination_label: Permission
sidebar_label: Permission
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Permission', 'NERMPermission'] 
slug: /tools/sdk/go/nerm/models/permission
tags: ['SDK', 'Software Development Kit', 'Permission', 'NERMPermission']
---

# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the permission | [optional] [readonly] 
**RoleId** | Pointer to **string** | The id of the role | [optional] 
**Value** | Pointer to **int32** | The permissions level of access | [optional] 
**Subject** | Pointer to **int32** | The type of permission | [optional] 
**SubjectId** | Pointer to **string** | The ID of the object that the permission is giving access to | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Permission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Permission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Permission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Permission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoleId

`func (o *Permission) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *Permission) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *Permission) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *Permission) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetValue

`func (o *Permission) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Permission) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Permission) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Permission) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSubject

`func (o *Permission) GetSubject() int32`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Permission) GetSubjectOk() (*int32, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Permission) SetSubject(v int32)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Permission) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectId

`func (o *Permission) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *Permission) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *Permission) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *Permission) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.


