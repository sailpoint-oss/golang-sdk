---
id: nerm-login-workflow-options
title: LoginWorkflowOptions
pagination_label: LoginWorkflowOptions
sidebar_label: LoginWorkflowOptions
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'LoginWorkflowOptions', 'NERMLoginWorkflowOptions'] 
slug: /tools/sdk/go/nerm/models/login-workflow-options
tags: ['SDK', 'Software Development Kit', 'LoginWorkflowOptions', 'NERMLoginWorkflowOptions']
---

# LoginWorkflowOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | Pointer to **int32** | Used for login/password reset for when the password will expire. | [optional] 

## Methods

### NewLoginWorkflowOptions

`func NewLoginWorkflowOptions() *LoginWorkflowOptions`

NewLoginWorkflowOptions instantiates a new LoginWorkflowOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWorkflowOptionsWithDefaults

`func NewLoginWorkflowOptionsWithDefaults() *LoginWorkflowOptions`

NewLoginWorkflowOptionsWithDefaults instantiates a new LoginWorkflowOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *LoginWorkflowOptions) GetExpirationTime() int32`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *LoginWorkflowOptions) GetExpirationTimeOk() (*int32, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *LoginWorkflowOptions) SetExpirationTime(v int32)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *LoginWorkflowOptions) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.


