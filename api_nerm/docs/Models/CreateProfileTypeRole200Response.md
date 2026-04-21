---
id: nerm-create-profile-type-role200-response
title: CreateProfileTypeRole200Response
pagination_label: CreateProfileTypeRole200Response
sidebar_label: CreateProfileTypeRole200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateProfileTypeRole200Response', 'NERMCreateProfileTypeRole200Response'] 
slug: /tools/sdk/go/nerm/models/create-profile-type-role200-response
tags: ['SDK', 'Software Development Kit', 'CreateProfileTypeRole200Response', 'NERMCreateProfileTypeRole200Response']
---

# CreateProfileTypeRole200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypeRoles** | Pointer to [**ProfileTypeRoles**](profile-type-roles) |  | [optional] 

## Methods

### NewCreateProfileTypeRole200Response

`func NewCreateProfileTypeRole200Response() *CreateProfileTypeRole200Response`

NewCreateProfileTypeRole200Response instantiates a new CreateProfileTypeRole200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileTypeRole200ResponseWithDefaults

`func NewCreateProfileTypeRole200ResponseWithDefaults() *CreateProfileTypeRole200Response`

NewCreateProfileTypeRole200ResponseWithDefaults instantiates a new CreateProfileTypeRole200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypeRoles

`func (o *CreateProfileTypeRole200Response) GetProfileTypeRoles() ProfileTypeRoles`

GetProfileTypeRoles returns the ProfileTypeRoles field if non-nil, zero value otherwise.

### GetProfileTypeRolesOk

`func (o *CreateProfileTypeRole200Response) GetProfileTypeRolesOk() (*ProfileTypeRoles, bool)`

GetProfileTypeRolesOk returns a tuple with the ProfileTypeRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeRoles

`func (o *CreateProfileTypeRole200Response) SetProfileTypeRoles(v ProfileTypeRoles)`

SetProfileTypeRoles sets ProfileTypeRoles field to given value.

### HasProfileTypeRoles

`func (o *CreateProfileTypeRole200Response) HasProfileTypeRoles() bool`

HasProfileTypeRoles returns a boolean if a field has been set.


