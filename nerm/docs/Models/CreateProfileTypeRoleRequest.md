---
id: nerm-create-profile-type-role-request
title: CreateProfileTypeRoleRequest
pagination_label: CreateProfileTypeRoleRequest
sidebar_label: CreateProfileTypeRoleRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateProfileTypeRoleRequest', 'NERMCreateProfileTypeRoleRequest'] 
slug: /tools/sdk/go/nerm/models/create-profile-type-role-request
tags: ['SDK', 'Software Development Kit', 'CreateProfileTypeRoleRequest', 'NERMCreateProfileTypeRoleRequest']
---

# CreateProfileTypeRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypeRole** | Pointer to [**ProfileTypeRoles1**](profile-type-roles1) |  | [optional] 

## Methods

### NewCreateProfileTypeRoleRequest

`func NewCreateProfileTypeRoleRequest() *CreateProfileTypeRoleRequest`

NewCreateProfileTypeRoleRequest instantiates a new CreateProfileTypeRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileTypeRoleRequestWithDefaults

`func NewCreateProfileTypeRoleRequestWithDefaults() *CreateProfileTypeRoleRequest`

NewCreateProfileTypeRoleRequestWithDefaults instantiates a new CreateProfileTypeRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypeRole

`func (o *CreateProfileTypeRoleRequest) GetProfileTypeRole() ProfileTypeRoles1`

GetProfileTypeRole returns the ProfileTypeRole field if non-nil, zero value otherwise.

### GetProfileTypeRoleOk

`func (o *CreateProfileTypeRoleRequest) GetProfileTypeRoleOk() (*ProfileTypeRoles1, bool)`

GetProfileTypeRoleOk returns a tuple with the ProfileTypeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeRole

`func (o *CreateProfileTypeRoleRequest) SetProfileTypeRole(v ProfileTypeRoles1)`

SetProfileTypeRole sets ProfileTypeRole field to given value.

### HasProfileTypeRole

`func (o *CreateProfileTypeRoleRequest) HasProfileTypeRole() bool`

HasProfileTypeRole returns a boolean if a field has been set.


