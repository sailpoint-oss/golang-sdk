---
id: nerm-system-role-permission1
title: SystemRolePermission1
pagination_label: SystemRolePermission1
sidebar_label: SystemRolePermission1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SystemRolePermission1', 'NERMSystemRolePermission1'] 
slug: /tools/sdk/go/nerm/models/system-role-permission1
tags: ['SDK', 'Software Development Kit', 'SystemRolePermission1', 'NERMSystemRolePermission1']
---

# SystemRolePermission1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemRoleId** | Pointer to **string** | The id of the system role | [optional] 
**SubjectId** | Pointer to **string** | The ID of the object that the permission is giving access to | [optional] 
**Value** | Pointer to **int32** | The permissions level of access | [optional] 
**Subject** | Pointer to **int32** | The type of permission | [optional] 

## Methods

### NewSystemRolePermission1

`func NewSystemRolePermission1() *SystemRolePermission1`

NewSystemRolePermission1 instantiates a new SystemRolePermission1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemRolePermission1WithDefaults

`func NewSystemRolePermission1WithDefaults() *SystemRolePermission1`

NewSystemRolePermission1WithDefaults instantiates a new SystemRolePermission1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemRoleId

`func (o *SystemRolePermission1) GetSystemRoleId() string`

GetSystemRoleId returns the SystemRoleId field if non-nil, zero value otherwise.

### GetSystemRoleIdOk

`func (o *SystemRolePermission1) GetSystemRoleIdOk() (*string, bool)`

GetSystemRoleIdOk returns a tuple with the SystemRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRoleId

`func (o *SystemRolePermission1) SetSystemRoleId(v string)`

SetSystemRoleId sets SystemRoleId field to given value.

### HasSystemRoleId

`func (o *SystemRolePermission1) HasSystemRoleId() bool`

HasSystemRoleId returns a boolean if a field has been set.

### GetSubjectId

`func (o *SystemRolePermission1) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SystemRolePermission1) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SystemRolePermission1) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *SystemRolePermission1) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetValue

`func (o *SystemRolePermission1) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemRolePermission1) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemRolePermission1) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *SystemRolePermission1) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSubject

`func (o *SystemRolePermission1) GetSubject() int32`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SystemRolePermission1) GetSubjectOk() (*int32, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SystemRolePermission1) SetSubject(v int32)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SystemRolePermission1) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


