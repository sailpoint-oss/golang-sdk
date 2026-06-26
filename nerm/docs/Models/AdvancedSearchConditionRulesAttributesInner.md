---
id: nerm-advanced-search-condition-rules-attributes-inner
title: AdvancedSearchConditionRulesAttributesInner
pagination_label: AdvancedSearchConditionRulesAttributesInner
sidebar_label: AdvancedSearchConditionRulesAttributesInner
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AdvancedSearchConditionRulesAttributesInner', 'NERMAdvancedSearchConditionRulesAttributesInner'] 
slug: /tools/sdk/go/nerm/models/advanced-search-condition-rules-attributes-inner
tags: ['SDK', 'Software Development Kit', 'AdvancedSearchConditionRulesAttributesInner', 'NERMAdvancedSearchConditionRulesAttributesInner']
---

# AdvancedSearchConditionRulesAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
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

### NewAdvancedSearchConditionRulesAttributesInner

`func NewAdvancedSearchConditionRulesAttributesInner(type_ string, comparisonOperator string, value string, conditionObjectType string, conditionObjectId string, secondaryValue string, ) *AdvancedSearchConditionRulesAttributesInner`

NewAdvancedSearchConditionRulesAttributesInner instantiates a new AdvancedSearchConditionRulesAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedSearchConditionRulesAttributesInnerWithDefaults

`func NewAdvancedSearchConditionRulesAttributesInnerWithDefaults() *AdvancedSearchConditionRulesAttributesInner`

NewAdvancedSearchConditionRulesAttributesInnerWithDefaults instantiates a new AdvancedSearchConditionRulesAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdvancedSearchConditionRulesAttributesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedSearchConditionRulesAttributesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedSearchConditionRulesAttributesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *AdvancedSearchConditionRulesAttributesInner) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AdvancedSearchConditionRulesAttributesInner) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AdvancedSearchConditionRulesAttributesInner) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *AdvancedSearchConditionRulesAttributesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvancedSearchConditionRulesAttributesInner) SetType(v string)`

SetType sets Type field to given value.


### GetComparisonOperator

`func (o *AdvancedSearchConditionRulesAttributesInner) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *AdvancedSearchConditionRulesAttributesInner) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.


### GetValue

`func (o *AdvancedSearchConditionRulesAttributesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AdvancedSearchConditionRulesAttributesInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetConditionObjectType

`func (o *AdvancedSearchConditionRulesAttributesInner) GetConditionObjectType() string`

GetConditionObjectType returns the ConditionObjectType field if non-nil, zero value otherwise.

### GetConditionObjectTypeOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetConditionObjectTypeOk() (*string, bool)`

GetConditionObjectTypeOk returns a tuple with the ConditionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectType

`func (o *AdvancedSearchConditionRulesAttributesInner) SetConditionObjectType(v string)`

SetConditionObjectType sets ConditionObjectType field to given value.


### GetConditionObjectId

`func (o *AdvancedSearchConditionRulesAttributesInner) GetConditionObjectId() string`

GetConditionObjectId returns the ConditionObjectId field if non-nil, zero value otherwise.

### GetConditionObjectIdOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetConditionObjectIdOk() (*string, bool)`

GetConditionObjectIdOk returns a tuple with the ConditionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectId

`func (o *AdvancedSearchConditionRulesAttributesInner) SetConditionObjectId(v string)`

SetConditionObjectId sets ConditionObjectId field to given value.


### GetSecondaryAttributeType

`func (o *AdvancedSearchConditionRulesAttributesInner) GetSecondaryAttributeType() string`

GetSecondaryAttributeType returns the SecondaryAttributeType field if non-nil, zero value otherwise.

### GetSecondaryAttributeTypeOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetSecondaryAttributeTypeOk() (*string, bool)`

GetSecondaryAttributeTypeOk returns a tuple with the SecondaryAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAttributeType

`func (o *AdvancedSearchConditionRulesAttributesInner) SetSecondaryAttributeType(v string)`

SetSecondaryAttributeType sets SecondaryAttributeType field to given value.

### HasSecondaryAttributeType

`func (o *AdvancedSearchConditionRulesAttributesInner) HasSecondaryAttributeType() bool`

HasSecondaryAttributeType returns a boolean if a field has been set.

### GetSecondaryAttributeId

`func (o *AdvancedSearchConditionRulesAttributesInner) GetSecondaryAttributeId() string`

GetSecondaryAttributeId returns the SecondaryAttributeId field if non-nil, zero value otherwise.

### GetSecondaryAttributeIdOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetSecondaryAttributeIdOk() (*string, bool)`

GetSecondaryAttributeIdOk returns a tuple with the SecondaryAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAttributeId

`func (o *AdvancedSearchConditionRulesAttributesInner) SetSecondaryAttributeId(v string)`

SetSecondaryAttributeId sets SecondaryAttributeId field to given value.

### HasSecondaryAttributeId

`func (o *AdvancedSearchConditionRulesAttributesInner) HasSecondaryAttributeId() bool`

HasSecondaryAttributeId returns a boolean if a field has been set.

### GetSecondaryValue

`func (o *AdvancedSearchConditionRulesAttributesInner) GetSecondaryValue() string`

GetSecondaryValue returns the SecondaryValue field if non-nil, zero value otherwise.

### GetSecondaryValueOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetSecondaryValueOk() (*string, bool)`

GetSecondaryValueOk returns a tuple with the SecondaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryValue

`func (o *AdvancedSearchConditionRulesAttributesInner) SetSecondaryValue(v string)`

SetSecondaryValue sets SecondaryValue field to given value.


### GetTertiaryValue

`func (o *AdvancedSearchConditionRulesAttributesInner) GetTertiaryValue() string`

GetTertiaryValue returns the TertiaryValue field if non-nil, zero value otherwise.

### GetTertiaryValueOk

`func (o *AdvancedSearchConditionRulesAttributesInner) GetTertiaryValueOk() (*string, bool)`

GetTertiaryValueOk returns a tuple with the TertiaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTertiaryValue

`func (o *AdvancedSearchConditionRulesAttributesInner) SetTertiaryValue(v string)`

SetTertiaryValue sets TertiaryValue field to given value.

### HasTertiaryValue

`func (o *AdvancedSearchConditionRulesAttributesInner) HasTertiaryValue() bool`

HasTertiaryValue returns a boolean if a field has been set.


