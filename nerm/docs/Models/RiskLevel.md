---
id: nerm-risk-level
title: RiskLevel
pagination_label: RiskLevel
sidebar_label: RiskLevel
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RiskLevel', 'NERMRiskLevel'] 
slug: /tools/sdk/go/nerm/models/risk-level
tags: ['SDK', 'Software Development Kit', 'RiskLevel', 'NERMRiskLevel']
---

# RiskLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**Label** | Pointer to **string** |  | [optional] 
**Points** | Pointer to **float32** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewRiskLevel

`func NewRiskLevel() *RiskLevel`

NewRiskLevel instantiates a new RiskLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLevelWithDefaults

`func NewRiskLevelWithDefaults() *RiskLevel`

NewRiskLevelWithDefaults instantiates a new RiskLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskLevel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskLevel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskLevel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *RiskLevel) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RiskLevel) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RiskLevel) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *RiskLevel) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLabel

`func (o *RiskLevel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RiskLevel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RiskLevel) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RiskLevel) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPoints

`func (o *RiskLevel) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *RiskLevel) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *RiskLevel) SetPoints(v float32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *RiskLevel) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetOrder

`func (o *RiskLevel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RiskLevel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RiskLevel) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RiskLevel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


