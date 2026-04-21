---
id: nerm-get-attributes400-response
title: GetAttributes400Response
pagination_label: GetAttributes400Response
sidebar_label: GetAttributes400Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetAttributes400Response', 'NERMGetAttributes400Response'] 
slug: /tools/sdk/go/nerm/models/get-attributes400-response
tags: ['SDK', 'Software Development Kit', 'GetAttributes400Response', 'NERMGetAttributes400Response']
---

# GetAttributes400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetAttributes400Response

`func NewGetAttributes400Response() *GetAttributes400Response`

NewGetAttributes400Response instantiates a new GetAttributes400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAttributes400ResponseWithDefaults

`func NewGetAttributes400ResponseWithDefaults() *GetAttributes400Response`

NewGetAttributes400ResponseWithDefaults instantiates a new GetAttributes400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetAttributes400Response) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAttributes400Response) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAttributes400Response) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *GetAttributes400Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *GetAttributes400Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetAttributes400Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetAttributes400Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetAttributes400Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


