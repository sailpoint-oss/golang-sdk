---
id: v1-intel-rare-access-slice
title: IntelRareAccessSlice
pagination_label: IntelRareAccessSlice
sidebar_label: IntelRareAccessSlice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IntelRareAccessSlice', 'V1IntelRareAccessSlice'] 
slug: /tools/sdk/go/intelligence/models/intel-rare-access-slice
tags: ['SDK', 'Software Development Kit', 'IntelRareAccessSlice', 'V1IntelRareAccessSlice']
---

# IntelRareAccessSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]IntelOutlierAccessItem**](intel-outlier-access-item) | First page of rare access items for the identity. | 
**Next** | Pointer to **string** | Absolute URL to the next rareAccess page; present only when more results exist. | [optional] 

## Methods

### NewIntelRareAccessSlice

`func NewIntelRareAccessSlice(items []IntelOutlierAccessItem, ) *IntelRareAccessSlice`

NewIntelRareAccessSlice instantiates a new IntelRareAccessSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelRareAccessSliceWithDefaults

`func NewIntelRareAccessSliceWithDefaults() *IntelRareAccessSlice`

NewIntelRareAccessSliceWithDefaults instantiates a new IntelRareAccessSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *IntelRareAccessSlice) GetItems() []IntelOutlierAccessItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IntelRareAccessSlice) GetItemsOk() (*[]IntelOutlierAccessItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IntelRareAccessSlice) SetItems(v []IntelOutlierAccessItem)`

SetItems sets Items field to given value.


### GetNext

`func (o *IntelRareAccessSlice) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *IntelRareAccessSlice) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *IntelRareAccessSlice) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *IntelRareAccessSlice) HasNext() bool`

HasNext returns a boolean if a field has been set.


