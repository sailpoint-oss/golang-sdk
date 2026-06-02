---
id: v1-intelidentityaccesshistorybody
title: Intelidentityaccesshistorybody
pagination_label: Intelidentityaccesshistorybody
sidebar_label: Intelidentityaccesshistorybody
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentityaccesshistorybody', 'V1Intelidentityaccesshistorybody'] 
slug: /tools/sdk/go/intelligencepackagev1/models/intelidentityaccesshistorybody
tags: ['SDK', 'Software Development Kit', 'Intelidentityaccesshistorybody', 'V1Intelidentityaccesshistorybody']
---

# Intelidentityaccesshistorybody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]map[string]interface{}** | Each event is relayed from identity-history. Schema varies by event type; consumers should treat unknown fields as opaque using additionalProperties.  | 

## Methods

### NewIntelidentityaccesshistorybody

`func NewIntelidentityaccesshistorybody(events []map[string]interface{}, ) *Intelidentityaccesshistorybody`

NewIntelidentityaccesshistorybody instantiates a new Intelidentityaccesshistorybody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentityaccesshistorybodyWithDefaults

`func NewIntelidentityaccesshistorybodyWithDefaults() *Intelidentityaccesshistorybody`

NewIntelidentityaccesshistorybodyWithDefaults instantiates a new Intelidentityaccesshistorybody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *Intelidentityaccesshistorybody) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Intelidentityaccesshistorybody) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Intelidentityaccesshistorybody) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.



