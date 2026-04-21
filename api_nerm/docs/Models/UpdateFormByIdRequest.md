---
id: nerm-update-form-by-id-request
title: UpdateFormByIdRequest
pagination_label: UpdateFormByIdRequest
sidebar_label: UpdateFormByIdRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UpdateFormByIdRequest', 'NERMUpdateFormByIdRequest'] 
slug: /tools/sdk/go/nerm/models/update-form-by-id-request
tags: ['SDK', 'Software Development Kit', 'UpdateFormByIdRequest', 'NERMUpdateFormByIdRequest']
---

# UpdateFormByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**Form1**](form1) |  | [optional] 

## Methods

### NewUpdateFormByIdRequest

`func NewUpdateFormByIdRequest() *UpdateFormByIdRequest`

NewUpdateFormByIdRequest instantiates a new UpdateFormByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFormByIdRequestWithDefaults

`func NewUpdateFormByIdRequestWithDefaults() *UpdateFormByIdRequest`

NewUpdateFormByIdRequestWithDefaults instantiates a new UpdateFormByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *UpdateFormByIdRequest) GetRole() Form1`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateFormByIdRequest) GetRoleOk() (*Form1, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateFormByIdRequest) SetRole(v Form1)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateFormByIdRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


