---
id: nerm-profile-type-profile-type-namings-inner
title: ProfileTypeProfileTypeNamingsInner
pagination_label: ProfileTypeProfileTypeNamingsInner
sidebar_label: ProfileTypeProfileTypeNamingsInner
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileTypeProfileTypeNamingsInner', 'NERMProfileTypeProfileTypeNamingsInner'] 
slug: /tools/sdk/go/nerm/models/profile-type-profile-type-namings-inner
tags: ['SDK', 'Software Development Kit', 'ProfileTypeProfileTypeNamingsInner', 'NERMProfileTypeProfileTypeNamingsInner']
---

# ProfileTypeProfileTypeNamingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the profile type naming. | [optional] 
**Uid** | Pointer to **string** | The user-specified identifier of the profile type naming. | [optional] 
**ProfileTypeId** | Pointer to **string** | The ID of the associated profile type. | [optional] 
**NeAttributeId** | Pointer to **string** | The ID of the associated ne attribute. | [optional] 
**Order** | Pointer to **int32** | The order that the namings are used in. | [optional] 

## Methods

### NewProfileTypeProfileTypeNamingsInner

`func NewProfileTypeProfileTypeNamingsInner() *ProfileTypeProfileTypeNamingsInner`

NewProfileTypeProfileTypeNamingsInner instantiates a new ProfileTypeProfileTypeNamingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileTypeProfileTypeNamingsInnerWithDefaults

`func NewProfileTypeProfileTypeNamingsInnerWithDefaults() *ProfileTypeProfileTypeNamingsInner`

NewProfileTypeProfileTypeNamingsInnerWithDefaults instantiates a new ProfileTypeProfileTypeNamingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileTypeProfileTypeNamingsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileTypeProfileTypeNamingsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileTypeProfileTypeNamingsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileTypeProfileTypeNamingsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *ProfileTypeProfileTypeNamingsInner) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfileTypeProfileTypeNamingsInner) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfileTypeProfileTypeNamingsInner) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfileTypeProfileTypeNamingsInner) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *ProfileTypeProfileTypeNamingsInner) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *ProfileTypeProfileTypeNamingsInner) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *ProfileTypeProfileTypeNamingsInner) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *ProfileTypeProfileTypeNamingsInner) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.

### GetNeAttributeId

`func (o *ProfileTypeProfileTypeNamingsInner) GetNeAttributeId() string`

GetNeAttributeId returns the NeAttributeId field if non-nil, zero value otherwise.

### GetNeAttributeIdOk

`func (o *ProfileTypeProfileTypeNamingsInner) GetNeAttributeIdOk() (*string, bool)`

GetNeAttributeIdOk returns a tuple with the NeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeId

`func (o *ProfileTypeProfileTypeNamingsInner) SetNeAttributeId(v string)`

SetNeAttributeId sets NeAttributeId field to given value.

### HasNeAttributeId

`func (o *ProfileTypeProfileTypeNamingsInner) HasNeAttributeId() bool`

HasNeAttributeId returns a boolean if a field has been set.

### GetOrder

`func (o *ProfileTypeProfileTypeNamingsInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ProfileTypeProfileTypeNamingsInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ProfileTypeProfileTypeNamingsInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ProfileTypeProfileTypeNamingsInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


