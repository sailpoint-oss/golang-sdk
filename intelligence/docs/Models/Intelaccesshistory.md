---
id: v1-intelaccesshistory
title: Intelaccesshistory
pagination_label: Intelaccesshistory
sidebar_label: Intelaccesshistory
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelaccesshistory', 'V1Intelaccesshistory'] 
slug: /tools/sdk/go/intelligence/models/intelaccesshistory
tags: ['SDK', 'Software Development Kit', 'Intelaccesshistory', 'V1Intelaccesshistory']
---

# Intelaccesshistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessItems** | [**Intelaccesshistoryaccessitemsslice**](intelaccesshistoryaccessitemsslice) | First page of access-item history events for the identity. | 
**Certifications** | [**Intelaccesshistorycertificationsslice**](intelaccesshistorycertificationsslice) | First page of certification history events for the identity. | 

## Methods

### NewIntelaccesshistory

`func NewIntelaccesshistory(accessItems Intelaccesshistoryaccessitemsslice, certifications Intelaccesshistorycertificationsslice, ) *Intelaccesshistory`

NewIntelaccesshistory instantiates a new Intelaccesshistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelaccesshistoryWithDefaults

`func NewIntelaccesshistoryWithDefaults() *Intelaccesshistory`

NewIntelaccesshistoryWithDefaults instantiates a new Intelaccesshistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessItems

`func (o *Intelaccesshistory) GetAccessItems() Intelaccesshistoryaccessitemsslice`

GetAccessItems returns the AccessItems field if non-nil, zero value otherwise.

### GetAccessItemsOk

`func (o *Intelaccesshistory) GetAccessItemsOk() (*Intelaccesshistoryaccessitemsslice, bool)`

GetAccessItemsOk returns a tuple with the AccessItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessItems

`func (o *Intelaccesshistory) SetAccessItems(v Intelaccesshistoryaccessitemsslice)`

SetAccessItems sets AccessItems field to given value.


### GetCertifications

`func (o *Intelaccesshistory) GetCertifications() Intelaccesshistorycertificationsslice`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *Intelaccesshistory) GetCertificationsOk() (*Intelaccesshistorycertificationsslice, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *Intelaccesshistory) SetCertifications(v Intelaccesshistorycertificationsslice)`

SetCertifications sets Certifications field to given value.



