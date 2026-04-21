---
id: nerm-attribute-options
title: AttributeOptions
pagination_label: AttributeOptions
sidebar_label: AttributeOptions
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AttributeOptions', 'NERMAttributeOptions'] 
slug: /tools/sdk/go/nerm/models/attribute-options
tags: ['SDK', 'Software Development Kit', 'AttributeOptions', 'NERMAttributeOptions']
---

# AttributeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttributeOptions** | Pointer to [**[]AttributeOption**](attribute-option) |  | [optional] 

## Methods

### NewAttributeOptions

`func NewAttributeOptions() *AttributeOptions`

NewAttributeOptions instantiates a new AttributeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeOptionsWithDefaults

`func NewAttributeOptionsWithDefaults() *AttributeOptions`

NewAttributeOptionsWithDefaults instantiates a new AttributeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttributeOptions

`func (o *AttributeOptions) GetNeAttributeOptions() []AttributeOption`

GetNeAttributeOptions returns the NeAttributeOptions field if non-nil, zero value otherwise.

### GetNeAttributeOptionsOk

`func (o *AttributeOptions) GetNeAttributeOptionsOk() (*[]AttributeOption, bool)`

GetNeAttributeOptionsOk returns a tuple with the NeAttributeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeOptions

`func (o *AttributeOptions) SetNeAttributeOptions(v []AttributeOption)`

SetNeAttributeOptions sets NeAttributeOptions field to given value.

### HasNeAttributeOptions

`func (o *AttributeOptions) HasNeAttributeOptions() bool`

HasNeAttributeOptions returns a boolean if a field has been set.


