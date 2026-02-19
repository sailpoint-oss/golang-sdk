---
id: v2026-list-entitlements401-response
title: ListEntitlements401Response
pagination_label: ListEntitlements401Response
sidebar_label: ListEntitlements401Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ListEntitlements401Response', 'V2026ListEntitlements401Response'] 
slug: /tools/sdk/go/v2026/models/list-entitlements401-response
tags: ['SDK', 'Software Development Kit', 'ListEntitlements401Response', 'V2026ListEntitlements401Response']
---

# ListEntitlements401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **map[string]interface{}** | A message describing the error | [optional] 

## Methods

### NewListEntitlements401Response

`func NewListEntitlements401Response() *ListEntitlements401Response`

NewListEntitlements401Response instantiates a new ListEntitlements401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitlements401ResponseWithDefaults

`func NewListEntitlements401ResponseWithDefaults() *ListEntitlements401Response`

NewListEntitlements401ResponseWithDefaults instantiates a new ListEntitlements401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ListEntitlements401Response) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListEntitlements401Response) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListEntitlements401Response) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ListEntitlements401Response) HasError() bool`

HasError returns a boolean if a field has been set.


