---
id: nerm-profile-attribute-rule-date
title: ProfileAttributeRuleDate
pagination_label: ProfileAttributeRuleDate
sidebar_label: ProfileAttributeRuleDate
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileAttributeRuleDate', 'NERMProfileAttributeRuleDate'] 
slug: /tools/sdk/go/nerm/models/profile-attribute-rule-date
tags: ['SDK', 'Software Development Kit', 'ProfileAttributeRuleDate', 'NERMProfileAttributeRuleDate']
---

# ProfileAttributeRuleDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**ConditionObjectType** | **string** |  | 
**ConditionObjectId** | Pointer to **string** |  | [optional] 
**SecondaryAttributeType** | Pointer to **string** |  | [optional] 
**SecondaryAttributeId** | Pointer to **string** |  | [optional] 
**ComparisonOperator** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 
**SecondaryValue** | Pointer to **string** |  | [optional] 
**TertiaryValue** | Pointer to **string** |  | [optional] 

## Methods

### NewProfileAttributeRuleDate

`func NewProfileAttributeRuleDate(type_ string, conditionObjectType string, value string, ) *ProfileAttributeRuleDate`

NewProfileAttributeRuleDate instantiates a new ProfileAttributeRuleDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileAttributeRuleDateWithDefaults

`func NewProfileAttributeRuleDateWithDefaults() *ProfileAttributeRuleDate`

NewProfileAttributeRuleDateWithDefaults instantiates a new ProfileAttributeRuleDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileAttributeRuleDate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileAttributeRuleDate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileAttributeRuleDate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileAttributeRuleDate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *ProfileAttributeRuleDate) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfileAttributeRuleDate) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfileAttributeRuleDate) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfileAttributeRuleDate) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *ProfileAttributeRuleDate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileAttributeRuleDate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileAttributeRuleDate) SetType(v string)`

SetType sets Type field to given value.


### GetConditionObjectType

`func (o *ProfileAttributeRuleDate) GetConditionObjectType() string`

GetConditionObjectType returns the ConditionObjectType field if non-nil, zero value otherwise.

### GetConditionObjectTypeOk

`func (o *ProfileAttributeRuleDate) GetConditionObjectTypeOk() (*string, bool)`

GetConditionObjectTypeOk returns a tuple with the ConditionObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectType

`func (o *ProfileAttributeRuleDate) SetConditionObjectType(v string)`

SetConditionObjectType sets ConditionObjectType field to given value.


### GetConditionObjectId

`func (o *ProfileAttributeRuleDate) GetConditionObjectId() string`

GetConditionObjectId returns the ConditionObjectId field if non-nil, zero value otherwise.

### GetConditionObjectIdOk

`func (o *ProfileAttributeRuleDate) GetConditionObjectIdOk() (*string, bool)`

GetConditionObjectIdOk returns a tuple with the ConditionObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionObjectId

`func (o *ProfileAttributeRuleDate) SetConditionObjectId(v string)`

SetConditionObjectId sets ConditionObjectId field to given value.

### HasConditionObjectId

`func (o *ProfileAttributeRuleDate) HasConditionObjectId() bool`

HasConditionObjectId returns a boolean if a field has been set.

### GetSecondaryAttributeType

`func (o *ProfileAttributeRuleDate) GetSecondaryAttributeType() string`

GetSecondaryAttributeType returns the SecondaryAttributeType field if non-nil, zero value otherwise.

### GetSecondaryAttributeTypeOk

`func (o *ProfileAttributeRuleDate) GetSecondaryAttributeTypeOk() (*string, bool)`

GetSecondaryAttributeTypeOk returns a tuple with the SecondaryAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAttributeType

`func (o *ProfileAttributeRuleDate) SetSecondaryAttributeType(v string)`

SetSecondaryAttributeType sets SecondaryAttributeType field to given value.

### HasSecondaryAttributeType

`func (o *ProfileAttributeRuleDate) HasSecondaryAttributeType() bool`

HasSecondaryAttributeType returns a boolean if a field has been set.

### GetSecondaryAttributeId

`func (o *ProfileAttributeRuleDate) GetSecondaryAttributeId() string`

GetSecondaryAttributeId returns the SecondaryAttributeId field if non-nil, zero value otherwise.

### GetSecondaryAttributeIdOk

`func (o *ProfileAttributeRuleDate) GetSecondaryAttributeIdOk() (*string, bool)`

GetSecondaryAttributeIdOk returns a tuple with the SecondaryAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAttributeId

`func (o *ProfileAttributeRuleDate) SetSecondaryAttributeId(v string)`

SetSecondaryAttributeId sets SecondaryAttributeId field to given value.

### HasSecondaryAttributeId

`func (o *ProfileAttributeRuleDate) HasSecondaryAttributeId() bool`

HasSecondaryAttributeId returns a boolean if a field has been set.

### GetComparisonOperator

`func (o *ProfileAttributeRuleDate) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *ProfileAttributeRuleDate) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *ProfileAttributeRuleDate) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.

### HasComparisonOperator

`func (o *ProfileAttributeRuleDate) HasComparisonOperator() bool`

HasComparisonOperator returns a boolean if a field has been set.

### GetValue

`func (o *ProfileAttributeRuleDate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProfileAttributeRuleDate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProfileAttributeRuleDate) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecondaryValue

`func (o *ProfileAttributeRuleDate) GetSecondaryValue() string`

GetSecondaryValue returns the SecondaryValue field if non-nil, zero value otherwise.

### GetSecondaryValueOk

`func (o *ProfileAttributeRuleDate) GetSecondaryValueOk() (*string, bool)`

GetSecondaryValueOk returns a tuple with the SecondaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryValue

`func (o *ProfileAttributeRuleDate) SetSecondaryValue(v string)`

SetSecondaryValue sets SecondaryValue field to given value.

### HasSecondaryValue

`func (o *ProfileAttributeRuleDate) HasSecondaryValue() bool`

HasSecondaryValue returns a boolean if a field has been set.

### GetTertiaryValue

`func (o *ProfileAttributeRuleDate) GetTertiaryValue() string`

GetTertiaryValue returns the TertiaryValue field if non-nil, zero value otherwise.

### GetTertiaryValueOk

`func (o *ProfileAttributeRuleDate) GetTertiaryValueOk() (*string, bool)`

GetTertiaryValueOk returns a tuple with the TertiaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTertiaryValue

`func (o *ProfileAttributeRuleDate) SetTertiaryValue(v string)`

SetTertiaryValue sets TertiaryValue field to given value.

### HasTertiaryValue

`func (o *ProfileAttributeRuleDate) HasTertiaryValue() bool`

HasTertiaryValue returns a boolean if a field has been set.


