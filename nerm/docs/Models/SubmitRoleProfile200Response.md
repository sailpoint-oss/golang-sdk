---
id: nerm-submit-role-profile200-response
title: SubmitRoleProfile200Response
pagination_label: SubmitRoleProfile200Response
sidebar_label: SubmitRoleProfile200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitRoleProfile200Response', 'NERMSubmitRoleProfile200Response'] 
slug: /tools/sdk/go/nerm/models/submit-role-profile200-response
tags: ['SDK', 'Software Development Kit', 'SubmitRoleProfile200Response', 'NERMSubmitRoleProfile200Response']
---

# SubmitRoleProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleProfile** | Pointer to [**RoleProfile**](role-profile) |  | [optional] 

## Methods

### NewSubmitRoleProfile200Response

`func NewSubmitRoleProfile200Response() *SubmitRoleProfile200Response`

NewSubmitRoleProfile200Response instantiates a new SubmitRoleProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRoleProfile200ResponseWithDefaults

`func NewSubmitRoleProfile200ResponseWithDefaults() *SubmitRoleProfile200Response`

NewSubmitRoleProfile200ResponseWithDefaults instantiates a new SubmitRoleProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleProfile

`func (o *SubmitRoleProfile200Response) GetRoleProfile() RoleProfile`

GetRoleProfile returns the RoleProfile field if non-nil, zero value otherwise.

### GetRoleProfileOk

`func (o *SubmitRoleProfile200Response) GetRoleProfileOk() (*RoleProfile, bool)`

GetRoleProfileOk returns a tuple with the RoleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleProfile

`func (o *SubmitRoleProfile200Response) SetRoleProfile(v RoleProfile)`

SetRoleProfile sets RoleProfile field to given value.

### HasRoleProfile

`func (o *SubmitRoleProfile200Response) HasRoleProfile() bool`

HasRoleProfile returns a boolean if a field has been set.


