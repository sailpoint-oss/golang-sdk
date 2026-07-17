---
id: v1-intel-access-history-access-items-slice
title: IntelAccessHistoryAccessItemsSlice
pagination_label: IntelAccessHistoryAccessItemsSlice
sidebar_label: IntelAccessHistoryAccessItemsSlice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IntelAccessHistoryAccessItemsSlice', 'V1IntelAccessHistoryAccessItemsSlice'] 
slug: /tools/sdk/go/intelligence/models/intel-access-history-access-items-slice
tags: ['SDK', 'Software Development Kit', 'IntelAccessHistoryAccessItemsSlice', 'V1IntelAccessHistoryAccessItemsSlice']
---

# IntelAccessHistoryAccessItemsSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]IntelAccessItemHistoryEvent** | First page of access-item history events for the identity. | 
**Next** | Pointer to **string** | Absolute URL to the next access-items page; present only when more results exist. | [optional] 

## Methods

### NewIntelAccessHistoryAccessItemsSlice

`func NewIntelAccessHistoryAccessItemsSlice(items []IntelAccessItemHistoryEvent, ) *IntelAccessHistoryAccessItemsSlice`

NewIntelAccessHistoryAccessItemsSlice instantiates a new IntelAccessHistoryAccessItemsSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelAccessHistoryAccessItemsSliceWithDefaults

`func NewIntelAccessHistoryAccessItemsSliceWithDefaults() *IntelAccessHistoryAccessItemsSlice`

NewIntelAccessHistoryAccessItemsSliceWithDefaults instantiates a new IntelAccessHistoryAccessItemsSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *IntelAccessHistoryAccessItemsSlice) GetItems() []IntelAccessItemHistoryEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IntelAccessHistoryAccessItemsSlice) GetItemsOk() (*[]IntelAccessItemHistoryEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IntelAccessHistoryAccessItemsSlice) SetItems(v []IntelAccessItemHistoryEvent)`

SetItems sets Items field to given value.


### GetNext

`func (o *IntelAccessHistoryAccessItemsSlice) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *IntelAccessHistoryAccessItemsSlice) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *IntelAccessHistoryAccessItemsSlice) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *IntelAccessHistoryAccessItemsSlice) HasNext() bool`

HasNext returns a boolean if a field has been set.


