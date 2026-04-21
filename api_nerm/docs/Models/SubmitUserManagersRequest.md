---
id: nerm-submit-user-managers-request
title: SubmitUserManagersRequest
pagination_label: SubmitUserManagersRequest
sidebar_label: SubmitUserManagersRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserManagersRequest', 'NERMSubmitUserManagersRequest'] 
slug: /tools/sdk/go/nerm/models/submit-user-managers-request
tags: ['SDK', 'Software Development Kit', 'SubmitUserManagersRequest', 'NERMSubmitUserManagersRequest']
---

# SubmitUserManagersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserManagers** | Pointer to [**[]UserManager1**](user-manager1) |  | [optional] 

## Methods

### NewSubmitUserManagersRequest

`func NewSubmitUserManagersRequest() *SubmitUserManagersRequest`

NewSubmitUserManagersRequest instantiates a new SubmitUserManagersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserManagersRequestWithDefaults

`func NewSubmitUserManagersRequestWithDefaults() *SubmitUserManagersRequest`

NewSubmitUserManagersRequestWithDefaults instantiates a new SubmitUserManagersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserManagers

`func (o *SubmitUserManagersRequest) GetUserManagers() []UserManager1`

GetUserManagers returns the UserManagers field if non-nil, zero value otherwise.

### GetUserManagersOk

`func (o *SubmitUserManagersRequest) GetUserManagersOk() (*[]UserManager1, bool)`

GetUserManagersOk returns a tuple with the UserManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagers

`func (o *SubmitUserManagersRequest) SetUserManagers(v []UserManager1)`

SetUserManagers sets UserManagers field to given value.

### HasUserManagers

`func (o *SubmitUserManagersRequest) HasUserManagers() bool`

HasUserManagers returns a boolean if a field has been set.


