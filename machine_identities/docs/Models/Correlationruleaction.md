---
id: v1-correlationruleaction
title: Correlationruleaction
pagination_label: Correlationruleaction
sidebar_label: Correlationruleaction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Correlationruleaction', 'V1Correlationruleaction'] 
slug: /tools/sdk/go/machineidentities/models/correlationruleaction
tags: ['SDK', 'Software Development Kit', 'Correlationruleaction', 'V1Correlationruleaction']
---

# Correlationruleaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The target owner type resolved by this action. | 
**Payload** | Pointer to **map[string]interface{}** | Action-specific payload. | [optional] 

## Methods

### NewCorrelationruleaction

`func NewCorrelationruleaction(type_ string, ) *Correlationruleaction`

NewCorrelationruleaction instantiates a new Correlationruleaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelationruleactionWithDefaults

`func NewCorrelationruleactionWithDefaults() *Correlationruleaction`

NewCorrelationruleactionWithDefaults instantiates a new Correlationruleaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Correlationruleaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Correlationruleaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Correlationruleaction) SetType(v string)`

SetType sets Type field to given value.


### GetPayload

`func (o *Correlationruleaction) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Correlationruleaction) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Correlationruleaction) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Correlationruleaction) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


