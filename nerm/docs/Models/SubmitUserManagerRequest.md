---
id: nerm-submit-user-manager-request
title: SubmitUserManagerRequest
pagination_label: SubmitUserManagerRequest
sidebar_label: SubmitUserManagerRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserManagerRequest', 'NERMSubmitUserManagerRequest'] 
slug: /tools/sdk/go/nerm/models/submit-user-manager-request
tags: ['SDK', 'Software Development Kit', 'SubmitUserManagerRequest', 'NERMSubmitUserManagerRequest']
---

# SubmitUserManagerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserManager** | Pointer to [**UserManager1**](user-manager1) |  | [optional] 

## Methods

### NewSubmitUserManagerRequest

`func NewSubmitUserManagerRequest() *SubmitUserManagerRequest`

NewSubmitUserManagerRequest instantiates a new SubmitUserManagerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserManagerRequestWithDefaults

`func NewSubmitUserManagerRequestWithDefaults() *SubmitUserManagerRequest`

NewSubmitUserManagerRequestWithDefaults instantiates a new SubmitUserManagerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserManager

`func (o *SubmitUserManagerRequest) GetUserManager() UserManager1`

GetUserManager returns the UserManager field if non-nil, zero value otherwise.

### GetUserManagerOk

`func (o *SubmitUserManagerRequest) GetUserManagerOk() (*UserManager1, bool)`

GetUserManagerOk returns a tuple with the UserManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManager

`func (o *SubmitUserManagerRequest) SetUserManager(v UserManager1)`

SetUserManager sets UserManager field to given value.

### HasUserManager

`func (o *SubmitUserManagerRequest) HasUserManager() bool`

HasUserManager returns a boolean if a field has been set.


