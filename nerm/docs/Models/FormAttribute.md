---
id: nerm-form-attribute
title: FormAttribute
pagination_label: FormAttribute
sidebar_label: FormAttribute
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'FormAttribute', 'NERMFormAttribute'] 
slug: /tools/sdk/go/nerm/models/form-attribute
tags: ['SDK', 'Software Development Kit', 'FormAttribute', 'NERMFormAttribute']
---

# FormAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormId** | Pointer to **string** | The id of the form | [optional] 
**NeAttributeId** | Pointer to **string** | The id of the attribute | [optional] 
**AttrType** | Pointer to **string** | The attribute type | [optional] 
**Order** | Pointer to **int32** | The ordinal position of the attribute on the Form.  The order value determines the order or sequence the Form values are presented in the user interface. Each FormAttribute on a Form must have a unique order value. Order valuess can start at zero (0), but often start at one (1). The FormAttribute with order 1 is presented before the FormAttribute with order 2, and so on. Gaps in the order can exist and the system ignores them. | [optional] 
**Id** | Pointer to **string** | The id of the form attribute | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The date-time the record created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **SailPointTime** | The date-time the record was last updated. | [optional] [readonly] 

## Methods

### NewFormAttribute

`func NewFormAttribute() *FormAttribute`

NewFormAttribute instantiates a new FormAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormAttributeWithDefaults

`func NewFormAttributeWithDefaults() *FormAttribute`

NewFormAttributeWithDefaults instantiates a new FormAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormId

`func (o *FormAttribute) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *FormAttribute) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *FormAttribute) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *FormAttribute) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetNeAttributeId

`func (o *FormAttribute) GetNeAttributeId() string`

GetNeAttributeId returns the NeAttributeId field if non-nil, zero value otherwise.

### GetNeAttributeIdOk

`func (o *FormAttribute) GetNeAttributeIdOk() (*string, bool)`

GetNeAttributeIdOk returns a tuple with the NeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeId

`func (o *FormAttribute) SetNeAttributeId(v string)`

SetNeAttributeId sets NeAttributeId field to given value.

### HasNeAttributeId

`func (o *FormAttribute) HasNeAttributeId() bool`

HasNeAttributeId returns a boolean if a field has been set.

### GetAttrType

`func (o *FormAttribute) GetAttrType() string`

GetAttrType returns the AttrType field if non-nil, zero value otherwise.

### GetAttrTypeOk

`func (o *FormAttribute) GetAttrTypeOk() (*string, bool)`

GetAttrTypeOk returns a tuple with the AttrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrType

`func (o *FormAttribute) SetAttrType(v string)`

SetAttrType sets AttrType field to given value.

### HasAttrType

`func (o *FormAttribute) HasAttrType() bool`

HasAttrType returns a boolean if a field has been set.

### GetOrder

`func (o *FormAttribute) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FormAttribute) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FormAttribute) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FormAttribute) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetId

`func (o *FormAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FormAttribute) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormAttribute) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormAttribute) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FormAttribute) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FormAttribute) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormAttribute) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormAttribute) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FormAttribute) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


