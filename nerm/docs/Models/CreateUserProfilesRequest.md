---
id: nerm-create-user-profiles-request
title: CreateUserProfilesRequest
pagination_label: CreateUserProfilesRequest
sidebar_label: CreateUserProfilesRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateUserProfilesRequest', 'NERMCreateUserProfilesRequest'] 
slug: /tools/sdk/go/nerm/models/create-user-profiles-request
tags: ['SDK', 'Software Development Kit', 'CreateUserProfilesRequest', 'NERMCreateUserProfilesRequest']
---

# CreateUserProfilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserProfiles** | Pointer to [**[]UserProfile1**](user-profile1) |  | [optional] 

## Methods

### NewCreateUserProfilesRequest

`func NewCreateUserProfilesRequest() *CreateUserProfilesRequest`

NewCreateUserProfilesRequest instantiates a new CreateUserProfilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserProfilesRequestWithDefaults

`func NewCreateUserProfilesRequestWithDefaults() *CreateUserProfilesRequest`

NewCreateUserProfilesRequestWithDefaults instantiates a new CreateUserProfilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserProfiles

`func (o *CreateUserProfilesRequest) GetUserProfiles() []UserProfile1`

GetUserProfiles returns the UserProfiles field if non-nil, zero value otherwise.

### GetUserProfilesOk

`func (o *CreateUserProfilesRequest) GetUserProfilesOk() (*[]UserProfile1, bool)`

GetUserProfilesOk returns a tuple with the UserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfiles

`func (o *CreateUserProfilesRequest) SetUserProfiles(v []UserProfile1)`

SetUserProfiles sets UserProfiles field to given value.

### HasUserProfiles

`func (o *CreateUserProfilesRequest) HasUserProfiles() bool`

HasUserProfiles returns a boolean if a field has been set.


