---
id: nerm-advanced-search
title: AdvancedSearch
pagination_label: AdvancedSearch
sidebar_label: AdvancedSearch
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AdvancedSearch', 'NERMAdvancedSearch'] 
slug: /tools/sdk/go/nerm/models/advanced-search
tags: ['SDK', 'Software Development Kit', 'AdvancedSearch', 'NERMAdvancedSearch']
---

# AdvancedSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**Label** | Pointer to **string** |  | [optional] 
**ConditionRulesAttributes** | Pointer to [**[]AdvancedSearchConditionRulesAttributesInner**](advanced-search-condition-rules-attributes-inner) |  | [optional] 

## Methods

### NewAdvancedSearch

`func NewAdvancedSearch() *AdvancedSearch`

NewAdvancedSearch instantiates a new AdvancedSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedSearchWithDefaults

`func NewAdvancedSearchWithDefaults() *AdvancedSearch`

NewAdvancedSearchWithDefaults instantiates a new AdvancedSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdvancedSearch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedSearch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedSearch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedSearch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *AdvancedSearch) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AdvancedSearch) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AdvancedSearch) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AdvancedSearch) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLabel

`func (o *AdvancedSearch) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AdvancedSearch) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AdvancedSearch) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AdvancedSearch) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetConditionRulesAttributes

`func (o *AdvancedSearch) GetConditionRulesAttributes() []AdvancedSearchConditionRulesAttributesInner`

GetConditionRulesAttributes returns the ConditionRulesAttributes field if non-nil, zero value otherwise.

### GetConditionRulesAttributesOk

`func (o *AdvancedSearch) GetConditionRulesAttributesOk() (*[]AdvancedSearchConditionRulesAttributesInner, bool)`

GetConditionRulesAttributesOk returns a tuple with the ConditionRulesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionRulesAttributes

`func (o *AdvancedSearch) SetConditionRulesAttributes(v []AdvancedSearchConditionRulesAttributesInner)`

SetConditionRulesAttributes sets ConditionRulesAttributes field to given value.

### HasConditionRulesAttributes

`func (o *AdvancedSearch) HasConditionRulesAttributes() bool`

HasConditionRulesAttributes returns a boolean if a field has been set.


