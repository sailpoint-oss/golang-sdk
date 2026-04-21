---
id: nermv2025-delegations-get500-response
title: DelegationsGet500Response
pagination_label: DelegationsGet500Response
sidebar_label: DelegationsGet500Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DelegationsGet500Response', 'NERMV2025DelegationsGet500Response'] 
slug: /tools/sdk/go/nermv2025/models/delegations-get500-response
tags: ['SDK', 'Software Development Kit', 'DelegationsGet500Response', 'NERMV2025DelegationsGet500Response']
---

# DelegationsGet500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** | A message describing the error | [optional] 

## Methods

### NewDelegationsGet500Response

`func NewDelegationsGet500Response() *DelegationsGet500Response`

NewDelegationsGet500Response instantiates a new DelegationsGet500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationsGet500ResponseWithDefaults

`func NewDelegationsGet500ResponseWithDefaults() *DelegationsGet500Response`

NewDelegationsGet500ResponseWithDefaults instantiates a new DelegationsGet500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DelegationsGet500Response) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DelegationsGet500Response) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DelegationsGet500Response) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *DelegationsGet500Response) HasError() bool`

HasError returns a boolean if a field has been set.


