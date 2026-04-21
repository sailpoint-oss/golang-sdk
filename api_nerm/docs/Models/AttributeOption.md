---
id: nerm-attribute-option
title: AttributeOption
pagination_label: AttributeOption
sidebar_label: AttributeOption
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AttributeOption', 'NERMAttributeOption'] 
slug: /tools/sdk/go/nerm/models/attribute-option
tags: ['SDK', 'Software Development Kit', 'AttributeOption', 'NERMAttributeOption']
---

# AttributeOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**NeAttributeId** | Pointer to **string** |  | [optional] 
**Option** | Pointer to **string** |  | [optional] 

## Methods

### NewAttributeOption

`func NewAttributeOption() *AttributeOption`

NewAttributeOption instantiates a new AttributeOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeOptionWithDefaults

`func NewAttributeOptionWithDefaults() *AttributeOption`

NewAttributeOptionWithDefaults instantiates a new AttributeOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttributeOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeOption) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttributeOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *AttributeOption) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AttributeOption) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AttributeOption) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AttributeOption) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetNeAttributeId

`func (o *AttributeOption) GetNeAttributeId() string`

GetNeAttributeId returns the NeAttributeId field if non-nil, zero value otherwise.

### GetNeAttributeIdOk

`func (o *AttributeOption) GetNeAttributeIdOk() (*string, bool)`

GetNeAttributeIdOk returns a tuple with the NeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeId

`func (o *AttributeOption) SetNeAttributeId(v string)`

SetNeAttributeId sets NeAttributeId field to given value.

### HasNeAttributeId

`func (o *AttributeOption) HasNeAttributeId() bool`

HasNeAttributeId returns a boolean if a field has been set.

### GetOption

`func (o *AttributeOption) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *AttributeOption) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *AttributeOption) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *AttributeOption) HasOption() bool`

HasOption returns a boolean if a field has been set.


