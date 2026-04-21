---
id: nerm-profile-attribute-rule-string
title: ProfileAttributeRuleString
pagination_label: ProfileAttributeRuleString
sidebar_label: ProfileAttributeRuleString
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileAttributeRuleString', 'NERMProfileAttributeRuleString'] 
slug: /tools/sdk/go/nerm/models/profile-attribute-rule-string
tags: ['SDK', 'Software Development Kit', 'ProfileAttributeRuleString', 'NERMProfileAttributeRuleString']
---

# ProfileAttributeRuleString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**ConditionObjectType** | **string** |  | 
**ConditionObjectId** | **string** |  | 
**ComparisonOperator** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewProfileAttributeRuleString

`func NewProfileAttributeRuleString(type_ string, conditionObjectType string, conditionObjectId string, comparisonOperator string, value string, ) *ProfileAttributeRuleString`

NewProfileAttributeRuleString instantiates a new ProfileAttributeRuleString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileAttributeRuleStringWithDefaults

`func NewProfileAttributeRuleStringWithDefaults() *ProfileAttributeRuleString`

NewProfileAttributeRuleStringWithDefaults instantiates a new ProfileAttributeRuleString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileAttributeRuleString) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileAttributeRuleString) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileAttributeRuleString) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileAttributeRuleString) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *ProfileAttributeRuleString) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfileAttributeRuleString) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfileAttributeRuleString) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfileAttributeRuleString) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *ProfileAttributeRuleString) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileAttributeRuleString) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileAttributeRuleString) SetType(v string)`

SetType sets Type field to given value.


### GetConditionObjectType

`func (o *ProfileAttributeRuleString) GetConditionObjectType() string`

GetConditionObjectType returns the ConditionObjectType field if non-nil, zero value otherwise.

### GetConditionObjectTypeOk

`func (o *ProfileAttributeRuleString) GetConditionObjectTypeOk() (*string, bool)`

GetConditionObjectTypeOk returns a tuple with the ConditionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectType

`func (o *ProfileAttributeRuleString) SetConditionObjectType(v string)`

SetConditionObjectType sets ConditionObjectType field to given value.


### GetConditionObjectId

`func (o *ProfileAttributeRuleString) GetConditionObjectId() string`

GetConditionObjectId returns the ConditionObjectId field if non-nil, zero value otherwise.

### GetConditionObjectIdOk

`func (o *ProfileAttributeRuleString) GetConditionObjectIdOk() (*string, bool)`

GetConditionObjectIdOk returns a tuple with the ConditionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectId

`func (o *ProfileAttributeRuleString) SetConditionObjectId(v string)`

SetConditionObjectId sets ConditionObjectId field to given value.


### GetComparisonOperator

`func (o *ProfileAttributeRuleString) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *ProfileAttributeRuleString) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *ProfileAttributeRuleString) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.


### GetValue

`func (o *ProfileAttributeRuleString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProfileAttributeRuleString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProfileAttributeRuleString) SetValue(v string)`

SetValue sets Value field to given value.



