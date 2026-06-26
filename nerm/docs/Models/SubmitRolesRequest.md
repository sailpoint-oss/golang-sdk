---
id: nerm-submit-roles-request
title: SubmitRolesRequest
pagination_label: SubmitRolesRequest
sidebar_label: SubmitRolesRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitRolesRequest', 'NERMSubmitRolesRequest'] 
slug: /tools/sdk/go/nerm/models/submit-roles-request
tags: ['SDK', 'Software Development Kit', 'SubmitRolesRequest', 'NERMSubmitRolesRequest']
---

# SubmitRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]Role1**](role1) |  | [optional] 

## Methods

### NewSubmitRolesRequest

`func NewSubmitRolesRequest() *SubmitRolesRequest`

NewSubmitRolesRequest instantiates a new SubmitRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRolesRequestWithDefaults

`func NewSubmitRolesRequestWithDefaults() *SubmitRolesRequest`

NewSubmitRolesRequestWithDefaults instantiates a new SubmitRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *SubmitRolesRequest) GetRoles() []Role1`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SubmitRolesRequest) GetRolesOk() (*[]Role1, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SubmitRolesRequest) SetRoles(v []Role1)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *SubmitRolesRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


