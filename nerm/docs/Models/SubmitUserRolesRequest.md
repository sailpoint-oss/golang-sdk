---
id: nerm-submit-user-roles-request
title: SubmitUserRolesRequest
pagination_label: SubmitUserRolesRequest
sidebar_label: SubmitUserRolesRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserRolesRequest', 'NERMSubmitUserRolesRequest'] 
slug: /tools/sdk/go/nerm/models/submit-user-roles-request
tags: ['SDK', 'Software Development Kit', 'SubmitUserRolesRequest', 'NERMSubmitUserRolesRequest']
---

# SubmitUserRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRoles** | Pointer to [**[]UserRole1**](user-role1) |  | [optional] 

## Methods

### NewSubmitUserRolesRequest

`func NewSubmitUserRolesRequest() *SubmitUserRolesRequest`

NewSubmitUserRolesRequest instantiates a new SubmitUserRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserRolesRequestWithDefaults

`func NewSubmitUserRolesRequestWithDefaults() *SubmitUserRolesRequest`

NewSubmitUserRolesRequestWithDefaults instantiates a new SubmitUserRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRoles

`func (o *SubmitUserRolesRequest) GetUserRoles() []UserRole1`

GetUserRoles returns the UserRoles field if non-nil, zero value otherwise.

### GetUserRolesOk

`func (o *SubmitUserRolesRequest) GetUserRolesOk() (*[]UserRole1, bool)`

GetUserRolesOk returns a tuple with the UserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoles

`func (o *SubmitUserRolesRequest) SetUserRoles(v []UserRole1)`

SetUserRoles sets UserRoles field to given value.

### HasUserRoles

`func (o *SubmitUserRolesRequest) HasUserRoles() bool`

HasUserRoles returns a boolean if a field has been set.


