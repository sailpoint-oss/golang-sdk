---
id: nerm-create-attribute-request
title: CreateAttributeRequest
pagination_label: CreateAttributeRequest
sidebar_label: CreateAttributeRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateAttributeRequest', 'NERMCreateAttributeRequest'] 
slug: /tools/sdk/go/nerm/models/create-attribute-request
tags: ['SDK', 'Software Development Kit', 'CreateAttributeRequest', 'NERMCreateAttributeRequest']
---

# CreateAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttribute** | Pointer to [**Attribute1**](attribute1) |  | [optional] 

## Methods

### NewCreateAttributeRequest

`func NewCreateAttributeRequest() *CreateAttributeRequest`

NewCreateAttributeRequest instantiates a new CreateAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAttributeRequestWithDefaults

`func NewCreateAttributeRequestWithDefaults() *CreateAttributeRequest`

NewCreateAttributeRequestWithDefaults instantiates a new CreateAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttribute

`func (o *CreateAttributeRequest) GetNeAttribute() Attribute1`

GetNeAttribute returns the NeAttribute field if non-nil, zero value otherwise.

### GetNeAttributeOk

`func (o *CreateAttributeRequest) GetNeAttributeOk() (*Attribute1, bool)`

GetNeAttributeOk returns a tuple with the NeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttribute

`func (o *CreateAttributeRequest) SetNeAttribute(v Attribute1)`

SetNeAttribute sets NeAttribute field to given value.

### HasNeAttribute

`func (o *CreateAttributeRequest) HasNeAttribute() bool`

HasNeAttribute returns a boolean if a field has been set.


