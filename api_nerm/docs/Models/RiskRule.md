---
id: nerm-risk-rule
title: RiskRule
pagination_label: RiskRule
sidebar_label: RiskRule
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RiskRule', 'NERMRiskRule'] 
slug: /tools/sdk/go/nerm/models/risk-rule
tags: ['SDK', 'Software Development Kit', 'RiskRule', 'NERMRiskRule']
---

# RiskRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**ComparisonOperator** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 
**SecondaryValue** | **string** |  | 

## Methods

### NewRiskRule

`func NewRiskRule(type_ string, value string, secondaryValue string, ) *RiskRule`

NewRiskRule instantiates a new RiskRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRuleWithDefaults

`func NewRiskRuleWithDefaults() *RiskRule`

NewRiskRuleWithDefaults instantiates a new RiskRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *RiskRule) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RiskRule) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RiskRule) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *RiskRule) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *RiskRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskRule) SetType(v string)`

SetType sets Type field to given value.


### GetComparisonOperator

`func (o *RiskRule) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *RiskRule) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *RiskRule) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.

### HasComparisonOperator

`func (o *RiskRule) HasComparisonOperator() bool`

HasComparisonOperator returns a boolean if a field has been set.

### GetValue

`func (o *RiskRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskRule) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecondaryValue

`func (o *RiskRule) GetSecondaryValue() string`

GetSecondaryValue returns the SecondaryValue field if non-nil, zero value otherwise.

### GetSecondaryValueOk

`func (o *RiskRule) GetSecondaryValueOk() (*string, bool)`

GetSecondaryValueOk returns a tuple with the SecondaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryValue

`func (o *RiskRule) SetSecondaryValue(v string)`

SetSecondaryValue sets SecondaryValue field to given value.



