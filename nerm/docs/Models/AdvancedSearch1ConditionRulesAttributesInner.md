---
id: nerm-advanced-search1-condition-rules-attributes-inner
title: AdvancedSearch1ConditionRulesAttributesInner
pagination_label: AdvancedSearch1ConditionRulesAttributesInner
sidebar_label: AdvancedSearch1ConditionRulesAttributesInner
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AdvancedSearch1ConditionRulesAttributesInner', 'NERMAdvancedSearch1ConditionRulesAttributesInner'] 
slug: /tools/sdk/go/nerm/models/advanced-search1-condition-rules-attributes-inner
tags: ['SDK', 'Software Development Kit', 'AdvancedSearch1ConditionRulesAttributesInner', 'NERMAdvancedSearch1ConditionRulesAttributesInner']
---

# AdvancedSearch1ConditionRulesAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ComparisonOperator** | **string** |  | 
**Value** | **string** |  | 
**ConditionObjectType** | **string** |  | 
**ConditionObjectId** | **string** |  | 
**SecondaryAttributeType** | Pointer to **string** |  | [optional] 
**SecondaryAttributeId** | Pointer to **string** |  | [optional] 
**SecondaryValue** | **string** |  | 
**TertiaryValue** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvancedSearch1ConditionRulesAttributesInner

`func NewAdvancedSearch1ConditionRulesAttributesInner(type_ string, comparisonOperator string, value string, conditionObjectType string, conditionObjectId string, secondaryValue string, ) *AdvancedSearch1ConditionRulesAttributesInner`

NewAdvancedSearch1ConditionRulesAttributesInner instantiates a new AdvancedSearch1ConditionRulesAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedSearch1ConditionRulesAttributesInnerWithDefaults

`func NewAdvancedSearch1ConditionRulesAttributesInnerWithDefaults() *AdvancedSearch1ConditionRulesAttributesInner`

NewAdvancedSearch1ConditionRulesAttributesInnerWithDefaults instantiates a new AdvancedSearch1ConditionRulesAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetType(v string)`

SetType sets Type field to given value.


### GetComparisonOperator

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.


### GetValue

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetConditionObjectType

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetConditionObjectType() string`

GetConditionObjectType returns the ConditionObjectType field if non-nil, zero value otherwise.

### GetConditionObjectTypeOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetConditionObjectTypeOk() (*string, bool)`

GetConditionObjectTypeOk returns a tuple with the ConditionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectType

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetConditionObjectType(v string)`

SetConditionObjectType sets ConditionObjectType field to given value.


### GetConditionObjectId

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetConditionObjectId() string`

GetConditionObjectId returns the ConditionObjectId field if non-nil, zero value otherwise.

### GetConditionObjectIdOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetConditionObjectIdOk() (*string, bool)`

GetConditionObjectIdOk returns a tuple with the ConditionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectId

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetConditionObjectId(v string)`

SetConditionObjectId sets ConditionObjectId field to given value.


### GetSecondaryAttributeType

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetSecondaryAttributeType() string`

GetSecondaryAttributeType returns the SecondaryAttributeType field if non-nil, zero value otherwise.

### GetSecondaryAttributeTypeOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetSecondaryAttributeTypeOk() (*string, bool)`

GetSecondaryAttributeTypeOk returns a tuple with the SecondaryAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAttributeType

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetSecondaryAttributeType(v string)`

SetSecondaryAttributeType sets SecondaryAttributeType field to given value.

### HasSecondaryAttributeType

`func (o *AdvancedSearch1ConditionRulesAttributesInner) HasSecondaryAttributeType() bool`

HasSecondaryAttributeType returns a boolean if a field has been set.

### GetSecondaryAttributeId

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetSecondaryAttributeId() string`

GetSecondaryAttributeId returns the SecondaryAttributeId field if non-nil, zero value otherwise.

### GetSecondaryAttributeIdOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetSecondaryAttributeIdOk() (*string, bool)`

GetSecondaryAttributeIdOk returns a tuple with the SecondaryAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAttributeId

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetSecondaryAttributeId(v string)`

SetSecondaryAttributeId sets SecondaryAttributeId field to given value.

### HasSecondaryAttributeId

`func (o *AdvancedSearch1ConditionRulesAttributesInner) HasSecondaryAttributeId() bool`

HasSecondaryAttributeId returns a boolean if a field has been set.

### GetSecondaryValue

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetSecondaryValue() string`

GetSecondaryValue returns the SecondaryValue field if non-nil, zero value otherwise.

### GetSecondaryValueOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetSecondaryValueOk() (*string, bool)`

GetSecondaryValueOk returns a tuple with the SecondaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryValue

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetSecondaryValue(v string)`

SetSecondaryValue sets SecondaryValue field to given value.


### GetTertiaryValue

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetTertiaryValue() string`

GetTertiaryValue returns the TertiaryValue field if non-nil, zero value otherwise.

### GetTertiaryValueOk

`func (o *AdvancedSearch1ConditionRulesAttributesInner) GetTertiaryValueOk() (*string, bool)`

GetTertiaryValueOk returns a tuple with the TertiaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTertiaryValue

`func (o *AdvancedSearch1ConditionRulesAttributesInner) SetTertiaryValue(v string)`

SetTertiaryValue sets TertiaryValue field to given value.

### HasTertiaryValue

`func (o *AdvancedSearch1ConditionRulesAttributesInner) HasTertiaryValue() bool`

HasTertiaryValue returns a boolean if a field has been set.


