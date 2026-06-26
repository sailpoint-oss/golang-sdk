---
id: nerm-submit-profile-type200-response
title: SubmitProfileType200Response
pagination_label: SubmitProfileType200Response
sidebar_label: SubmitProfileType200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitProfileType200Response', 'NERMSubmitProfileType200Response'] 
slug: /tools/sdk/go/nerm/models/submit-profile-type200-response
tags: ['SDK', 'Software Development Kit', 'SubmitProfileType200Response', 'NERMSubmitProfileType200Response']
---

# SubmitProfileType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileType** | Pointer to [**ProfileType**](profile-type) |  | [optional] 

## Methods

### NewSubmitProfileType200Response

`func NewSubmitProfileType200Response() *SubmitProfileType200Response`

NewSubmitProfileType200Response instantiates a new SubmitProfileType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitProfileType200ResponseWithDefaults

`func NewSubmitProfileType200ResponseWithDefaults() *SubmitProfileType200Response`

NewSubmitProfileType200ResponseWithDefaults instantiates a new SubmitProfileType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileType

`func (o *SubmitProfileType200Response) GetProfileType() ProfileType`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *SubmitProfileType200Response) GetProfileTypeOk() (*ProfileType, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *SubmitProfileType200Response) SetProfileType(v ProfileType)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *SubmitProfileType200Response) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.


