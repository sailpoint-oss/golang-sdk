---
id: v1-intelidentitynotfoundbody
title: Intelidentitynotfoundbody
pagination_label: Intelidentitynotfoundbody
sidebar_label: Intelidentitynotfoundbody
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentitynotfoundbody', 'V1Intelidentitynotfoundbody'] 
slug: /tools/sdk/go/intelligence/models/intelidentitynotfoundbody
tags: ['SDK', 'Software Development Kit', 'Intelidentitynotfoundbody', 'V1Intelidentitynotfoundbody']
---

# Intelidentitynotfoundbody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailCode** | **string** | Constant detail code indicating that no identity matched the supplied filter. | 
**Message** | Pointer to **string** | Optional explanatory text describing why no identity was found. | [optional] 

## Methods

### NewIntelidentitynotfoundbody

`func NewIntelidentitynotfoundbody(detailCode string, ) *Intelidentitynotfoundbody`

NewIntelidentitynotfoundbody instantiates a new Intelidentitynotfoundbody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentitynotfoundbodyWithDefaults

`func NewIntelidentitynotfoundbodyWithDefaults() *Intelidentitynotfoundbody`

NewIntelidentitynotfoundbodyWithDefaults instantiates a new Intelidentitynotfoundbody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailCode

`func (o *Intelidentitynotfoundbody) GetDetailCode() string`

GetDetailCode returns the DetailCode field if non-nil, zero value otherwise.

### GetDetailCodeOk

`func (o *Intelidentitynotfoundbody) GetDetailCodeOk() (*string, bool)`

GetDetailCodeOk returns a tuple with the DetailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailCode

`func (o *Intelidentitynotfoundbody) SetDetailCode(v string)`

SetDetailCode sets DetailCode field to given value.


### GetMessage

`func (o *Intelidentitynotfoundbody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Intelidentitynotfoundbody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Intelidentitynotfoundbody) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Intelidentitynotfoundbody) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


