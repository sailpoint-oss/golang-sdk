---
id: v1-intelrareaccessslice
title: Intelrareaccessslice
pagination_label: Intelrareaccessslice
sidebar_label: Intelrareaccessslice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelrareaccessslice', 'V1Intelrareaccessslice'] 
slug: /tools/sdk/go/intelligence/models/intelrareaccessslice
tags: ['SDK', 'Software Development Kit', 'Intelrareaccessslice', 'V1Intelrareaccessslice']
---

# Intelrareaccessslice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Inteloutlieraccessitem**](inteloutlieraccessitem) | First page of rare access items for the identity. | 
**Next** | Pointer to **string** | Absolute URL to the next rareAccess page; present only when more results exist. | [optional] 

## Methods

### NewIntelrareaccessslice

`func NewIntelrareaccessslice(items []Inteloutlieraccessitem, ) *Intelrareaccessslice`

NewIntelrareaccessslice instantiates a new Intelrareaccessslice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelrareaccesssliceWithDefaults

`func NewIntelrareaccesssliceWithDefaults() *Intelrareaccessslice`

NewIntelrareaccesssliceWithDefaults instantiates a new Intelrareaccessslice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Intelrareaccessslice) GetItems() []Inteloutlieraccessitem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Intelrareaccessslice) GetItemsOk() (*[]Inteloutlieraccessitem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Intelrareaccessslice) SetItems(v []Inteloutlieraccessitem)`

SetItems sets Items field to given value.


### GetNext

`func (o *Intelrareaccessslice) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Intelrareaccessslice) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Intelrareaccessslice) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Intelrareaccessslice) HasNext() bool`

HasNext returns a boolean if a field has been set.


