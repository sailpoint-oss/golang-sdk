---
id: nerm-submit-user-manager200-response
title: SubmitUserManager200Response
pagination_label: SubmitUserManager200Response
sidebar_label: SubmitUserManager200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserManager200Response', 'NERMSubmitUserManager200Response'] 
slug: /tools/sdk/go/nerm/models/submit-user-manager200-response
tags: ['SDK', 'Software Development Kit', 'SubmitUserManager200Response', 'NERMSubmitUserManager200Response']
---

# SubmitUserManager200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserManager** | Pointer to [**UserManager**](user-manager) |  | [optional] 

## Methods

### NewSubmitUserManager200Response

`func NewSubmitUserManager200Response() *SubmitUserManager200Response`

NewSubmitUserManager200Response instantiates a new SubmitUserManager200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserManager200ResponseWithDefaults

`func NewSubmitUserManager200ResponseWithDefaults() *SubmitUserManager200Response`

NewSubmitUserManager200ResponseWithDefaults instantiates a new SubmitUserManager200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserManager

`func (o *SubmitUserManager200Response) GetUserManager() UserManager`

GetUserManager returns the UserManager field if non-nil, zero value otherwise.

### GetUserManagerOk

`func (o *SubmitUserManager200Response) GetUserManagerOk() (*UserManager, bool)`

GetUserManagerOk returns a tuple with the UserManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManager

`func (o *SubmitUserManager200Response) SetUserManager(v UserManager)`

SetUserManager sets UserManager field to given value.

### HasUserManager

`func (o *SubmitUserManager200Response) HasUserManager() bool`

HasUserManager returns a boolean if a field has been set.


