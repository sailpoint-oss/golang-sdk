---
id: nerm-create-permission-request
title: CreatePermissionRequest
pagination_label: CreatePermissionRequest
sidebar_label: CreatePermissionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreatePermissionRequest', 'NERMCreatePermissionRequest'] 
slug: /tools/sdk/go/nerm/models/create-permission-request
tags: ['SDK', 'Software Development Kit', 'CreatePermissionRequest', 'NERMCreatePermissionRequest']
---

# CreatePermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to [**Permission1**](permission1) |  | [optional] 

## Methods

### NewCreatePermissionRequest

`func NewCreatePermissionRequest() *CreatePermissionRequest`

NewCreatePermissionRequest instantiates a new CreatePermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePermissionRequestWithDefaults

`func NewCreatePermissionRequestWithDefaults() *CreatePermissionRequest`

NewCreatePermissionRequestWithDefaults instantiates a new CreatePermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *CreatePermissionRequest) GetPermission() Permission1`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CreatePermissionRequest) GetPermissionOk() (*Permission1, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CreatePermissionRequest) SetPermission(v Permission1)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *CreatePermissionRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


