---
id: v1-intelaccesshistoryaccessitemsslice
title: Intelaccesshistoryaccessitemsslice
pagination_label: Intelaccesshistoryaccessitemsslice
sidebar_label: Intelaccesshistoryaccessitemsslice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelaccesshistoryaccessitemsslice', 'V1Intelaccesshistoryaccessitemsslice'] 
slug: /tools/sdk/go/intelligence/models/intelaccesshistoryaccessitemsslice
tags: ['SDK', 'Software Development Kit', 'Intelaccesshistoryaccessitemsslice', 'V1Intelaccesshistoryaccessitemsslice']
---

# Intelaccesshistoryaccessitemsslice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]Intelaccessitemhistoryevent** | First page of access-item history events for the identity. | 
**Next** | Pointer to **string** | Absolute URL to the next access-items page; present only when more results exist. | [optional] 

## Methods

### NewIntelaccesshistoryaccessitemsslice

`func NewIntelaccesshistoryaccessitemsslice(items []Intelaccessitemhistoryevent, ) *Intelaccesshistoryaccessitemsslice`

NewIntelaccesshistoryaccessitemsslice instantiates a new Intelaccesshistoryaccessitemsslice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelaccesshistoryaccessitemssliceWithDefaults

`func NewIntelaccesshistoryaccessitemssliceWithDefaults() *Intelaccesshistoryaccessitemsslice`

NewIntelaccesshistoryaccessitemssliceWithDefaults instantiates a new Intelaccesshistoryaccessitemsslice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Intelaccesshistoryaccessitemsslice) GetItems() []Intelaccessitemhistoryevent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Intelaccesshistoryaccessitemsslice) GetItemsOk() (*[]Intelaccessitemhistoryevent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Intelaccesshistoryaccessitemsslice) SetItems(v []Intelaccessitemhistoryevent)`

SetItems sets Items field to given value.


### GetNext

`func (o *Intelaccesshistoryaccessitemsslice) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Intelaccesshistoryaccessitemsslice) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Intelaccesshistoryaccessitemsslice) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Intelaccesshistoryaccessitemsslice) HasNext() bool`

HasNext returns a boolean if a field has been set.


