---
id: nerm-update-profile-request-profile
title: UpdateProfileRequestProfile
pagination_label: UpdateProfileRequestProfile
sidebar_label: UpdateProfileRequestProfile
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UpdateProfileRequestProfile', 'NERMUpdateProfileRequestProfile'] 
slug: /tools/sdk/go/nerm/models/update-profile-request-profile
tags: ['SDK', 'Software Development Kit', 'UpdateProfileRequestProfile', 'NERMUpdateProfileRequestProfile']
---

# UpdateProfileRequestProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** | schema-mapped attributes to be updated | [optional] 

## Methods

### NewUpdateProfileRequestProfile

`func NewUpdateProfileRequestProfile() *UpdateProfileRequestProfile`

NewUpdateProfileRequestProfile instantiates a new UpdateProfileRequestProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileRequestProfileWithDefaults

`func NewUpdateProfileRequestProfileWithDefaults() *UpdateProfileRequestProfile`

NewUpdateProfileRequestProfileWithDefaults instantiates a new UpdateProfileRequestProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateProfileRequestProfile) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateProfileRequestProfile) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateProfileRequestProfile) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateProfileRequestProfile) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


