---
id: nerm-automated-workflow-condition-rules-attributes-inner
title: AutomatedWorkflowConditionRulesAttributesInner
pagination_label: AutomatedWorkflowConditionRulesAttributesInner
sidebar_label: AutomatedWorkflowConditionRulesAttributesInner
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AutomatedWorkflowConditionRulesAttributesInner', 'NERMAutomatedWorkflowConditionRulesAttributesInner'] 
slug: /tools/sdk/go/nerm/models/automated-workflow-condition-rules-attributes-inner
tags: ['SDK', 'Software Development Kit', 'AutomatedWorkflowConditionRulesAttributesInner', 'NERMAutomatedWorkflowConditionRulesAttributesInner']
---

# AutomatedWorkflowConditionRulesAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalOperator** | **string** | The type of condition rule. | 
**ComparisonOperator** | **string** | The operator used by the condition rule. | 
**ConditionObjectId** | Pointer to **string** | UUID for object targeted by the condition. | [optional] 
**ConditionObjectType** | **string** | Type of object targeted by the condition. | 
**SecondaryValue** | Pointer to **string** | Used for time_frame comparison. | [optional] 
**TertiaryValue** | Pointer to **string** | Represents the days of a temporal comparison. | [optional] 
**Value** | Pointer to **string** | The value being compared against. | [optional] 
**Order** | Pointer to **int32** | If there are multiple rules against one workflow, this is the order that they are ran in. | [optional] 
**Type** | **string** | The type of condition rule. | 

## Methods

### NewAutomatedWorkflowConditionRulesAttributesInner

`func NewAutomatedWorkflowConditionRulesAttributesInner(logicalOperator string, comparisonOperator string, conditionObjectType string, type_ string, ) *AutomatedWorkflowConditionRulesAttributesInner`

NewAutomatedWorkflowConditionRulesAttributesInner instantiates a new AutomatedWorkflowConditionRulesAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomatedWorkflowConditionRulesAttributesInnerWithDefaults

`func NewAutomatedWorkflowConditionRulesAttributesInnerWithDefaults() *AutomatedWorkflowConditionRulesAttributesInner`

NewAutomatedWorkflowConditionRulesAttributesInnerWithDefaults instantiates a new AutomatedWorkflowConditionRulesAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalOperator

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetLogicalOperator() string`

GetLogicalOperator returns the LogicalOperator field if non-nil, zero value otherwise.

### GetLogicalOperatorOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetLogicalOperatorOk() (*string, bool)`

GetLogicalOperatorOk returns a tuple with the LogicalOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalOperator

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetLogicalOperator(v string)`

SetLogicalOperator sets LogicalOperator field to given value.


### GetComparisonOperator

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.


### GetConditionObjectId

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetConditionObjectId() string`

GetConditionObjectId returns the ConditionObjectId field if non-nil, zero value otherwise.

### GetConditionObjectIdOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetConditionObjectIdOk() (*string, bool)`

GetConditionObjectIdOk returns a tuple with the ConditionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectId

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetConditionObjectId(v string)`

SetConditionObjectId sets ConditionObjectId field to given value.

### HasConditionObjectId

`func (o *AutomatedWorkflowConditionRulesAttributesInner) HasConditionObjectId() bool`

HasConditionObjectId returns a boolean if a field has been set.

### GetConditionObjectType

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetConditionObjectType() string`

GetConditionObjectType returns the ConditionObjectType field if non-nil, zero value otherwise.

### GetConditionObjectTypeOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetConditionObjectTypeOk() (*string, bool)`

GetConditionObjectTypeOk returns a tuple with the ConditionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectType

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetConditionObjectType(v string)`

SetConditionObjectType sets ConditionObjectType field to given value.


### GetSecondaryValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetSecondaryValue() string`

GetSecondaryValue returns the SecondaryValue field if non-nil, zero value otherwise.

### GetSecondaryValueOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetSecondaryValueOk() (*string, bool)`

GetSecondaryValueOk returns a tuple with the SecondaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetSecondaryValue(v string)`

SetSecondaryValue sets SecondaryValue field to given value.

### HasSecondaryValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) HasSecondaryValue() bool`

HasSecondaryValue returns a boolean if a field has been set.

### GetTertiaryValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetTertiaryValue() string`

GetTertiaryValue returns the TertiaryValue field if non-nil, zero value otherwise.

### GetTertiaryValueOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetTertiaryValueOk() (*string, bool)`

GetTertiaryValueOk returns a tuple with the TertiaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTertiaryValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetTertiaryValue(v string)`

SetTertiaryValue sets TertiaryValue field to given value.

### HasTertiaryValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) HasTertiaryValue() bool`

HasTertiaryValue returns a boolean if a field has been set.

### GetValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AutomatedWorkflowConditionRulesAttributesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOrder

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AutomatedWorkflowConditionRulesAttributesInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutomatedWorkflowConditionRulesAttributesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutomatedWorkflowConditionRulesAttributesInner) SetType(v string)`

SetType sets Type field to given value.



