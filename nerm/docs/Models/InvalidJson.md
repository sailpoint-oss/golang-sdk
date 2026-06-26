---
id: nerm-invalid-json
title: InvalidJson
pagination_label: InvalidJson
sidebar_label: InvalidJson
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'InvalidJson', 'NERMInvalidJson'] 
slug: /tools/sdk/go/nerm/models/invalid-json
tags: ['SDK', 'Software Development Kit', 'InvalidJson', 'NERMInvalidJson']
---

# InvalidJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInvalidJson

`func NewInvalidJson() *InvalidJson`

NewInvalidJson instantiates a new InvalidJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidJsonWithDefaults

`func NewInvalidJsonWithDefaults() *InvalidJson`

NewInvalidJsonWithDefaults instantiates a new InvalidJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *InvalidJson) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InvalidJson) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InvalidJson) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *InvalidJson) HasError() bool`

HasError returns a boolean if a field has been set.


