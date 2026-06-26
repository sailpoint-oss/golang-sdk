---
id: nerm-create-permission200-response
title: CreatePermission200Response
pagination_label: CreatePermission200Response
sidebar_label: CreatePermission200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreatePermission200Response', 'NERMCreatePermission200Response'] 
slug: /tools/sdk/go/nerm/models/create-permission200-response
tags: ['SDK', 'Software Development Kit', 'CreatePermission200Response', 'NERMCreatePermission200Response']
---

# CreatePermission200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to [**Permission**](permission) |  | [optional] 

## Methods

### NewCreatePermission200Response

`func NewCreatePermission200Response() *CreatePermission200Response`

NewCreatePermission200Response instantiates a new CreatePermission200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePermission200ResponseWithDefaults

`func NewCreatePermission200ResponseWithDefaults() *CreatePermission200Response`

NewCreatePermission200ResponseWithDefaults instantiates a new CreatePermission200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *CreatePermission200Response) GetPermission() Permission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CreatePermission200Response) GetPermissionOk() (*Permission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CreatePermission200Response) SetPermission(v Permission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *CreatePermission200Response) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


