---
id: nerm-submit-user-role-request
title: SubmitUserRoleRequest
pagination_label: SubmitUserRoleRequest
sidebar_label: SubmitUserRoleRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserRoleRequest', 'NERMSubmitUserRoleRequest'] 
slug: /tools/sdk/go/nerm/models/submit-user-role-request
tags: ['SDK', 'Software Development Kit', 'SubmitUserRoleRequest', 'NERMSubmitUserRoleRequest']
---

# SubmitUserRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRole** | Pointer to [**UserRole1**](user-role1) |  | [optional] 

## Methods

### NewSubmitUserRoleRequest

`func NewSubmitUserRoleRequest() *SubmitUserRoleRequest`

NewSubmitUserRoleRequest instantiates a new SubmitUserRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserRoleRequestWithDefaults

`func NewSubmitUserRoleRequestWithDefaults() *SubmitUserRoleRequest`

NewSubmitUserRoleRequestWithDefaults instantiates a new SubmitUserRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRole

`func (o *SubmitUserRoleRequest) GetUserRole() UserRole1`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *SubmitUserRoleRequest) GetUserRoleOk() (*UserRole1, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *SubmitUserRoleRequest) SetUserRole(v UserRole1)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *SubmitUserRoleRequest) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.


