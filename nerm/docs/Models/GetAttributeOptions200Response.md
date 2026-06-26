---
id: nerm-get-attribute-options200-response
title: GetAttributeOptions200Response
pagination_label: GetAttributeOptions200Response
sidebar_label: GetAttributeOptions200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetAttributeOptions200Response', 'NERMGetAttributeOptions200Response'] 
slug: /tools/sdk/go/nerm/models/get-attribute-options200-response
tags: ['SDK', 'Software Development Kit', 'GetAttributeOptions200Response', 'NERMGetAttributeOptions200Response']
---

# GetAttributeOptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttributeOptions** | Pointer to [**[]AttributeOption**](attribute-option) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetAttributeOptions200Response

`func NewGetAttributeOptions200Response() *GetAttributeOptions200Response`

NewGetAttributeOptions200Response instantiates a new GetAttributeOptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttributeOptions200ResponseWithDefaults

`func NewGetAttributeOptions200ResponseWithDefaults() *GetAttributeOptions200Response`

NewGetAttributeOptions200ResponseWithDefaults instantiates a new GetAttributeOptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttributeOptions

`func (o *GetAttributeOptions200Response) GetNeAttributeOptions() []AttributeOption`

GetNeAttributeOptions returns the NeAttributeOptions field if non-nil, zero value otherwise.

### GetNeAttributeOptionsOk

`func (o *GetAttributeOptions200Response) GetNeAttributeOptionsOk() (*[]AttributeOption, bool)`

GetNeAttributeOptionsOk returns a tuple with the NeAttributeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeOptions

`func (o *GetAttributeOptions200Response) SetNeAttributeOptions(v []AttributeOption)`

SetNeAttributeOptions sets NeAttributeOptions field to given value.

### HasNeAttributeOptions

`func (o *GetAttributeOptions200Response) HasNeAttributeOptions() bool`

HasNeAttributeOptions returns a boolean if a field has been set.

### GetMetadata

`func (o *GetAttributeOptions200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetAttributeOptions200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetAttributeOptions200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetAttributeOptions200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


