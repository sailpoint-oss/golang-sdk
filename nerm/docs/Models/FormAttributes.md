---
id: nerm-form-attributes
title: FormAttributes
pagination_label: FormAttributes
sidebar_label: FormAttributes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'FormAttributes', 'NERMFormAttributes'] 
slug: /tools/sdk/go/nerm/models/form-attributes
tags: ['SDK', 'Software Development Kit', 'FormAttributes', 'NERMFormAttributes']
---

# FormAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormId** | Pointer to **string** | The id of the form | [optional] 
**NeAttributeId** | Pointer to **string** | The id of the attribute | [optional] 
**AttrType** | Pointer to **string** | The attribute type | [optional] 
**Order** | Pointer to **int32** | The ordinal position of the attribute on the Form.  The order value determines the order or sequence the Form values are presented in the user interface. Each FormAttribute on a Form must have a unique order value. Order valuess can start at zero (0), but often start at one (1). The FormAttribute with order 1 is presented before the FormAttribute with order 2, and so on. Gaps in the order can exist and the system ignores them. | [optional] 

## Methods

### NewFormAttributes

`func NewFormAttributes() *FormAttributes`

NewFormAttributes instantiates a new FormAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormAttributesWithDefaults

`func NewFormAttributesWithDefaults() *FormAttributes`

NewFormAttributesWithDefaults instantiates a new FormAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormId

`func (o *FormAttributes) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *FormAttributes) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *FormAttributes) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *FormAttributes) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetNeAttributeId

`func (o *FormAttributes) GetNeAttributeId() string`

GetNeAttributeId returns the NeAttributeId field if non-nil, zero value otherwise.

### GetNeAttributeIdOk

`func (o *FormAttributes) GetNeAttributeIdOk() (*string, bool)`

GetNeAttributeIdOk returns a tuple with the NeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeId

`func (o *FormAttributes) SetNeAttributeId(v string)`

SetNeAttributeId sets NeAttributeId field to given value.

### HasNeAttributeId

`func (o *FormAttributes) HasNeAttributeId() bool`

HasNeAttributeId returns a boolean if a field has been set.

### GetAttrType

`func (o *FormAttributes) GetAttrType() string`

GetAttrType returns the AttrType field if non-nil, zero value otherwise.

### GetAttrTypeOk

`func (o *FormAttributes) GetAttrTypeOk() (*string, bool)`

GetAttrTypeOk returns a tuple with the AttrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrType

`func (o *FormAttributes) SetAttrType(v string)`

SetAttrType sets AttrType field to given value.

### HasAttrType

`func (o *FormAttributes) HasAttrType() bool`

HasAttrType returns a boolean if a field has been set.

### GetOrder

`func (o *FormAttributes) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FormAttributes) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FormAttributes) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FormAttributes) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


