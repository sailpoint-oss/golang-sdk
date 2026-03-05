---
id: v2026-get-access-request-config429-response
title: GetAccessRequestConfig429Response
pagination_label: GetAccessRequestConfig429Response
sidebar_label: GetAccessRequestConfig429Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetAccessRequestConfig429Response', 'V2026GetAccessRequestConfig429Response'] 
slug: /tools/sdk/go/v2026/models/get-access-request-config429-response
tags: ['SDK', 'Software Development Kit', 'GetAccessRequestConfig429Response', 'V2026GetAccessRequestConfig429Response']
---

# GetAccessRequestConfig429Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **map[string]interface{}** | A message describing the error | [optional] 

## Methods

### NewGetAccessRequestConfig429Response

`func NewGetAccessRequestConfig429Response() *GetAccessRequestConfig429Response`

NewGetAccessRequestConfig429Response instantiates a new GetAccessRequestConfig429Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessRequestConfig429ResponseWithDefaults

`func NewGetAccessRequestConfig429ResponseWithDefaults() *GetAccessRequestConfig429Response`

NewGetAccessRequestConfig429ResponseWithDefaults instantiates a new GetAccessRequestConfig429Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetAccessRequestConfig429Response) GetMessage() map[string]interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAccessRequestConfig429Response) GetMessageOk() (*map[string]interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAccessRequestConfig429Response) SetMessage(v map[string]interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetAccessRequestConfig429Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


