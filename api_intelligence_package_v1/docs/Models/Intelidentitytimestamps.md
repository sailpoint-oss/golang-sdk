---
id: v1-intelidentitytimestamps
title: Intelidentitytimestamps
pagination_label: Intelidentitytimestamps
sidebar_label: Intelidentitytimestamps
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentitytimestamps', 'V1Intelidentitytimestamps'] 
slug: /tools/sdk/go/intelligencepackagev1/models/intelidentitytimestamps
tags: ['SDK', 'Software Development Kit', 'Intelidentitytimestamps', 'V1Intelidentitytimestamps']
---

# Intelidentitytimestamps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **SailPointTime** | Timestamp when the identity record was first created in Identity Security Cloud. | 
**ModifiedAt** | **SailPointTime** | Timestamp when the identity record was last modified in Identity Security Cloud. | 

## Methods

### NewIntelidentitytimestamps

`func NewIntelidentitytimestamps(createdAt SailPointTime, modifiedAt SailPointTime, ) *Intelidentitytimestamps`

NewIntelidentitytimestamps instantiates a new Intelidentitytimestamps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentitytimestampsWithDefaults

`func NewIntelidentitytimestampsWithDefaults() *Intelidentitytimestamps`

NewIntelidentitytimestampsWithDefaults instantiates a new Intelidentitytimestamps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Intelidentitytimestamps) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Intelidentitytimestamps) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Intelidentitytimestamps) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Intelidentitytimestamps) GetModifiedAt() SailPointTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Intelidentitytimestamps) GetModifiedAtOk() (*SailPointTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Intelidentitytimestamps) SetModifiedAt(v SailPointTime)`

SetModifiedAt sets ModifiedAt field to given value.



