---
id: v1-misbulkrequest
title: Misbulkrequest
pagination_label: Misbulkrequest
sidebar_label: Misbulkrequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Misbulkrequest', 'V1Misbulkrequest'] 
slug: /tools/sdk/go/machineaccountsv1/models/misbulkrequest
tags: ['SDK', 'Software Development Kit', 'Misbulkrequest', 'V1Misbulkrequest']
---

# Misbulkrequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | Machine identity or machine account IDs to include in the bulk operation. | 

## Methods

### NewMisbulkrequest

`func NewMisbulkrequest(ids []string, ) *Misbulkrequest`

NewMisbulkrequest instantiates a new Misbulkrequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMisbulkrequestWithDefaults

`func NewMisbulkrequestWithDefaults() *Misbulkrequest`

NewMisbulkrequestWithDefaults instantiates a new Misbulkrequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *Misbulkrequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *Misbulkrequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *Misbulkrequest) SetIds(v []string)`

SetIds sets Ids field to given value.



