---
id: v1-intelcertificationhistoryevent
title: Intelcertificationhistoryevent
pagination_label: Intelcertificationhistoryevent
sidebar_label: Intelcertificationhistoryevent
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelcertificationhistoryevent', 'V1Intelcertificationhistoryevent'] 
slug: /tools/sdk/go/intelligence/models/intelcertificationhistoryevent
tags: ['SDK', 'Software Development Kit', 'Intelcertificationhistoryevent', 'V1Intelcertificationhistoryevent']
---

# Intelcertificationhistoryevent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | Type of certification history event. | 
**DateTime** | Pointer to **SailPointTime** | Event timestamp. | [optional] 
**CertificationId** | Pointer to **string** | Identifier of the certification. | [optional] 
**CertificationName** | Pointer to **string** | Display name of the certification. | [optional] 
**SignedDate** | Pointer to **SailPointTime** | Timestamp when the certification was signed. | [optional] 

## Methods

### NewIntelcertificationhistoryevent

`func NewIntelcertificationhistoryevent(eventType string, ) *Intelcertificationhistoryevent`

NewIntelcertificationhistoryevent instantiates a new Intelcertificationhistoryevent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelcertificationhistoryeventWithDefaults

`func NewIntelcertificationhistoryeventWithDefaults() *Intelcertificationhistoryevent`

NewIntelcertificationhistoryeventWithDefaults instantiates a new Intelcertificationhistoryevent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *Intelcertificationhistoryevent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Intelcertificationhistoryevent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Intelcertificationhistoryevent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetDateTime

`func (o *Intelcertificationhistoryevent) GetDateTime() SailPointTime`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *Intelcertificationhistoryevent) GetDateTimeOk() (*SailPointTime, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *Intelcertificationhistoryevent) SetDateTime(v SailPointTime)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *Intelcertificationhistoryevent) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetCertificationId

`func (o *Intelcertificationhistoryevent) GetCertificationId() string`

GetCertificationId returns the CertificationId field if non-nil, zero value otherwise.

### GetCertificationIdOk

`func (o *Intelcertificationhistoryevent) GetCertificationIdOk() (*string, bool)`

GetCertificationIdOk returns a tuple with the CertificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationId

`func (o *Intelcertificationhistoryevent) SetCertificationId(v string)`

SetCertificationId sets CertificationId field to given value.

### HasCertificationId

`func (o *Intelcertificationhistoryevent) HasCertificationId() bool`

HasCertificationId returns a boolean if a field has been set.

### GetCertificationName

`func (o *Intelcertificationhistoryevent) GetCertificationName() string`

GetCertificationName returns the CertificationName field if non-nil, zero value otherwise.

### GetCertificationNameOk

`func (o *Intelcertificationhistoryevent) GetCertificationNameOk() (*string, bool)`

GetCertificationNameOk returns a tuple with the CertificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationName

`func (o *Intelcertificationhistoryevent) SetCertificationName(v string)`

SetCertificationName sets CertificationName field to given value.

### HasCertificationName

`func (o *Intelcertificationhistoryevent) HasCertificationName() bool`

HasCertificationName returns a boolean if a field has been set.

### GetSignedDate

`func (o *Intelcertificationhistoryevent) GetSignedDate() SailPointTime`

GetSignedDate returns the SignedDate field if non-nil, zero value otherwise.

### GetSignedDateOk

`func (o *Intelcertificationhistoryevent) GetSignedDateOk() (*SailPointTime, bool)`

GetSignedDateOk returns a tuple with the SignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedDate

`func (o *Intelcertificationhistoryevent) SetSignedDate(v SailPointTime)`

SetSignedDate sets SignedDate field to given value.

### HasSignedDate

`func (o *Intelcertificationhistoryevent) HasSignedDate() bool`

HasSignedDate returns a boolean if a field has been set.


