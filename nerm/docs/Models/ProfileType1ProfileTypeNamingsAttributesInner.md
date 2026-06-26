---
id: nerm-profile-type1-profile-type-namings-attributes-inner
title: ProfileType1ProfileTypeNamingsAttributesInner
pagination_label: ProfileType1ProfileTypeNamingsAttributesInner
sidebar_label: ProfileType1ProfileTypeNamingsAttributesInner
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileType1ProfileTypeNamingsAttributesInner', 'NERMProfileType1ProfileTypeNamingsAttributesInner'] 
slug: /tools/sdk/go/nerm/models/profile-type1-profile-type-namings-attributes-inner
tags: ['SDK', 'Software Development Kit', 'ProfileType1ProfileTypeNamingsAttributesInner', 'NERMProfileType1ProfileTypeNamingsAttributesInner']
---

# ProfileType1ProfileTypeNamingsAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttributeId** | Pointer to **string** | The ID of the associated ne attribute. | [optional] 
**Order** | Pointer to **int32** | The order that the namings are used in. | [optional] 

## Methods

### NewProfileType1ProfileTypeNamingsAttributesInner

`func NewProfileType1ProfileTypeNamingsAttributesInner() *ProfileType1ProfileTypeNamingsAttributesInner`

NewProfileType1ProfileTypeNamingsAttributesInner instantiates a new ProfileType1ProfileTypeNamingsAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileType1ProfileTypeNamingsAttributesInnerWithDefaults

`func NewProfileType1ProfileTypeNamingsAttributesInnerWithDefaults() *ProfileType1ProfileTypeNamingsAttributesInner`

NewProfileType1ProfileTypeNamingsAttributesInnerWithDefaults instantiates a new ProfileType1ProfileTypeNamingsAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttributeId

`func (o *ProfileType1ProfileTypeNamingsAttributesInner) GetNeAttributeId() string`

GetNeAttributeId returns the NeAttributeId field if non-nil, zero value otherwise.

### GetNeAttributeIdOk

`func (o *ProfileType1ProfileTypeNamingsAttributesInner) GetNeAttributeIdOk() (*string, bool)`

GetNeAttributeIdOk returns a tuple with the NeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeId

`func (o *ProfileType1ProfileTypeNamingsAttributesInner) SetNeAttributeId(v string)`

SetNeAttributeId sets NeAttributeId field to given value.

### HasNeAttributeId

`func (o *ProfileType1ProfileTypeNamingsAttributesInner) HasNeAttributeId() bool`

HasNeAttributeId returns a boolean if a field has been set.

### GetOrder

`func (o *ProfileType1ProfileTypeNamingsAttributesInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ProfileType1ProfileTypeNamingsAttributesInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ProfileType1ProfileTypeNamingsAttributesInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ProfileType1ProfileTypeNamingsAttributesInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


