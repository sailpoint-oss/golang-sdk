---
id: nerm-submit-user200-response
title: SubmitUser200Response
pagination_label: SubmitUser200Response
sidebar_label: SubmitUser200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUser200Response', 'NERMSubmitUser200Response'] 
slug: /tools/sdk/go/nerm/models/submit-user200-response
tags: ['SDK', 'Software Development Kit', 'SubmitUser200Response', 'NERMSubmitUser200Response']
---

# SubmitUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](user) |  | [optional] 

## Methods

### NewSubmitUser200Response

`func NewSubmitUser200Response() *SubmitUser200Response`

NewSubmitUser200Response instantiates a new SubmitUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUser200ResponseWithDefaults

`func NewSubmitUser200ResponseWithDefaults() *SubmitUser200Response`

NewSubmitUser200ResponseWithDefaults instantiates a new SubmitUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *SubmitUser200Response) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SubmitUser200Response) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SubmitUser200Response) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *SubmitUser200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.


