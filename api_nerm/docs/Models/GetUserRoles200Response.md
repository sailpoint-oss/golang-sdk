---
id: nerm-get-user-roles200-response
title: GetUserRoles200Response
pagination_label: GetUserRoles200Response
sidebar_label: GetUserRoles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetUserRoles200Response', 'NERMGetUserRoles200Response'] 
slug: /tools/sdk/go/nerm/models/get-user-roles200-response
tags: ['SDK', 'Software Development Kit', 'GetUserRoles200Response', 'NERMGetUserRoles200Response']
---

# GetUserRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRoles** | Pointer to [**[]UserRole**](user-role) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetUserRoles200Response

`func NewGetUserRoles200Response() *GetUserRoles200Response`

NewGetUserRoles200Response instantiates a new GetUserRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserRoles200ResponseWithDefaults

`func NewGetUserRoles200ResponseWithDefaults() *GetUserRoles200Response`

NewGetUserRoles200ResponseWithDefaults instantiates a new GetUserRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRoles

`func (o *GetUserRoles200Response) GetUserRoles() []UserRole`

GetUserRoles returns the UserRoles field if non-nil, zero value otherwise.

### GetUserRolesOk

`func (o *GetUserRoles200Response) GetUserRolesOk() (*[]UserRole, bool)`

GetUserRolesOk returns a tuple with the UserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoles

`func (o *GetUserRoles200Response) SetUserRoles(v []UserRole)`

SetUserRoles sets UserRoles field to given value.

### HasUserRoles

`func (o *GetUserRoles200Response) HasUserRoles() bool`

HasUserRoles returns a boolean if a field has been set.

### GetMetadata

`func (o *GetUserRoles200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetUserRoles200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetUserRoles200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetUserRoles200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


