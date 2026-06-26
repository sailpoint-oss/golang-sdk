---
id: nerm-get-role-profiles200-response
title: GetRoleProfiles200Response
pagination_label: GetRoleProfiles200Response
sidebar_label: GetRoleProfiles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetRoleProfiles200Response', 'NERMGetRoleProfiles200Response'] 
slug: /tools/sdk/go/nerm/models/get-role-profiles200-response
tags: ['SDK', 'Software Development Kit', 'GetRoleProfiles200Response', 'NERMGetRoleProfiles200Response']
---

# GetRoleProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleProfiles** | Pointer to [**[]RoleProfile**](role-profile) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetRoleProfiles200Response

`func NewGetRoleProfiles200Response() *GetRoleProfiles200Response`

NewGetRoleProfiles200Response instantiates a new GetRoleProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoleProfiles200ResponseWithDefaults

`func NewGetRoleProfiles200ResponseWithDefaults() *GetRoleProfiles200Response`

NewGetRoleProfiles200ResponseWithDefaults instantiates a new GetRoleProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleProfiles

`func (o *GetRoleProfiles200Response) GetRoleProfiles() []RoleProfile`

GetRoleProfiles returns the RoleProfiles field if non-nil, zero value otherwise.

### GetRoleProfilesOk

`func (o *GetRoleProfiles200Response) GetRoleProfilesOk() (*[]RoleProfile, bool)`

GetRoleProfilesOk returns a tuple with the RoleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleProfiles

`func (o *GetRoleProfiles200Response) SetRoleProfiles(v []RoleProfile)`

SetRoleProfiles sets RoleProfiles field to given value.

### HasRoleProfiles

`func (o *GetRoleProfiles200Response) HasRoleProfiles() bool`

HasRoleProfiles returns a boolean if a field has been set.

### GetMetadata

`func (o *GetRoleProfiles200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetRoleProfiles200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetRoleProfiles200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetRoleProfiles200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


