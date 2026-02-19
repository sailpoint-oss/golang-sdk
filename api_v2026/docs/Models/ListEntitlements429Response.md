---
id: v2026-list-entitlements429-response
title: ListEntitlements429Response
pagination_label: ListEntitlements429Response
sidebar_label: ListEntitlements429Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ListEntitlements429Response', 'V2026ListEntitlements429Response'] 
slug: /tools/sdk/go/v2026/models/list-entitlements429-response
tags: ['SDK', 'Software Development Kit', 'ListEntitlements429Response', 'V2026ListEntitlements429Response']
---

# ListEntitlements429Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **map[string]interface{}** | A message describing the error | [optional] 

## Methods

### NewListEntitlements429Response

`func NewListEntitlements429Response() *ListEntitlements429Response`

NewListEntitlements429Response instantiates a new ListEntitlements429Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitlements429ResponseWithDefaults

`func NewListEntitlements429ResponseWithDefaults() *ListEntitlements429Response`

NewListEntitlements429ResponseWithDefaults instantiates a new ListEntitlements429Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListEntitlements429Response) GetMessage() map[string]interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListEntitlements429Response) GetMessageOk() (*map[string]interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListEntitlements429Response) SetMessage(v map[string]interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListEntitlements429Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


