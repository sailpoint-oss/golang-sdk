---
id: nerm-get-risk-score200-response
title: GetRiskScore200Response
pagination_label: GetRiskScore200Response
sidebar_label: GetRiskScore200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetRiskScore200Response', 'NERMGetRiskScore200Response'] 
slug: /tools/sdk/go/nerm/models/get-risk-score200-response
tags: ['SDK', 'Software Development Kit', 'GetRiskScore200Response', 'NERMGetRiskScore200Response']
---

# GetRiskScore200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskScore** | Pointer to [**RiskScore**](risk-score) |  | [optional] 

## Methods

### NewGetRiskScore200Response

`func NewGetRiskScore200Response() *GetRiskScore200Response`

NewGetRiskScore200Response instantiates a new GetRiskScore200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRiskScore200ResponseWithDefaults

`func NewGetRiskScore200ResponseWithDefaults() *GetRiskScore200Response`

NewGetRiskScore200ResponseWithDefaults instantiates a new GetRiskScore200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskScore

`func (o *GetRiskScore200Response) GetRiskScore() RiskScore`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *GetRiskScore200Response) GetRiskScoreOk() (*RiskScore, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *GetRiskScore200Response) SetRiskScore(v RiskScore)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *GetRiskScore200Response) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.


