---
id: v1-intel-accounts-slice
title: IntelAccountsSlice
pagination_label: IntelAccountsSlice
sidebar_label: IntelAccountsSlice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IntelAccountsSlice', 'V1IntelAccountsSlice'] 
slug: /tools/sdk/go/intelligence/models/intel-accounts-slice
tags: ['SDK', 'Software Development Kit', 'IntelAccountsSlice', 'V1IntelAccountsSlice']
---

# IntelAccountsSlice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]IntelAccessAccountWire**](intel-access-account-wire) | First page of accounts for the identity. | 
**Next** | Pointer to **string** | Absolute URL to the next accounts page; present only when more results exist. | [optional] 

## Methods

### NewIntelAccountsSlice

`func NewIntelAccountsSlice(items []IntelAccessAccountWire, ) *IntelAccountsSlice`

NewIntelAccountsSlice instantiates a new IntelAccountsSlice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelAccountsSliceWithDefaults

`func NewIntelAccountsSliceWithDefaults() *IntelAccountsSlice`

NewIntelAccountsSliceWithDefaults instantiates a new IntelAccountsSlice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *IntelAccountsSlice) GetItems() []IntelAccessAccountWire`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IntelAccountsSlice) GetItemsOk() (*[]IntelAccessAccountWire, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IntelAccountsSlice) SetItems(v []IntelAccessAccountWire)`

SetItems sets Items field to given value.


### GetNext

`func (o *IntelAccountsSlice) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *IntelAccountsSlice) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *IntelAccountsSlice) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *IntelAccountsSlice) HasNext() bool`

HasNext returns a boolean if a field has been set.


