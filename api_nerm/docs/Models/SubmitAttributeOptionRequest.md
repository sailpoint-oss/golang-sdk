---
id: nerm-submit-attribute-option-request
title: SubmitAttributeOptionRequest
pagination_label: SubmitAttributeOptionRequest
sidebar_label: SubmitAttributeOptionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitAttributeOptionRequest', 'NERMSubmitAttributeOptionRequest'] 
slug: /tools/sdk/go/nerm/models/submit-attribute-option-request
tags: ['SDK', 'Software Development Kit', 'SubmitAttributeOptionRequest', 'NERMSubmitAttributeOptionRequest']
---

# SubmitAttributeOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttributeOption** | Pointer to [**AttributeOption1**](attribute-option1) |  | [optional] 

## Methods

### NewSubmitAttributeOptionRequest

`func NewSubmitAttributeOptionRequest() *SubmitAttributeOptionRequest`

NewSubmitAttributeOptionRequest instantiates a new SubmitAttributeOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitAttributeOptionRequestWithDefaults

`func NewSubmitAttributeOptionRequestWithDefaults() *SubmitAttributeOptionRequest`

NewSubmitAttributeOptionRequestWithDefaults instantiates a new SubmitAttributeOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttributeOption

`func (o *SubmitAttributeOptionRequest) GetNeAttributeOption() AttributeOption1`

GetNeAttributeOption returns the NeAttributeOption field if non-nil, zero value otherwise.

### GetNeAttributeOptionOk

`func (o *SubmitAttributeOptionRequest) GetNeAttributeOptionOk() (*AttributeOption1, bool)`

GetNeAttributeOptionOk returns a tuple with the NeAttributeOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeOption

`func (o *SubmitAttributeOptionRequest) SetNeAttributeOption(v AttributeOption1)`

SetNeAttributeOption sets NeAttributeOption field to given value.

### HasNeAttributeOption

`func (o *SubmitAttributeOptionRequest) HasNeAttributeOption() bool`

HasNeAttributeOption returns a boolean if a field has been set.


