---
id: nerm-update-profile-request
title: UpdateProfileRequest
pagination_label: UpdateProfileRequest
sidebar_label: UpdateProfileRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UpdateProfileRequest', 'NERMUpdateProfileRequest'] 
slug: /tools/sdk/go/nerm/models/update-profile-request
tags: ['SDK', 'Software Development Kit', 'UpdateProfileRequest', 'NERMUpdateProfileRequest']
---

# UpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | [**UpdateProfileRequestProfile**](update-profile-request-profile) |  | 

## Methods

### NewUpdateProfileRequest

`func NewUpdateProfileRequest(profile UpdateProfileRequestProfile, ) *UpdateProfileRequest`

NewUpdateProfileRequest instantiates a new UpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileRequestWithDefaults

`func NewUpdateProfileRequestWithDefaults() *UpdateProfileRequest`

NewUpdateProfileRequestWithDefaults instantiates a new UpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *UpdateProfileRequest) GetProfile() UpdateProfileRequestProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateProfileRequest) GetProfileOk() (*UpdateProfileRequestProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateProfileRequest) SetProfile(v UpdateProfileRequestProfile)`

SetProfile sets Profile field to given value.



