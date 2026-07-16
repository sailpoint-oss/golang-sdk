---
id: v1-correlationrule
title: Correlationrule
pagination_label: Correlationrule
sidebar_label: Correlationrule
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Correlationrule', 'V1Correlationrule'] 
slug: /tools/sdk/go/machineidentities/models/correlationrule
tags: ['SDK', 'Software Development Kit', 'Correlationrule', 'V1Correlationrule']
---

# Correlationrule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Omit for new rules (server mints a UUID). Send only when updating a rule that already exists on this config (merge on PATCH). Unknown ids are rejected. | [optional] 
**Priority** | **int32** | The evaluation priority of the rule. Lower values are evaluated first. | 
**DefaultRule** | **bool** | Whether this rule is the default rule for the config. | 
**RuleType** | **string** | The rule subject type. When either ruleType or ruleAction.type is GOVERNANCE_GROUP, both must be; ruleType GOVERNANCE_GROUP is allowed only when the parent config type is OWNER_SECONDARY. | 
**RuleAction** | [**Correlationruleaction**](correlationruleaction) |  | 
**ConditionExpressions** | [**[]Correlationcondition**](correlationcondition) | The conditions that must match for this rule to apply. | 

## Methods

### NewCorrelationrule

`func NewCorrelationrule(priority int32, defaultRule bool, ruleType string, ruleAction Correlationruleaction, conditionExpressions []Correlationcondition, ) *Correlationrule`

NewCorrelationrule instantiates a new Correlationrule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelationruleWithDefaults

`func NewCorrelationruleWithDefaults() *Correlationrule`

NewCorrelationruleWithDefaults instantiates a new Correlationrule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Correlationrule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Correlationrule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Correlationrule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Correlationrule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *Correlationrule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Correlationrule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Correlationrule) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetDefaultRule

`func (o *Correlationrule) GetDefaultRule() bool`

GetDefaultRule returns the DefaultRule field if non-nil, zero value otherwise.

### GetDefaultRuleOk

`func (o *Correlationrule) GetDefaultRuleOk() (*bool, bool)`

GetDefaultRuleOk returns a tuple with the DefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRule

`func (o *Correlationrule) SetDefaultRule(v bool)`

SetDefaultRule sets DefaultRule field to given value.


### GetRuleType

`func (o *Correlationrule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *Correlationrule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *Correlationrule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetRuleAction

`func (o *Correlationrule) GetRuleAction() Correlationruleaction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *Correlationrule) GetRuleActionOk() (*Correlationruleaction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *Correlationrule) SetRuleAction(v Correlationruleaction)`

SetRuleAction sets RuleAction field to given value.


### GetConditionExpressions

`func (o *Correlationrule) GetConditionExpressions() []Correlationcondition`

GetConditionExpressions returns the ConditionExpressions field if non-nil, zero value otherwise.

### GetConditionExpressionsOk

`func (o *Correlationrule) GetConditionExpressionsOk() (*[]Correlationcondition, bool)`

GetConditionExpressionsOk returns a tuple with the ConditionExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionExpressions

`func (o *Correlationrule) SetConditionExpressions(v []Correlationcondition)`

SetConditionExpressions sets ConditionExpressions field to given value.



