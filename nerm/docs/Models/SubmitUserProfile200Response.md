---
id: nerm-submit-user-profile200-response
title: SubmitUserProfile200Response
pagination_label: SubmitUserProfile200Response
sidebar_label: SubmitUserProfile200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserProfile200Response', 'NERMSubmitUserProfile200Response'] 
slug: /tools/sdk/go/nerm/models/submit-user-profile200-response
tags: ['SDK', 'Software Development Kit', 'SubmitUserProfile200Response', 'NERMSubmitUserProfile200Response']
---

# SubmitUserProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserProfile** | Pointer to [**UserProfile**](user-profile) |  | [optional] 

## Methods

### NewSubmitUserProfile200Response

`func NewSubmitUserProfile200Response() *SubmitUserProfile200Response`

NewSubmitUserProfile200Response instantiates a new SubmitUserProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserProfile200ResponseWithDefaults

`func NewSubmitUserProfile200ResponseWithDefaults() *SubmitUserProfile200Response`

NewSubmitUserProfile200ResponseWithDefaults instantiates a new SubmitUserProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserProfile

`func (o *SubmitUserProfile200Response) GetUserProfile() UserProfile`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *SubmitUserProfile200Response) GetUserProfileOk() (*UserProfile, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *SubmitUserProfile200Response) SetUserProfile(v UserProfile)`

SetUserProfile sets UserProfile field to given value.

### HasUserProfile

`func (o *SubmitUserProfile200Response) HasUserProfile() bool`

HasUserProfile returns a boolean if a field has been set.


