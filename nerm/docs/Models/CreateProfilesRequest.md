---
id: nerm-create-profiles-request
title: CreateProfilesRequest
pagination_label: CreateProfilesRequest
sidebar_label: CreateProfilesRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateProfilesRequest', 'NERMCreateProfilesRequest'] 
slug: /tools/sdk/go/nerm/models/create-profiles-request
tags: ['SDK', 'Software Development Kit', 'CreateProfilesRequest', 'NERMCreateProfilesRequest']
---

# CreateProfilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]Profile1**](profile1) |  | [optional] 

## Methods

### NewCreateProfilesRequest

`func NewCreateProfilesRequest() *CreateProfilesRequest`

NewCreateProfilesRequest instantiates a new CreateProfilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfilesRequestWithDefaults

`func NewCreateProfilesRequestWithDefaults() *CreateProfilesRequest`

NewCreateProfilesRequestWithDefaults instantiates a new CreateProfilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *CreateProfilesRequest) GetProfiles() []Profile1`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *CreateProfilesRequest) GetProfilesOk() (*[]Profile1, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *CreateProfilesRequest) SetProfiles(v []Profile1)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *CreateProfilesRequest) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


