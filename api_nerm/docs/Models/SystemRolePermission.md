---
id: nerm-system-role-permission
title: SystemRolePermission
pagination_label: SystemRolePermission
sidebar_label: SystemRolePermission
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SystemRolePermission', 'NERMSystemRolePermission'] 
slug: /tools/sdk/go/nerm/models/system-role-permission
tags: ['SDK', 'Software Development Kit', 'SystemRolePermission', 'NERMSystemRolePermission']
---

# SystemRolePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the system role permission | [optional] [readonly] 
**SystemRoleId** | Pointer to **string** | The id of the system role | [optional] 
**Value** | Pointer to **int32** | The permissions level of access | [optional] 
**Subject** | Pointer to **int32** | The type of permission | [optional] 
**SubjectId** | Pointer to **string** | The ID of the object that the permission is giving access to | [optional] 

## Methods

### NewSystemRolePermission

`func NewSystemRolePermission() *SystemRolePermission`

NewSystemRolePermission instantiates a new SystemRolePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemRolePermissionWithDefaults

`func NewSystemRolePermissionWithDefaults() *SystemRolePermission`

NewSystemRolePermissionWithDefaults instantiates a new SystemRolePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemRolePermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemRolePermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemRolePermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SystemRolePermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSystemRoleId

`func (o *SystemRolePermission) GetSystemRoleId() string`

GetSystemRoleId returns the SystemRoleId field if non-nil, zero value otherwise.

### GetSystemRoleIdOk

`func (o *SystemRolePermission) GetSystemRoleIdOk() (*string, bool)`

GetSystemRoleIdOk returns a tuple with the SystemRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRoleId

`func (o *SystemRolePermission) SetSystemRoleId(v string)`

SetSystemRoleId sets SystemRoleId field to given value.

### HasSystemRoleId

`func (o *SystemRolePermission) HasSystemRoleId() bool`

HasSystemRoleId returns a boolean if a field has been set.

### GetValue

`func (o *SystemRolePermission) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemRolePermission) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemRolePermission) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *SystemRolePermission) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSubject

`func (o *SystemRolePermission) GetSubject() int32`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SystemRolePermission) GetSubjectOk() (*int32, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SystemRolePermission) SetSubject(v int32)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SystemRolePermission) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectId

`func (o *SystemRolePermission) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SystemRolePermission) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SystemRolePermission) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *SystemRolePermission) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.


