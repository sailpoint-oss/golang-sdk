---
id: v1-inteloutliersslice
title: Inteloutliersslice
pagination_label: Inteloutliersslice
sidebar_label: Inteloutliersslice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Inteloutliersslice', 'V1Inteloutliersslice'] 
slug: /tools/sdk/go/intelligence/models/inteloutliersslice
tags: ['SDK', 'Software Development Kit', 'Inteloutliersslice', 'V1Inteloutliersslice']
---

# Inteloutliersslice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RareAccess** | [**Intelrareaccessslice**](intelrareaccessslice) | First page of rare access items for the identity. | 

## Methods

### NewInteloutliersslice

`func NewInteloutliersslice(rareAccess Intelrareaccessslice, ) *Inteloutliersslice`

NewInteloutliersslice instantiates a new Inteloutliersslice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteloutlierssliceWithDefaults

`func NewInteloutlierssliceWithDefaults() *Inteloutliersslice`

NewInteloutlierssliceWithDefaults instantiates a new Inteloutliersslice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRareAccess

`func (o *Inteloutliersslice) GetRareAccess() Intelrareaccessslice`

GetRareAccess returns the RareAccess field if non-nil, zero value otherwise.

### GetRareAccessOk

`func (o *Inteloutliersslice) GetRareAccessOk() (*Intelrareaccessslice, bool)`

GetRareAccessOk returns a tuple with the RareAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRareAccess

`func (o *Inteloutliersslice) SetRareAccess(v Intelrareaccessslice)`

SetRareAccess sets RareAccess field to given value.



