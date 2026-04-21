---
id: nerm-advanced-search1
title: AdvancedSearch1
pagination_label: AdvancedSearch1
sidebar_label: AdvancedSearch1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AdvancedSearch1', 'NERMAdvancedSearch1'] 
slug: /tools/sdk/go/nerm/models/advanced-search1
tags: ['SDK', 'Software Development Kit', 'AdvancedSearch1', 'NERMAdvancedSearch1']
---

# AdvancedSearch1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**ConditionRulesAttributes** | Pointer to [**[]AdvancedSearch1ConditionRulesAttributesInner**](advanced-search1-condition-rules-attributes-inner) |  | [optional] 

## Methods

### NewAdvancedSearch1

`func NewAdvancedSearch1() *AdvancedSearch1`

NewAdvancedSearch1 instantiates a new AdvancedSearch1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedSearch1WithDefaults

`func NewAdvancedSearch1WithDefaults() *AdvancedSearch1`

NewAdvancedSearch1WithDefaults instantiates a new AdvancedSearch1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *AdvancedSearch1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AdvancedSearch1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AdvancedSearch1) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AdvancedSearch1) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetConditionRulesAttributes

`func (o *AdvancedSearch1) GetConditionRulesAttributes() []AdvancedSearch1ConditionRulesAttributesInner`

GetConditionRulesAttributes returns the ConditionRulesAttributes field if non-nil, zero value otherwise.

### GetConditionRulesAttributesOk

`func (o *AdvancedSearch1) GetConditionRulesAttributesOk() (*[]AdvancedSearch1ConditionRulesAttributesInner, bool)`

GetConditionRulesAttributesOk returns a tuple with the ConditionRulesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionRulesAttributes

`func (o *AdvancedSearch1) SetConditionRulesAttributes(v []AdvancedSearch1ConditionRulesAttributesInner)`

SetConditionRulesAttributes sets ConditionRulesAttributes field to given value.

### HasConditionRulesAttributes

`func (o *AdvancedSearch1) HasConditionRulesAttributes() bool`

HasConditionRulesAttributes returns a boolean if a field has been set.


