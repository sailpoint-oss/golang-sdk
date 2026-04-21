---
id: nerm-risk-score
title: RiskScore
pagination_label: RiskScore
sidebar_label: RiskScore
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RiskScore', 'NERMRiskScore'] 
slug: /tools/sdk/go/nerm/models/risk-score
tags: ['SDK', 'Software Development Kit', 'RiskScore', 'NERMRiskScore']
---

# RiskScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**OverallScore** | Pointer to **float32** |  | [optional] 
**OverallRiskLevelId** | Pointer to **string** |  | [optional] 
**ImpactScore** | Pointer to **float32** |  | [optional] 
**ImpactRiskLevelId** | Pointer to **string** |  | [optional] 
**ProbabilityScore** | Pointer to **float32** |  | [optional] 
**ProbabilityRiskLevelId** | Pointer to **string** |  | [optional] 

## Methods

### NewRiskScore

`func NewRiskScore() *RiskScore`

NewRiskScore instantiates a new RiskScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskScoreWithDefaults

`func NewRiskScoreWithDefaults() *RiskScore`

NewRiskScoreWithDefaults instantiates a new RiskScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskScore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskScore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskScore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskScore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *RiskScore) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RiskScore) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RiskScore) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *RiskScore) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetObjectId

`func (o *RiskScore) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RiskScore) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RiskScore) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *RiskScore) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *RiskScore) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RiskScore) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RiskScore) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *RiskScore) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOverallScore

`func (o *RiskScore) GetOverallScore() float32`

GetOverallScore returns the OverallScore field if non-nil, zero value otherwise.

### GetOverallScoreOk

`func (o *RiskScore) GetOverallScoreOk() (*float32, bool)`

GetOverallScoreOk returns a tuple with the OverallScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallScore

`func (o *RiskScore) SetOverallScore(v float32)`

SetOverallScore sets OverallScore field to given value.

### HasOverallScore

`func (o *RiskScore) HasOverallScore() bool`

HasOverallScore returns a boolean if a field has been set.

### GetOverallRiskLevelId

`func (o *RiskScore) GetOverallRiskLevelId() string`

GetOverallRiskLevelId returns the OverallRiskLevelId field if non-nil, zero value otherwise.

### GetOverallRiskLevelIdOk

`func (o *RiskScore) GetOverallRiskLevelIdOk() (*string, bool)`

GetOverallRiskLevelIdOk returns a tuple with the OverallRiskLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallRiskLevelId

`func (o *RiskScore) SetOverallRiskLevelId(v string)`

SetOverallRiskLevelId sets OverallRiskLevelId field to given value.

### HasOverallRiskLevelId

`func (o *RiskScore) HasOverallRiskLevelId() bool`

HasOverallRiskLevelId returns a boolean if a field has been set.

### GetImpactScore

`func (o *RiskScore) GetImpactScore() float32`

GetImpactScore returns the ImpactScore field if non-nil, zero value otherwise.

### GetImpactScoreOk

`func (o *RiskScore) GetImpactScoreOk() (*float32, bool)`

GetImpactScoreOk returns a tuple with the ImpactScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactScore

`func (o *RiskScore) SetImpactScore(v float32)`

SetImpactScore sets ImpactScore field to given value.

### HasImpactScore

`func (o *RiskScore) HasImpactScore() bool`

HasImpactScore returns a boolean if a field has been set.

### GetImpactRiskLevelId

`func (o *RiskScore) GetImpactRiskLevelId() string`

GetImpactRiskLevelId returns the ImpactRiskLevelId field if non-nil, zero value otherwise.

### GetImpactRiskLevelIdOk

`func (o *RiskScore) GetImpactRiskLevelIdOk() (*string, bool)`

GetImpactRiskLevelIdOk returns a tuple with the ImpactRiskLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactRiskLevelId

`func (o *RiskScore) SetImpactRiskLevelId(v string)`

SetImpactRiskLevelId sets ImpactRiskLevelId field to given value.

### HasImpactRiskLevelId

`func (o *RiskScore) HasImpactRiskLevelId() bool`

HasImpactRiskLevelId returns a boolean if a field has been set.

### GetProbabilityScore

`func (o *RiskScore) GetProbabilityScore() float32`

GetProbabilityScore returns the ProbabilityScore field if non-nil, zero value otherwise.

### GetProbabilityScoreOk

`func (o *RiskScore) GetProbabilityScoreOk() (*float32, bool)`

GetProbabilityScoreOk returns a tuple with the ProbabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbabilityScore

`func (o *RiskScore) SetProbabilityScore(v float32)`

SetProbabilityScore sets ProbabilityScore field to given value.

### HasProbabilityScore

`func (o *RiskScore) HasProbabilityScore() bool`

HasProbabilityScore returns a boolean if a field has been set.

### GetProbabilityRiskLevelId

`func (o *RiskScore) GetProbabilityRiskLevelId() string`

GetProbabilityRiskLevelId returns the ProbabilityRiskLevelId field if non-nil, zero value otherwise.

### GetProbabilityRiskLevelIdOk

`func (o *RiskScore) GetProbabilityRiskLevelIdOk() (*string, bool)`

GetProbabilityRiskLevelIdOk returns a tuple with the ProbabilityRiskLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbabilityRiskLevelId

`func (o *RiskScore) SetProbabilityRiskLevelId(v string)`

SetProbabilityRiskLevelId sets ProbabilityRiskLevelId field to given value.

### HasProbabilityRiskLevelId

`func (o *RiskScore) HasProbabilityRiskLevelId() bool`

HasProbabilityRiskLevelId returns a boolean if a field has been set.


