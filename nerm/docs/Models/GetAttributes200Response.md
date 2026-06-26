---
id: nerm-get-attributes200-response
title: GetAttributes200Response
pagination_label: GetAttributes200Response
sidebar_label: GetAttributes200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetAttributes200Response', 'NERMGetAttributes200Response'] 
slug: /tools/sdk/go/nerm/models/get-attributes200-response
tags: ['SDK', 'Software Development Kit', 'GetAttributes200Response', 'NERMGetAttributes200Response']
---

# GetAttributes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttributes** | Pointer to [**[]Attribute**](attribute) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetAttributes200Response

`func NewGetAttributes200Response() *GetAttributes200Response`

NewGetAttributes200Response instantiates a new GetAttributes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttributes200ResponseWithDefaults

`func NewGetAttributes200ResponseWithDefaults() *GetAttributes200Response`

NewGetAttributes200ResponseWithDefaults instantiates a new GetAttributes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttributes

`func (o *GetAttributes200Response) GetNeAttributes() []Attribute`

GetNeAttributes returns the NeAttributes field if non-nil, zero value otherwise.

### GetNeAttributesOk

`func (o *GetAttributes200Response) GetNeAttributesOk() (*[]Attribute, bool)`

GetNeAttributesOk returns a tuple with the NeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributes

`func (o *GetAttributes200Response) SetNeAttributes(v []Attribute)`

SetNeAttributes sets NeAttributes field to given value.

### HasNeAttributes

`func (o *GetAttributes200Response) HasNeAttributes() bool`

HasNeAttributes returns a boolean if a field has been set.

### GetMetadata

`func (o *GetAttributes200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetAttributes200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetAttributes200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetAttributes200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


