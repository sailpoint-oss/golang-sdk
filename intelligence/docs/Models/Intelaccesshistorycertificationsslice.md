---
id: v1-intelaccesshistorycertificationsslice
title: Intelaccesshistorycertificationsslice
pagination_label: Intelaccesshistorycertificationsslice
sidebar_label: Intelaccesshistorycertificationsslice
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelaccesshistorycertificationsslice', 'V1Intelaccesshistorycertificationsslice'] 
slug: /tools/sdk/go/intelligence/models/intelaccesshistorycertificationsslice
tags: ['SDK', 'Software Development Kit', 'Intelaccesshistorycertificationsslice', 'V1Intelaccesshistorycertificationsslice']
---

# Intelaccesshistorycertificationsslice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]Intelcertificationhistoryevent** | First page of certification history events for the identity. | 
**Next** | Pointer to **string** | Absolute URL to the next certifications page; present only when more results exist. | [optional] 

## Methods

### NewIntelaccesshistorycertificationsslice

`func NewIntelaccesshistorycertificationsslice(items []Intelcertificationhistoryevent, ) *Intelaccesshistorycertificationsslice`

NewIntelaccesshistorycertificationsslice instantiates a new Intelaccesshistorycertificationsslice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelaccesshistorycertificationssliceWithDefaults

`func NewIntelaccesshistorycertificationssliceWithDefaults() *Intelaccesshistorycertificationsslice`

NewIntelaccesshistorycertificationssliceWithDefaults instantiates a new Intelaccesshistorycertificationsslice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Intelaccesshistorycertificationsslice) GetItems() []Intelcertificationhistoryevent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Intelaccesshistorycertificationsslice) GetItemsOk() (*[]Intelcertificationhistoryevent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Intelaccesshistorycertificationsslice) SetItems(v []Intelcertificationhistoryevent)`

SetItems sets Items field to given value.


### GetNext

`func (o *Intelaccesshistorycertificationsslice) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Intelaccesshistorycertificationsslice) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Intelaccesshistorycertificationsslice) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Intelaccesshistorycertificationsslice) HasNext() bool`

HasNext returns a boolean if a field has been set.


