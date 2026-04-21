---
id: nerm-get-risk-scores200-response
title: GetRiskScores200Response
pagination_label: GetRiskScores200Response
sidebar_label: GetRiskScores200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetRiskScores200Response', 'NERMGetRiskScores200Response'] 
slug: /tools/sdk/go/nerm/models/get-risk-scores200-response
tags: ['SDK', 'Software Development Kit', 'GetRiskScores200Response', 'NERMGetRiskScores200Response']
---

# GetRiskScores200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskScores** | Pointer to [**[]RiskScore**](risk-score) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetRiskScores200Response

`func NewGetRiskScores200Response() *GetRiskScores200Response`

NewGetRiskScores200Response instantiates a new GetRiskScores200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRiskScores200ResponseWithDefaults

`func NewGetRiskScores200ResponseWithDefaults() *GetRiskScores200Response`

NewGetRiskScores200ResponseWithDefaults instantiates a new GetRiskScores200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskScores

`func (o *GetRiskScores200Response) GetRiskScores() []RiskScore`

GetRiskScores returns the RiskScores field if non-nil, zero value otherwise.

### GetRiskScoresOk

`func (o *GetRiskScores200Response) GetRiskScoresOk() (*[]RiskScore, bool)`

GetRiskScoresOk returns a tuple with the RiskScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScores

`func (o *GetRiskScores200Response) SetRiskScores(v []RiskScore)`

SetRiskScores sets RiskScores field to given value.

### HasRiskScores

`func (o *GetRiskScores200Response) HasRiskScores() bool`

HasRiskScores returns a boolean if a field has been set.

### GetMetadata

`func (o *GetRiskScores200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetRiskScores200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetRiskScores200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetRiskScores200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


