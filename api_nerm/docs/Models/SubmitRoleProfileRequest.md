---
id: nerm-submit-role-profile-request
title: SubmitRoleProfileRequest
pagination_label: SubmitRoleProfileRequest
sidebar_label: SubmitRoleProfileRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitRoleProfileRequest', 'NERMSubmitRoleProfileRequest'] 
slug: /tools/sdk/go/nerm/models/submit-role-profile-request
tags: ['SDK', 'Software Development Kit', 'SubmitRoleProfileRequest', 'NERMSubmitRoleProfileRequest']
---

# SubmitRoleProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleProfile** | Pointer to [**RoleProfile1**](role-profile1) |  | [optional] 

## Methods

### NewSubmitRoleProfileRequest

`func NewSubmitRoleProfileRequest() *SubmitRoleProfileRequest`

NewSubmitRoleProfileRequest instantiates a new SubmitRoleProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRoleProfileRequestWithDefaults

`func NewSubmitRoleProfileRequestWithDefaults() *SubmitRoleProfileRequest`

NewSubmitRoleProfileRequestWithDefaults instantiates a new SubmitRoleProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleProfile

`func (o *SubmitRoleProfileRequest) GetRoleProfile() RoleProfile1`

GetRoleProfile returns the RoleProfile field if non-nil, zero value otherwise.

### GetRoleProfileOk

`func (o *SubmitRoleProfileRequest) GetRoleProfileOk() (*RoleProfile1, bool)`

GetRoleProfileOk returns a tuple with the RoleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleProfile

`func (o *SubmitRoleProfileRequest) SetRoleProfile(v RoleProfile1)`

SetRoleProfile sets RoleProfile field to given value.

### HasRoleProfile

`func (o *SubmitRoleProfileRequest) HasRoleProfile() bool`

HasRoleProfile returns a boolean if a field has been set.


