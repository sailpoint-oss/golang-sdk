---
id: nerm-get-roles200-response
title: GetRoles200Response
pagination_label: GetRoles200Response
sidebar_label: GetRoles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetRoles200Response', 'NERMGetRoles200Response'] 
slug: /tools/sdk/go/nerm/models/get-roles200-response
tags: ['SDK', 'Software Development Kit', 'GetRoles200Response', 'NERMGetRoles200Response']
---

# GetRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]Role**](role) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetRoles200Response

`func NewGetRoles200Response() *GetRoles200Response`

NewGetRoles200Response instantiates a new GetRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoles200ResponseWithDefaults

`func NewGetRoles200ResponseWithDefaults() *GetRoles200Response`

NewGetRoles200ResponseWithDefaults instantiates a new GetRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *GetRoles200Response) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetRoles200Response) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetRoles200Response) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GetRoles200Response) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetMetadata

`func (o *GetRoles200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetRoles200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetRoles200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetRoles200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


