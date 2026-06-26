---
id: nerm-submit-user-profile-request
title: SubmitUserProfileRequest
pagination_label: SubmitUserProfileRequest
sidebar_label: SubmitUserProfileRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserProfileRequest', 'NERMSubmitUserProfileRequest'] 
slug: /tools/sdk/go/nerm/models/submit-user-profile-request
tags: ['SDK', 'Software Development Kit', 'SubmitUserProfileRequest', 'NERMSubmitUserProfileRequest']
---

# SubmitUserProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserProfile** | Pointer to [**UserProfile1**](user-profile1) |  | [optional] 

## Methods

### NewSubmitUserProfileRequest

`func NewSubmitUserProfileRequest() *SubmitUserProfileRequest`

NewSubmitUserProfileRequest instantiates a new SubmitUserProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserProfileRequestWithDefaults

`func NewSubmitUserProfileRequestWithDefaults() *SubmitUserProfileRequest`

NewSubmitUserProfileRequestWithDefaults instantiates a new SubmitUserProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserProfile

`func (o *SubmitUserProfileRequest) GetUserProfile() UserProfile1`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *SubmitUserProfileRequest) GetUserProfileOk() (*UserProfile1, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *SubmitUserProfileRequest) SetUserProfile(v UserProfile1)`

SetUserProfile sets UserProfile field to given value.

### HasUserProfile

`func (o *SubmitUserProfileRequest) HasUserProfile() bool`

HasUserProfile returns a boolean if a field has been set.


