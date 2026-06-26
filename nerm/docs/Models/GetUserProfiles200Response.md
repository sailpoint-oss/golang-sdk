---
id: nerm-get-user-profiles200-response
title: GetUserProfiles200Response
pagination_label: GetUserProfiles200Response
sidebar_label: GetUserProfiles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetUserProfiles200Response', 'NERMGetUserProfiles200Response'] 
slug: /tools/sdk/go/nerm/models/get-user-profiles200-response
tags: ['SDK', 'Software Development Kit', 'GetUserProfiles200Response', 'NERMGetUserProfiles200Response']
---

# GetUserProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserProfiles** | Pointer to [**[]UserProfile**](user-profile) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetUserProfiles200Response

`func NewGetUserProfiles200Response() *GetUserProfiles200Response`

NewGetUserProfiles200Response instantiates a new GetUserProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserProfiles200ResponseWithDefaults

`func NewGetUserProfiles200ResponseWithDefaults() *GetUserProfiles200Response`

NewGetUserProfiles200ResponseWithDefaults instantiates a new GetUserProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserProfiles

`func (o *GetUserProfiles200Response) GetUserProfiles() []UserProfile`

GetUserProfiles returns the UserProfiles field if non-nil, zero value otherwise.

### GetUserProfilesOk

`func (o *GetUserProfiles200Response) GetUserProfilesOk() (*[]UserProfile, bool)`

GetUserProfilesOk returns a tuple with the UserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfiles

`func (o *GetUserProfiles200Response) SetUserProfiles(v []UserProfile)`

SetUserProfiles sets UserProfiles field to given value.

### HasUserProfiles

`func (o *GetUserProfiles200Response) HasUserProfiles() bool`

HasUserProfiles returns a boolean if a field has been set.

### GetMetadata

`func (o *GetUserProfiles200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetUserProfiles200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetUserProfiles200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetUserProfiles200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


