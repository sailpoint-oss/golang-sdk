---
id: v2026-start-application-discovery403-response-one-of
title: StartApplicationDiscovery403ResponseOneOf
pagination_label: StartApplicationDiscovery403ResponseOneOf
sidebar_label: StartApplicationDiscovery403ResponseOneOf
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'StartApplicationDiscovery403ResponseOneOf', 'V2026StartApplicationDiscovery403ResponseOneOf'] 
slug: /tools/sdk/go/v2026/models/start-application-discovery403-response-one-of
tags: ['SDK', 'Software Development Kit', 'StartApplicationDiscovery403ResponseOneOf', 'V2026StartApplicationDiscovery403ResponseOneOf']
---

# StartApplicationDiscovery403ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Error message when quota is exceeded | 

## Methods

### NewStartApplicationDiscovery403ResponseOneOf

`func NewStartApplicationDiscovery403ResponseOneOf(error_ string, ) *StartApplicationDiscovery403ResponseOneOf`

NewStartApplicationDiscovery403ResponseOneOf instantiates a new StartApplicationDiscovery403ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartApplicationDiscovery403ResponseOneOfWithDefaults

`func NewStartApplicationDiscovery403ResponseOneOfWithDefaults() *StartApplicationDiscovery403ResponseOneOf`

NewStartApplicationDiscovery403ResponseOneOfWithDefaults instantiates a new StartApplicationDiscovery403ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StartApplicationDiscovery403ResponseOneOf) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StartApplicationDiscovery403ResponseOneOf) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StartApplicationDiscovery403ResponseOneOf) SetError(v string)`

SetError sets Error field to given value.



