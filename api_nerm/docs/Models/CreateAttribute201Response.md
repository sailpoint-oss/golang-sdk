---
id: nerm-create-attribute201-response
title: CreateAttribute201Response
pagination_label: CreateAttribute201Response
sidebar_label: CreateAttribute201Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateAttribute201Response', 'NERMCreateAttribute201Response'] 
slug: /tools/sdk/go/nerm/models/create-attribute201-response
tags: ['SDK', 'Software Development Kit', 'CreateAttribute201Response', 'NERMCreateAttribute201Response']
---

# CreateAttribute201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttribute** | Pointer to [**Attribute**](attribute) |  | [optional] 

## Methods

### NewCreateAttribute201Response

`func NewCreateAttribute201Response() *CreateAttribute201Response`

NewCreateAttribute201Response instantiates a new CreateAttribute201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAttribute201ResponseWithDefaults

`func NewCreateAttribute201ResponseWithDefaults() *CreateAttribute201Response`

NewCreateAttribute201ResponseWithDefaults instantiates a new CreateAttribute201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttribute

`func (o *CreateAttribute201Response) GetNeAttribute() Attribute`

GetNeAttribute returns the NeAttribute field if non-nil, zero value otherwise.

### GetNeAttributeOk

`func (o *CreateAttribute201Response) GetNeAttributeOk() (*Attribute, bool)`

GetNeAttributeOk returns a tuple with the NeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttribute

`func (o *CreateAttribute201Response) SetNeAttribute(v Attribute)`

SetNeAttribute sets NeAttribute field to given value.

### HasNeAttribute

`func (o *CreateAttribute201Response) HasNeAttribute() bool`

HasNeAttribute returns a boolean if a field has been set.


