---
id: nerm-submit-user-request
title: SubmitUserRequest
pagination_label: SubmitUserRequest
sidebar_label: SubmitUserRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserRequest', 'NERMSubmitUserRequest'] 
slug: /tools/sdk/go/nerm/models/submit-user-request
tags: ['SDK', 'Software Development Kit', 'SubmitUserRequest', 'NERMSubmitUserRequest']
---

# SubmitUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User1**](user1) |  | [optional] 

## Methods

### NewSubmitUserRequest

`func NewSubmitUserRequest() *SubmitUserRequest`

NewSubmitUserRequest instantiates a new SubmitUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserRequestWithDefaults

`func NewSubmitUserRequestWithDefaults() *SubmitUserRequest`

NewSubmitUserRequestWithDefaults instantiates a new SubmitUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *SubmitUserRequest) GetUser() User1`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SubmitUserRequest) GetUserOk() (*User1, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SubmitUserRequest) SetUser(v User1)`

SetUser sets User field to given value.

### HasUser

`func (o *SubmitUserRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


