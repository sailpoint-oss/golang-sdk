---
id: nerm-get-system-roles200-response
title: GetSystemRoles200Response
pagination_label: GetSystemRoles200Response
sidebar_label: GetSystemRoles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetSystemRoles200Response', 'NERMGetSystemRoles200Response'] 
slug: /tools/sdk/go/nerm/models/get-system-roles200-response
tags: ['SDK', 'Software Development Kit', 'GetSystemRoles200Response', 'NERMGetSystemRoles200Response']
---

# GetSystemRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemRoles** | Pointer to [**[]SystemRole**](system-role) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetSystemRoles200Response

`func NewGetSystemRoles200Response() *GetSystemRoles200Response`

NewGetSystemRoles200Response instantiates a new GetSystemRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSystemRoles200ResponseWithDefaults

`func NewGetSystemRoles200ResponseWithDefaults() *GetSystemRoles200Response`

NewGetSystemRoles200ResponseWithDefaults instantiates a new GetSystemRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemRoles

`func (o *GetSystemRoles200Response) GetSystemRoles() []SystemRole`

GetSystemRoles returns the SystemRoles field if non-nil, zero value otherwise.

### GetSystemRolesOk

`func (o *GetSystemRoles200Response) GetSystemRolesOk() (*[]SystemRole, bool)`

GetSystemRolesOk returns a tuple with the SystemRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemRoles

`func (o *GetSystemRoles200Response) SetSystemRoles(v []SystemRole)`

SetSystemRoles sets SystemRoles field to given value.

### HasSystemRoles

`func (o *GetSystemRoles200Response) HasSystemRoles() bool`

HasSystemRoles returns a boolean if a field has been set.

### GetMetadata

`func (o *GetSystemRoles200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSystemRoles200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSystemRoles200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetSystemRoles200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


