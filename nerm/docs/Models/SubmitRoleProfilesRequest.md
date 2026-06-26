---
id: nerm-submit-role-profiles-request
title: SubmitRoleProfilesRequest
pagination_label: SubmitRoleProfilesRequest
sidebar_label: SubmitRoleProfilesRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitRoleProfilesRequest', 'NERMSubmitRoleProfilesRequest'] 
slug: /tools/sdk/go/nerm/models/submit-role-profiles-request
tags: ['SDK', 'Software Development Kit', 'SubmitRoleProfilesRequest', 'NERMSubmitRoleProfilesRequest']
---

# SubmitRoleProfilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleProfiles** | Pointer to [**[]RoleProfile1**](role-profile1) |  | [optional] 

## Methods

### NewSubmitRoleProfilesRequest

`func NewSubmitRoleProfilesRequest() *SubmitRoleProfilesRequest`

NewSubmitRoleProfilesRequest instantiates a new SubmitRoleProfilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRoleProfilesRequestWithDefaults

`func NewSubmitRoleProfilesRequestWithDefaults() *SubmitRoleProfilesRequest`

NewSubmitRoleProfilesRequestWithDefaults instantiates a new SubmitRoleProfilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleProfiles

`func (o *SubmitRoleProfilesRequest) GetRoleProfiles() []RoleProfile1`

GetRoleProfiles returns the RoleProfiles field if non-nil, zero value otherwise.

### GetRoleProfilesOk

`func (o *SubmitRoleProfilesRequest) GetRoleProfilesOk() (*[]RoleProfile1, bool)`

GetRoleProfilesOk returns a tuple with the RoleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleProfiles

`func (o *SubmitRoleProfilesRequest) SetRoleProfiles(v []RoleProfile1)`

SetRoleProfiles sets RoleProfiles field to given value.

### HasRoleProfiles

`func (o *SubmitRoleProfilesRequest) HasRoleProfiles() bool`

HasRoleProfiles returns a boolean if a field has been set.


