---
id: v2026-role-propagation-config-input
title: RolePropagationConfigInput
pagination_label: RolePropagationConfigInput
sidebar_label: RolePropagationConfigInput
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RolePropagationConfigInput', 'V2026RolePropagationConfigInput'] 
slug: /tools/sdk/go/v2026/models/role-propagation-config-input
tags: ['SDK', 'Software Development Kit', 'RolePropagationConfigInput', 'V2026RolePropagationConfigInput']
---

# RolePropagationConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates if the Role Change Propagation process should be enabled for the tenant | [optional] [default to false]

## Methods

### NewRolePropagationConfigInput

`func NewRolePropagationConfigInput() *RolePropagationConfigInput`

NewRolePropagationConfigInput instantiates a new RolePropagationConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePropagationConfigInputWithDefaults

`func NewRolePropagationConfigInputWithDefaults() *RolePropagationConfigInput`

NewRolePropagationConfigInputWithDefaults instantiates a new RolePropagationConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RolePropagationConfigInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RolePropagationConfigInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RolePropagationConfigInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RolePropagationConfigInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


