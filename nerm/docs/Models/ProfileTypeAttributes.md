---
id: nerm-profile-type-attributes
title: ProfileTypeAttributes
pagination_label: ProfileTypeAttributes
sidebar_label: ProfileTypeAttributes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileTypeAttributes', 'NERMProfileTypeAttributes'] 
slug: /tools/sdk/go/nerm/models/profile-type-attributes
tags: ['SDK', 'Software Development Kit', 'ProfileTypeAttributes', 'NERMProfileTypeAttributes']
---

# ProfileTypeAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | How many ne attribute records exist | [optional] 
**Records** | Pointer to [**[]ProfileTypeAttribute**](profile-type-attribute) |  | [optional] 

## Methods

### NewProfileTypeAttributes

`func NewProfileTypeAttributes() *ProfileTypeAttributes`

NewProfileTypeAttributes instantiates a new ProfileTypeAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileTypeAttributesWithDefaults

`func NewProfileTypeAttributesWithDefaults() *ProfileTypeAttributes`

NewProfileTypeAttributesWithDefaults instantiates a new ProfileTypeAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ProfileTypeAttributes) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ProfileTypeAttributes) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ProfileTypeAttributes) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ProfileTypeAttributes) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRecords

`func (o *ProfileTypeAttributes) GetRecords() []ProfileTypeAttribute`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ProfileTypeAttributes) GetRecordsOk() (*[]ProfileTypeAttribute, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ProfileTypeAttributes) SetRecords(v []ProfileTypeAttribute)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ProfileTypeAttributes) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


