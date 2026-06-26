---
id: nerm-submit-profile-request
title: SubmitProfileRequest
pagination_label: SubmitProfileRequest
sidebar_label: SubmitProfileRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitProfileRequest', 'NERMSubmitProfileRequest'] 
slug: /tools/sdk/go/nerm/models/submit-profile-request
tags: ['SDK', 'Software Development Kit', 'SubmitProfileRequest', 'NERMSubmitProfileRequest']
---

# SubmitProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**Profile1**](profile1) |  | [optional] 

## Methods

### NewSubmitProfileRequest

`func NewSubmitProfileRequest() *SubmitProfileRequest`

NewSubmitProfileRequest instantiates a new SubmitProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitProfileRequestWithDefaults

`func NewSubmitProfileRequestWithDefaults() *SubmitProfileRequest`

NewSubmitProfileRequestWithDefaults instantiates a new SubmitProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *SubmitProfileRequest) GetProfile() Profile1`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SubmitProfileRequest) GetProfileOk() (*Profile1, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SubmitProfileRequest) SetProfile(v Profile1)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SubmitProfileRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


