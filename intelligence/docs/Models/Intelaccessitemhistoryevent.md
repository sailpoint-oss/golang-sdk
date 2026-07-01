---
id: v1-intelaccessitemhistoryevent
title: Intelaccessitemhistoryevent
pagination_label: Intelaccessitemhistoryevent
sidebar_label: Intelaccessitemhistoryevent
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelaccessitemhistoryevent', 'V1Intelaccessitemhistoryevent'] 
slug: /tools/sdk/go/intelligence/models/intelaccessitemhistoryevent
tags: ['SDK', 'Software Development Kit', 'Intelaccessitemhistoryevent', 'V1Intelaccessitemhistoryevent']
---

# Intelaccessitemhistoryevent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | Type of access-item history event. | 
**DateTime** | Pointer to **SailPointTime** | Event timestamp. | [optional] 

## Methods

### NewIntelaccessitemhistoryevent

`func NewIntelaccessitemhistoryevent(eventType string, ) *Intelaccessitemhistoryevent`

NewIntelaccessitemhistoryevent instantiates a new Intelaccessitemhistoryevent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelaccessitemhistoryeventWithDefaults

`func NewIntelaccessitemhistoryeventWithDefaults() *Intelaccessitemhistoryevent`

NewIntelaccessitemhistoryeventWithDefaults instantiates a new Intelaccessitemhistoryevent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *Intelaccessitemhistoryevent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Intelaccessitemhistoryevent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Intelaccessitemhistoryevent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetDateTime

`func (o *Intelaccessitemhistoryevent) GetDateTime() SailPointTime`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *Intelaccessitemhistoryevent) GetDateTimeOk() (*SailPointTime, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *Intelaccessitemhistoryevent) SetDateTime(v SailPointTime)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *Intelaccessitemhistoryevent) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.


