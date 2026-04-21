---
id: nerm-get-user-managers200-response
title: GetUserManagers200Response
pagination_label: GetUserManagers200Response
sidebar_label: GetUserManagers200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetUserManagers200Response', 'NERMGetUserManagers200Response'] 
slug: /tools/sdk/go/nerm/models/get-user-managers200-response
tags: ['SDK', 'Software Development Kit', 'GetUserManagers200Response', 'NERMGetUserManagers200Response']
---

# GetUserManagers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserManagers** | Pointer to [**[]UserManager**](user-manager) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetUserManagers200Response

`func NewGetUserManagers200Response() *GetUserManagers200Response`

NewGetUserManagers200Response instantiates a new GetUserManagers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserManagers200ResponseWithDefaults

`func NewGetUserManagers200ResponseWithDefaults() *GetUserManagers200Response`

NewGetUserManagers200ResponseWithDefaults instantiates a new GetUserManagers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserManagers

`func (o *GetUserManagers200Response) GetUserManagers() []UserManager`

GetUserManagers returns the UserManagers field if non-nil, zero value otherwise.

### GetUserManagersOk

`func (o *GetUserManagers200Response) GetUserManagersOk() (*[]UserManager, bool)`

GetUserManagersOk returns a tuple with the UserManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagers

`func (o *GetUserManagers200Response) SetUserManagers(v []UserManager)`

SetUserManagers sets UserManagers field to given value.

### HasUserManagers

`func (o *GetUserManagers200Response) HasUserManagers() bool`

HasUserManagers returns a boolean if a field has been set.

### GetMetadata

`func (o *GetUserManagers200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetUserManagers200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetUserManagers200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetUserManagers200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


