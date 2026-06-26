---
id: nerm-get-profile-types200-response
title: GetProfileTypes200Response
pagination_label: GetProfileTypes200Response
sidebar_label: GetProfileTypes200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetProfileTypes200Response', 'NERMGetProfileTypes200Response'] 
slug: /tools/sdk/go/nerm/models/get-profile-types200-response
tags: ['SDK', 'Software Development Kit', 'GetProfileTypes200Response', 'NERMGetProfileTypes200Response']
---

# GetProfileTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypes** | Pointer to [**[]ProfileType**](profile-type) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetProfileTypes200Response

`func NewGetProfileTypes200Response() *GetProfileTypes200Response`

NewGetProfileTypes200Response instantiates a new GetProfileTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfileTypes200ResponseWithDefaults

`func NewGetProfileTypes200ResponseWithDefaults() *GetProfileTypes200Response`

NewGetProfileTypes200ResponseWithDefaults instantiates a new GetProfileTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypes

`func (o *GetProfileTypes200Response) GetProfileTypes() []ProfileType`

GetProfileTypes returns the ProfileTypes field if non-nil, zero value otherwise.

### GetProfileTypesOk

`func (o *GetProfileTypes200Response) GetProfileTypesOk() (*[]ProfileType, bool)`

GetProfileTypesOk returns a tuple with the ProfileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypes

`func (o *GetProfileTypes200Response) SetProfileTypes(v []ProfileType)`

SetProfileTypes sets ProfileTypes field to given value.

### HasProfileTypes

`func (o *GetProfileTypes200Response) HasProfileTypes() bool`

HasProfileTypes returns a boolean if a field has been set.

### GetMetadata

`func (o *GetProfileTypes200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetProfileTypes200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetProfileTypes200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetProfileTypes200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


