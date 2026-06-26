---
id: nerm-risk-rule1
title: RiskRule1
pagination_label: RiskRule1
sidebar_label: RiskRule1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RiskRule1', 'NERMRiskRule1'] 
slug: /tools/sdk/go/nerm/models/risk-rule1
tags: ['SDK', 'Software Development Kit', 'RiskRule1', 'NERMRiskRule1']
---

# RiskRule1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ComparisonOperator** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 
**SecondaryValue** | **string** |  | 

## Methods

### NewRiskRule1

`func NewRiskRule1(type_ string, value string, secondaryValue string, ) *RiskRule1`

NewRiskRule1 instantiates a new RiskRule1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRule1WithDefaults

`func NewRiskRule1WithDefaults() *RiskRule1`

NewRiskRule1WithDefaults instantiates a new RiskRule1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RiskRule1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskRule1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskRule1) SetType(v string)`

SetType sets Type field to given value.


### GetComparisonOperator

`func (o *RiskRule1) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *RiskRule1) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *RiskRule1) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.

### HasComparisonOperator

`func (o *RiskRule1) HasComparisonOperator() bool`

HasComparisonOperator returns a boolean if a field has been set.

### GetValue

`func (o *RiskRule1) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskRule1) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskRule1) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecondaryValue

`func (o *RiskRule1) GetSecondaryValue() string`

GetSecondaryValue returns the SecondaryValue field if non-nil, zero value otherwise.

### GetSecondaryValueOk

`func (o *RiskRule1) GetSecondaryValueOk() (*string, bool)`

GetSecondaryValueOk returns a tuple with the SecondaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryValue

`func (o *RiskRule1) SetSecondaryValue(v string)`

SetSecondaryValue sets SecondaryValue field to given value.



