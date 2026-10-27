---
id: v1-source-subtype-override-request
title: SourceSubtypeOverrideRequest
pagination_label: SourceSubtypeOverrideRequest
sidebar_label: SourceSubtypeOverrideRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SourceSubtypeOverrideRequest', 'V1SourceSubtypeOverrideRequest'] 
slug: /tools/sdk/go/machineaccountsubtypes/models/source-subtype-override-request
tags: ['SDK', 'Software Development Kit', 'SourceSubtypeOverrideRequest', 'V1SourceSubtypeOverrideRequest']
---

# SourceSubtypeOverrideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | ID of the source to create the overriding source subtype on. | 

## Methods

### NewSourceSubtypeOverrideRequest

`func NewSourceSubtypeOverrideRequest(sourceId string, ) *SourceSubtypeOverrideRequest`

NewSourceSubtypeOverrideRequest instantiates a new SourceSubtypeOverrideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceSubtypeOverrideRequestWithDefaults

`func NewSourceSubtypeOverrideRequestWithDefaults() *SourceSubtypeOverrideRequest`

NewSourceSubtypeOverrideRequestWithDefaults instantiates a new SourceSubtypeOverrideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *SourceSubtypeOverrideRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourceSubtypeOverrideRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourceSubtypeOverrideRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.



