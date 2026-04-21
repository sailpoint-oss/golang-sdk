---
id: nerm-profile-type-rule
title: ProfileTypeRule
pagination_label: ProfileTypeRule
sidebar_label: ProfileTypeRule
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileTypeRule', 'NERMProfileTypeRule'] 
slug: /tools/sdk/go/nerm/models/profile-type-rule
tags: ['SDK', 'Software Development Kit', 'ProfileTypeRule', 'NERMProfileTypeRule']
---

# ProfileTypeRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**Type** | **string** |  | 
**ComparisonOperator** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewProfileTypeRule

`func NewProfileTypeRule(type_ string, ) *ProfileTypeRule`

NewProfileTypeRule instantiates a new ProfileTypeRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileTypeRuleWithDefaults

`func NewProfileTypeRuleWithDefaults() *ProfileTypeRule`

NewProfileTypeRuleWithDefaults instantiates a new ProfileTypeRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileTypeRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileTypeRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileTypeRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileTypeRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *ProfileTypeRule) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfileTypeRule) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfileTypeRule) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfileTypeRule) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *ProfileTypeRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileTypeRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileTypeRule) SetType(v string)`

SetType sets Type field to given value.


### GetComparisonOperator

`func (o *ProfileTypeRule) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *ProfileTypeRule) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *ProfileTypeRule) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.

### HasComparisonOperator

`func (o *ProfileTypeRule) HasComparisonOperator() bool`

HasComparisonOperator returns a boolean if a field has been set.

### GetValue

`func (o *ProfileTypeRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProfileTypeRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProfileTypeRule) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProfileTypeRule) HasValue() bool`

HasValue returns a boolean if a field has been set.


