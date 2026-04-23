---
id: v2026-jit-activation-deactivate-request
title: JitActivationDeactivateRequest
pagination_label: JitActivationDeactivateRequest
sidebar_label: JitActivationDeactivateRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'JitActivationDeactivateRequest', 'V2026JitActivationDeactivateRequest'] 
slug: /tools/sdk/go/v2026/models/jit-activation-deactivate-request
tags: ['SDK', 'Software Development Kit', 'JitActivationDeactivateRequest', 'V2026JitActivationDeactivateRequest']
---

# JitActivationDeactivateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Entitlement connection identifier for the activation to deactivate. | 

## Methods

### NewJitActivationDeactivateRequest

`func NewJitActivationDeactivateRequest(connectionId string, ) *JitActivationDeactivateRequest`

NewJitActivationDeactivateRequest instantiates a new JitActivationDeactivateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitActivationDeactivateRequestWithDefaults

`func NewJitActivationDeactivateRequestWithDefaults() *JitActivationDeactivateRequest`

NewJitActivationDeactivateRequestWithDefaults instantiates a new JitActivationDeactivateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *JitActivationDeactivateRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *JitActivationDeactivateRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *JitActivationDeactivateRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.



