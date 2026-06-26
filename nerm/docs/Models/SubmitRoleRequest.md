---
id: nerm-submit-role-request
title: SubmitRoleRequest
pagination_label: SubmitRoleRequest
sidebar_label: SubmitRoleRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitRoleRequest', 'NERMSubmitRoleRequest'] 
slug: /tools/sdk/go/nerm/models/submit-role-request
tags: ['SDK', 'Software Development Kit', 'SubmitRoleRequest', 'NERMSubmitRoleRequest']
---

# SubmitRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**Role1**](role1) |  | [optional] 

## Methods

### NewSubmitRoleRequest

`func NewSubmitRoleRequest() *SubmitRoleRequest`

NewSubmitRoleRequest instantiates a new SubmitRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRoleRequestWithDefaults

`func NewSubmitRoleRequestWithDefaults() *SubmitRoleRequest`

NewSubmitRoleRequestWithDefaults instantiates a new SubmitRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *SubmitRoleRequest) GetRole() Role1`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SubmitRoleRequest) GetRoleOk() (*Role1, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SubmitRoleRequest) SetRole(v Role1)`

SetRole sets Role field to given value.

### HasRole

`func (o *SubmitRoleRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


