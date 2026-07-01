---
id: v1-intelaccountsslice
title: Intelaccountsslice
pagination_label: Intelaccountsslice
sidebar_label: Intelaccountsslice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelaccountsslice', 'V1Intelaccountsslice'] 
slug: /tools/sdk/go/intelligence/models/intelaccountsslice
tags: ['SDK', 'Software Development Kit', 'Intelaccountsslice', 'V1Intelaccountsslice']
---

# Intelaccountsslice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Intelaccessaccountwire**](intelaccessaccountwire) | First page of accounts for the identity. | 
**Next** | Pointer to **string** | Absolute URL to the next accounts page; present only when more results exist. | [optional] 

## Methods

### NewIntelaccountsslice

`func NewIntelaccountsslice(items []Intelaccessaccountwire, ) *Intelaccountsslice`

NewIntelaccountsslice instantiates a new Intelaccountsslice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelaccountssliceWithDefaults

`func NewIntelaccountssliceWithDefaults() *Intelaccountsslice`

NewIntelaccountssliceWithDefaults instantiates a new Intelaccountsslice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Intelaccountsslice) GetItems() []Intelaccessaccountwire`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Intelaccountsslice) GetItemsOk() (*[]Intelaccessaccountwire, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Intelaccountsslice) SetItems(v []Intelaccessaccountwire)`

SetItems sets Items field to given value.


### GetNext

`func (o *Intelaccountsslice) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Intelaccountsslice) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Intelaccountsslice) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Intelaccountsslice) HasNext() bool`

HasNext returns a boolean if a field has been set.


