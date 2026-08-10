---
id: v1-intelmachineaccountsslice
title: Intelmachineaccountsslice
pagination_label: Intelmachineaccountsslice
sidebar_label: Intelmachineaccountsslice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelmachineaccountsslice', 'V1Intelmachineaccountsslice'] 
slug: /tools/sdk/go/intelligence/models/intelmachineaccountsslice
tags: ['SDK', 'Software Development Kit', 'Intelmachineaccountsslice', 'V1Intelmachineaccountsslice']
---

# Intelmachineaccountsslice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Intelmachineaccountwire**](intelmachineaccountwire) | Machine account rows correlated to the non-human identity. | 

## Methods

### NewIntelmachineaccountsslice

`func NewIntelmachineaccountsslice(items []Intelmachineaccountwire, ) *Intelmachineaccountsslice`

NewIntelmachineaccountsslice instantiates a new Intelmachineaccountsslice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelmachineaccountssliceWithDefaults

`func NewIntelmachineaccountssliceWithDefaults() *Intelmachineaccountsslice`

NewIntelmachineaccountssliceWithDefaults instantiates a new Intelmachineaccountsslice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Intelmachineaccountsslice) GetItems() []Intelmachineaccountwire`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Intelmachineaccountsslice) GetItemsOk() (*[]Intelmachineaccountwire, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Intelmachineaccountsslice) SetItems(v []Intelmachineaccountwire)`

SetItems sets Items field to given value.



