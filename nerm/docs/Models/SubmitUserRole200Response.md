---
id: nerm-submit-user-role200-response
title: SubmitUserRole200Response
pagination_label: SubmitUserRole200Response
sidebar_label: SubmitUserRole200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserRole200Response', 'NERMSubmitUserRole200Response'] 
slug: /tools/sdk/go/nerm/models/submit-user-role200-response
tags: ['SDK', 'Software Development Kit', 'SubmitUserRole200Response', 'NERMSubmitUserRole200Response']
---

# SubmitUserRole200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRole** | Pointer to [**UserRole**](user-role) |  | [optional] 

## Methods

### NewSubmitUserRole200Response

`func NewSubmitUserRole200Response() *SubmitUserRole200Response`

NewSubmitUserRole200Response instantiates a new SubmitUserRole200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserRole200ResponseWithDefaults

`func NewSubmitUserRole200ResponseWithDefaults() *SubmitUserRole200Response`

NewSubmitUserRole200ResponseWithDefaults instantiates a new SubmitUserRole200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRole

`func (o *SubmitUserRole200Response) GetUserRole() UserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *SubmitUserRole200Response) GetUserRoleOk() (*UserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *SubmitUserRole200Response) SetUserRole(v UserRole)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *SubmitUserRole200Response) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.


