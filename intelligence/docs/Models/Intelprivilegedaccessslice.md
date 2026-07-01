---
id: v1-intelprivilegedaccessslice
title: Intelprivilegedaccessslice
pagination_label: Intelprivilegedaccessslice
sidebar_label: Intelprivilegedaccessslice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelprivilegedaccessslice', 'V1Intelprivilegedaccessslice'] 
slug: /tools/sdk/go/intelligence/models/intelprivilegedaccessslice
tags: ['SDK', 'Software Development Kit', 'Intelprivilegedaccessslice', 'V1Intelprivilegedaccessslice']
---

# Intelprivilegedaccessslice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Intelprivilegedaccessitemwire**](intelprivilegedaccessitemwire) | Privileged access items for the identity. | 

## Methods

### NewIntelprivilegedaccessslice

`func NewIntelprivilegedaccessslice(items []Intelprivilegedaccessitemwire, ) *Intelprivilegedaccessslice`

NewIntelprivilegedaccessslice instantiates a new Intelprivilegedaccessslice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelprivilegedaccesssliceWithDefaults

`func NewIntelprivilegedaccesssliceWithDefaults() *Intelprivilegedaccessslice`

NewIntelprivilegedaccesssliceWithDefaults instantiates a new Intelprivilegedaccessslice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Intelprivilegedaccessslice) GetItems() []Intelprivilegedaccessitemwire`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Intelprivilegedaccessslice) GetItemsOk() (*[]Intelprivilegedaccessitemwire, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Intelprivilegedaccessslice) SetItems(v []Intelprivilegedaccessitemwire)`

SetItems sets Items field to given value.



