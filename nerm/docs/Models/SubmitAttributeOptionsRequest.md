---
id: nerm-submit-attribute-options-request
title: SubmitAttributeOptionsRequest
pagination_label: SubmitAttributeOptionsRequest
sidebar_label: SubmitAttributeOptionsRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitAttributeOptionsRequest', 'NERMSubmitAttributeOptionsRequest'] 
slug: /tools/sdk/go/nerm/models/submit-attribute-options-request
tags: ['SDK', 'Software Development Kit', 'SubmitAttributeOptionsRequest', 'NERMSubmitAttributeOptionsRequest']
---

# SubmitAttributeOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttributeOptions** | Pointer to [**[]AttributeOption1**](attribute-option1) |  | [optional] 

## Methods

### NewSubmitAttributeOptionsRequest

`func NewSubmitAttributeOptionsRequest() *SubmitAttributeOptionsRequest`

NewSubmitAttributeOptionsRequest instantiates a new SubmitAttributeOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitAttributeOptionsRequestWithDefaults

`func NewSubmitAttributeOptionsRequestWithDefaults() *SubmitAttributeOptionsRequest`

NewSubmitAttributeOptionsRequestWithDefaults instantiates a new SubmitAttributeOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttributeOptions

`func (o *SubmitAttributeOptionsRequest) GetNeAttributeOptions() []AttributeOption1`

GetNeAttributeOptions returns the NeAttributeOptions field if non-nil, zero value otherwise.

### GetNeAttributeOptionsOk

`func (o *SubmitAttributeOptionsRequest) GetNeAttributeOptionsOk() (*[]AttributeOption1, bool)`

GetNeAttributeOptionsOk returns a tuple with the NeAttributeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeOptions

`func (o *SubmitAttributeOptionsRequest) SetNeAttributeOptions(v []AttributeOption1)`

SetNeAttributeOptions sets NeAttributeOptions field to given value.

### HasNeAttributeOptions

`func (o *SubmitAttributeOptionsRequest) HasNeAttributeOptions() bool`

HasNeAttributeOptions returns a boolean if a field has been set.


