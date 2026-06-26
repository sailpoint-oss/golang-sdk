---
id: nerm-profile-status-rule
title: ProfileStatusRule
pagination_label: ProfileStatusRule
sidebar_label: ProfileStatusRule
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileStatusRule', 'NERMProfileStatusRule'] 
slug: /tools/sdk/go/nerm/models/profile-status-rule
tags: ['SDK', 'Software Development Kit', 'ProfileStatusRule', 'NERMProfileStatusRule']
---

# ProfileStatusRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**ComparisonOperator** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewProfileStatusRule

`func NewProfileStatusRule(type_ string, ) *ProfileStatusRule`

NewProfileStatusRule instantiates a new ProfileStatusRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileStatusRuleWithDefaults

`func NewProfileStatusRuleWithDefaults() *ProfileStatusRule`

NewProfileStatusRuleWithDefaults instantiates a new ProfileStatusRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileStatusRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileStatusRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileStatusRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileStatusRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *ProfileStatusRule) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfileStatusRule) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfileStatusRule) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfileStatusRule) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *ProfileStatusRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileStatusRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileStatusRule) SetType(v string)`

SetType sets Type field to given value.


### GetComparisonOperator

`func (o *ProfileStatusRule) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *ProfileStatusRule) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *ProfileStatusRule) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.

### HasComparisonOperator

`func (o *ProfileStatusRule) HasComparisonOperator() bool`

HasComparisonOperator returns a boolean if a field has been set.

### GetValue

`func (o *ProfileStatusRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProfileStatusRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProfileStatusRule) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProfileStatusRule) HasValue() bool`

HasValue returns a boolean if a field has been set.


