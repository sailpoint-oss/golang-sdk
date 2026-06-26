---
id: nerm-profile-attribute-rule-id
title: ProfileAttributeRuleId
pagination_label: ProfileAttributeRuleId
sidebar_label: ProfileAttributeRuleId
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileAttributeRuleId', 'NERMProfileAttributeRuleId'] 
slug: /tools/sdk/go/nerm/models/profile-attribute-rule-id
tags: ['SDK', 'Software Development Kit', 'ProfileAttributeRuleId', 'NERMProfileAttributeRuleId']
---

# ProfileAttributeRuleId

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

### NewProfileAttributeRuleId

`func NewProfileAttributeRuleId(type_ string, conditionObjectType string, conditionObjectId string, comparisonOperator string, value string, ) *ProfileAttributeRuleId`

NewProfileAttributeRuleId instantiates a new ProfileAttributeRuleId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileAttributeRuleIdWithDefaults

`func NewProfileAttributeRuleIdWithDefaults() *ProfileAttributeRuleId`

NewProfileAttributeRuleIdWithDefaults instantiates a new ProfileAttributeRuleId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileAttributeRuleId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileAttributeRuleId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileAttributeRuleId) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileAttributeRuleId) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *ProfileAttributeRuleId) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfileAttributeRuleId) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfileAttributeRuleId) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfileAttributeRuleId) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *ProfileAttributeRuleId) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileAttributeRuleId) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileAttributeRuleId) SetType(v string)`

SetType sets Type field to given value.


### GetConditionObjectType

`func (o *ProfileAttributeRuleId) GetConditionObjectType() string`

GetConditionObjectType returns the ConditionObjectType field if non-nil, zero value otherwise.

### GetConditionObjectTypeOk

`func (o *ProfileAttributeRuleId) GetConditionObjectTypeOk() (*string, bool)`

GetConditionObjectTypeOk returns a tuple with the ConditionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectType

`func (o *ProfileAttributeRuleId) SetConditionObjectType(v string)`

SetConditionObjectType sets ConditionObjectType field to given value.


### GetConditionObjectId

`func (o *ProfileAttributeRuleId) GetConditionObjectId() string`

GetConditionObjectId returns the ConditionObjectId field if non-nil, zero value otherwise.

### GetConditionObjectIdOk

`func (o *ProfileAttributeRuleId) GetConditionObjectIdOk() (*string, bool)`

GetConditionObjectIdOk returns a tuple with the ConditionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectId

`func (o *ProfileAttributeRuleId) SetConditionObjectId(v string)`

SetConditionObjectId sets ConditionObjectId field to given value.


### GetComparisonOperator

`func (o *ProfileAttributeRuleId) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *ProfileAttributeRuleId) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *ProfileAttributeRuleId) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.


### GetValue

`func (o *ProfileAttributeRuleId) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProfileAttributeRuleId) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProfileAttributeRuleId) SetValue(v string)`

SetValue sets Value field to given value.



