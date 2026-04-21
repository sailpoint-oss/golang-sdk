---
id: nerm-get-attributes500-response
title: GetAttributes500Response
pagination_label: GetAttributes500Response
sidebar_label: GetAttributes500Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetAttributes500Response', 'NERMGetAttributes500Response'] 
slug: /tools/sdk/go/nerm/models/get-attributes500-response
tags: ['SDK', 'Software Development Kit', 'GetAttributes500Response', 'NERMGetAttributes500Response']
---

# GetAttributes500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** | A message describing the error | [optional] 

## Methods

### NewGetAttributes500Response

`func NewGetAttributes500Response() *GetAttributes500Response`

NewGetAttributes500Response instantiates a new GetAttributes500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttributes500ResponseWithDefaults

`func NewGetAttributes500ResponseWithDefaults() *GetAttributes500Response`

NewGetAttributes500ResponseWithDefaults instantiates a new GetAttributes500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetAttributes500Response) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAttributes500Response) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAttributes500Response) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *GetAttributes500Response) HasError() bool`

HasError returns a boolean if a field has been set.


