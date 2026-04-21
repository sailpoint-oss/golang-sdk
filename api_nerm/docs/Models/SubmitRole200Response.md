---
id: nerm-submit-role200-response
title: SubmitRole200Response
pagination_label: SubmitRole200Response
sidebar_label: SubmitRole200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitRole200Response', 'NERMSubmitRole200Response'] 
slug: /tools/sdk/go/nerm/models/submit-role200-response
tags: ['SDK', 'Software Development Kit', 'SubmitRole200Response', 'NERMSubmitRole200Response']
---

# SubmitRole200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**Role**](role) |  | [optional] 

## Methods

### NewSubmitRole200Response

`func NewSubmitRole200Response() *SubmitRole200Response`

NewSubmitRole200Response instantiates a new SubmitRole200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRole200ResponseWithDefaults

`func NewSubmitRole200ResponseWithDefaults() *SubmitRole200Response`

NewSubmitRole200ResponseWithDefaults instantiates a new SubmitRole200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *SubmitRole200Response) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SubmitRole200Response) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SubmitRole200Response) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *SubmitRole200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.


