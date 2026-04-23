---
id: v2026-jit-activation-extend-request
title: JitActivationExtendRequest
pagination_label: JitActivationExtendRequest
sidebar_label: JitActivationExtendRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'JitActivationExtendRequest', 'V2026JitActivationExtendRequest'] 
slug: /tools/sdk/go/v2026/models/jit-activation-extend-request
tags: ['SDK', 'Software Development Kit', 'JitActivationExtendRequest', 'V2026JitActivationExtendRequest']
---

# JitActivationExtendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Entitlement connection identifier for the activation to extend. | 
**ActivationPeriodExtensionMins** | **int32** | Number of minutes to extend the activation period. | 

## Methods

### NewJitActivationExtendRequest

`func NewJitActivationExtendRequest(connectionId string, activationPeriodExtensionMins int32, ) *JitActivationExtendRequest`

NewJitActivationExtendRequest instantiates a new JitActivationExtendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitActivationExtendRequestWithDefaults

`func NewJitActivationExtendRequestWithDefaults() *JitActivationExtendRequest`

NewJitActivationExtendRequestWithDefaults instantiates a new JitActivationExtendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *JitActivationExtendRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *JitActivationExtendRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *JitActivationExtendRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetActivationPeriodExtensionMins

`func (o *JitActivationExtendRequest) GetActivationPeriodExtensionMins() int32`

GetActivationPeriodExtensionMins returns the ActivationPeriodExtensionMins field if non-nil, zero value otherwise.

### GetActivationPeriodExtensionMinsOk

`func (o *JitActivationExtendRequest) GetActivationPeriodExtensionMinsOk() (*int32, bool)`

GetActivationPeriodExtensionMinsOk returns a tuple with the ActivationPeriodExtensionMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPeriodExtensionMins

`func (o *JitActivationExtendRequest) SetActivationPeriodExtensionMins(v int32)`

SetActivationPeriodExtensionMins sets ActivationPeriodExtensionMins field to given value.



