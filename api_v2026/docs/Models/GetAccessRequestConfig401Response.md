---
id: v2026-get-access-request-config401-response
title: GetAccessRequestConfig401Response
pagination_label: GetAccessRequestConfig401Response
sidebar_label: GetAccessRequestConfig401Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetAccessRequestConfig401Response', 'V2026GetAccessRequestConfig401Response'] 
slug: /tools/sdk/go/v2026/models/get-access-request-config401-response
tags: ['SDK', 'Software Development Kit', 'GetAccessRequestConfig401Response', 'V2026GetAccessRequestConfig401Response']
---

# GetAccessRequestConfig401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** | A message describing the error | [optional] 

## Methods

### NewGetAccessRequestConfig401Response

`func NewGetAccessRequestConfig401Response() *GetAccessRequestConfig401Response`

NewGetAccessRequestConfig401Response instantiates a new GetAccessRequestConfig401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessRequestConfig401ResponseWithDefaults

`func NewGetAccessRequestConfig401ResponseWithDefaults() *GetAccessRequestConfig401Response`

NewGetAccessRequestConfig401ResponseWithDefaults instantiates a new GetAccessRequestConfig401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetAccessRequestConfig401Response) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAccessRequestConfig401Response) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAccessRequestConfig401Response) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *GetAccessRequestConfig401Response) HasError() bool`

HasError returns a boolean if a field has been set.


