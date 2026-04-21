---
id: nermv2025-delegations-post404-response
title: DelegationsPost404Response
pagination_label: DelegationsPost404Response
sidebar_label: DelegationsPost404Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DelegationsPost404Response', 'NERMV2025DelegationsPost404Response'] 
slug: /tools/sdk/go/nermv2025/models/delegations-post404-response
tags: ['SDK', 'Software Development Kit', 'DelegationsPost404Response', 'NERMV2025DelegationsPost404Response']
---

# DelegationsPost404Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** | The requested record, either ID or UID, was not found | [optional] 

## Methods

### NewDelegationsPost404Response

`func NewDelegationsPost404Response() *DelegationsPost404Response`

NewDelegationsPost404Response instantiates a new DelegationsPost404Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationsPost404ResponseWithDefaults

`func NewDelegationsPost404ResponseWithDefaults() *DelegationsPost404Response`

NewDelegationsPost404ResponseWithDefaults instantiates a new DelegationsPost404Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DelegationsPost404Response) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DelegationsPost404Response) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DelegationsPost404Response) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *DelegationsPost404Response) HasError() bool`

HasError returns a boolean if a field has been set.


