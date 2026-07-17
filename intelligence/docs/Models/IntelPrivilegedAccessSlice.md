---
id: v1-intel-privileged-access-slice
title: IntelPrivilegedAccessSlice
pagination_label: IntelPrivilegedAccessSlice
sidebar_label: IntelPrivilegedAccessSlice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IntelPrivilegedAccessSlice', 'V1IntelPrivilegedAccessSlice'] 
slug: /tools/sdk/go/intelligence/models/intel-privileged-access-slice
tags: ['SDK', 'Software Development Kit', 'IntelPrivilegedAccessSlice', 'V1IntelPrivilegedAccessSlice']
---

# IntelPrivilegedAccessSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]IntelPrivilegedAccessItemWire**](intel-privileged-access-item-wire) | Privileged access items for the identity. | 

## Methods

### NewIntelPrivilegedAccessSlice

`func NewIntelPrivilegedAccessSlice(items []IntelPrivilegedAccessItemWire, ) *IntelPrivilegedAccessSlice`

NewIntelPrivilegedAccessSlice instantiates a new IntelPrivilegedAccessSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelPrivilegedAccessSliceWithDefaults

`func NewIntelPrivilegedAccessSliceWithDefaults() *IntelPrivilegedAccessSlice`

NewIntelPrivilegedAccessSliceWithDefaults instantiates a new IntelPrivilegedAccessSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *IntelPrivilegedAccessSlice) GetItems() []IntelPrivilegedAccessItemWire`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IntelPrivilegedAccessSlice) GetItemsOk() (*[]IntelPrivilegedAccessItemWire, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IntelPrivilegedAccessSlice) SetItems(v []IntelPrivilegedAccessItemWire)`

SetItems sets Items field to given value.



