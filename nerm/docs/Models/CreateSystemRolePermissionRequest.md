---
id: nerm-create-system-role-permission-request
title: CreateSystemRolePermissionRequest
pagination_label: CreateSystemRolePermissionRequest
sidebar_label: CreateSystemRolePermissionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateSystemRolePermissionRequest', 'NERMCreateSystemRolePermissionRequest'] 
slug: /tools/sdk/go/nerm/models/create-system-role-permission-request
tags: ['SDK', 'Software Development Kit', 'CreateSystemRolePermissionRequest', 'NERMCreateSystemRolePermissionRequest']
---

# CreateSystemRolePermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemRolePermission** | Pointer to [**SystemRolePermission1**](system-role-permission1) |  | [optional] 

## Methods

### NewCreateSystemRolePermissionRequest

`func NewCreateSystemRolePermissionRequest() *CreateSystemRolePermissionRequest`

NewCreateSystemRolePermissionRequest instantiates a new CreateSystemRolePermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSystemRolePermissionRequestWithDefaults

`func NewCreateSystemRolePermissionRequestWithDefaults() *CreateSystemRolePermissionRequest`

NewCreateSystemRolePermissionRequestWithDefaults instantiates a new CreateSystemRolePermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemRolePermission

`func (o *CreateSystemRolePermissionRequest) GetSystemRolePermission() SystemRolePermission1`

GetSystemRolePermission returns the SystemRolePermission field if non-nil, zero value otherwise.

### GetSystemRolePermissionOk

`func (o *CreateSystemRolePermissionRequest) GetSystemRolePermissionOk() (*SystemRolePermission1, bool)`

GetSystemRolePermissionOk returns a tuple with the SystemRolePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRolePermission

`func (o *CreateSystemRolePermissionRequest) SetSystemRolePermission(v SystemRolePermission1)`

SetSystemRolePermission sets SystemRolePermission field to given value.

### HasSystemRolePermission

`func (o *CreateSystemRolePermissionRequest) HasSystemRolePermission() bool`

HasSystemRolePermission returns a boolean if a field has been set.


