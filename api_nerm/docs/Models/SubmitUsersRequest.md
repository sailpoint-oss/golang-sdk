---
id: nerm-submit-users-request
title: SubmitUsersRequest
pagination_label: SubmitUsersRequest
sidebar_label: SubmitUsersRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUsersRequest', 'NERMSubmitUsersRequest'] 
slug: /tools/sdk/go/nerm/models/submit-users-request
tags: ['SDK', 'Software Development Kit', 'SubmitUsersRequest', 'NERMSubmitUsersRequest']
---

# SubmitUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]User1**](user1) |  | [optional] 

## Methods

### NewSubmitUsersRequest

`func NewSubmitUsersRequest() *SubmitUsersRequest`

NewSubmitUsersRequest instantiates a new SubmitUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUsersRequestWithDefaults

`func NewSubmitUsersRequestWithDefaults() *SubmitUsersRequest`

NewSubmitUsersRequestWithDefaults instantiates a new SubmitUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *SubmitUsersRequest) GetUsers() []User1`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SubmitUsersRequest) GetUsersOk() (*[]User1, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SubmitUsersRequest) SetUsers(v []User1)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *SubmitUsersRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


