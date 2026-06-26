---
id: nerm-submit-profile-type-request
title: SubmitProfileTypeRequest
pagination_label: SubmitProfileTypeRequest
sidebar_label: SubmitProfileTypeRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitProfileTypeRequest', 'NERMSubmitProfileTypeRequest'] 
slug: /tools/sdk/go/nerm/models/submit-profile-type-request
tags: ['SDK', 'Software Development Kit', 'SubmitProfileTypeRequest', 'NERMSubmitProfileTypeRequest']
---

# SubmitProfileTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileType** | Pointer to [**ProfileType1**](profile-type1) |  | [optional] 

## Methods

### NewSubmitProfileTypeRequest

`func NewSubmitProfileTypeRequest() *SubmitProfileTypeRequest`

NewSubmitProfileTypeRequest instantiates a new SubmitProfileTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitProfileTypeRequestWithDefaults

`func NewSubmitProfileTypeRequestWithDefaults() *SubmitProfileTypeRequest`

NewSubmitProfileTypeRequestWithDefaults instantiates a new SubmitProfileTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileType

`func (o *SubmitProfileTypeRequest) GetProfileType() ProfileType1`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *SubmitProfileTypeRequest) GetProfileTypeOk() (*ProfileType1, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *SubmitProfileTypeRequest) SetProfileType(v ProfileType1)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *SubmitProfileTypeRequest) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.


