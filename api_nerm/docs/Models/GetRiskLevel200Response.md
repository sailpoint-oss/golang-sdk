---
id: nerm-get-risk-level200-response
title: GetRiskLevel200Response
pagination_label: GetRiskLevel200Response
sidebar_label: GetRiskLevel200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetRiskLevel200Response', 'NERMGetRiskLevel200Response'] 
slug: /tools/sdk/go/nerm/models/get-risk-level200-response
tags: ['SDK', 'Software Development Kit', 'GetRiskLevel200Response', 'NERMGetRiskLevel200Response']
---

# GetRiskLevel200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskLevel** | Pointer to [**RiskLevel**](risk-level) |  | [optional] 

## Methods

### NewGetRiskLevel200Response

`func NewGetRiskLevel200Response() *GetRiskLevel200Response`

NewGetRiskLevel200Response instantiates a new GetRiskLevel200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRiskLevel200ResponseWithDefaults

`func NewGetRiskLevel200ResponseWithDefaults() *GetRiskLevel200Response`

NewGetRiskLevel200ResponseWithDefaults instantiates a new GetRiskLevel200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskLevel

`func (o *GetRiskLevel200Response) GetRiskLevel() RiskLevel`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *GetRiskLevel200Response) GetRiskLevelOk() (*RiskLevel, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *GetRiskLevel200Response) SetRiskLevel(v RiskLevel)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *GetRiskLevel200Response) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.


