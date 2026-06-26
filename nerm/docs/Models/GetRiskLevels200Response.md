---
id: nerm-get-risk-levels200-response
title: GetRiskLevels200Response
pagination_label: GetRiskLevels200Response
sidebar_label: GetRiskLevels200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetRiskLevels200Response', 'NERMGetRiskLevels200Response'] 
slug: /tools/sdk/go/nerm/models/get-risk-levels200-response
tags: ['SDK', 'Software Development Kit', 'GetRiskLevels200Response', 'NERMGetRiskLevels200Response']
---

# GetRiskLevels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskLevels** | Pointer to [**[]RiskLevel**](risk-level) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetRiskLevels200Response

`func NewGetRiskLevels200Response() *GetRiskLevels200Response`

NewGetRiskLevels200Response instantiates a new GetRiskLevels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRiskLevels200ResponseWithDefaults

`func NewGetRiskLevels200ResponseWithDefaults() *GetRiskLevels200Response`

NewGetRiskLevels200ResponseWithDefaults instantiates a new GetRiskLevels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskLevels

`func (o *GetRiskLevels200Response) GetRiskLevels() []RiskLevel`

GetRiskLevels returns the RiskLevels field if non-nil, zero value otherwise.

### GetRiskLevelsOk

`func (o *GetRiskLevels200Response) GetRiskLevelsOk() (*[]RiskLevel, bool)`

GetRiskLevelsOk returns a tuple with the RiskLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevels

`func (o *GetRiskLevels200Response) SetRiskLevels(v []RiskLevel)`

SetRiskLevels sets RiskLevels field to given value.

### HasRiskLevels

`func (o *GetRiskLevels200Response) HasRiskLevels() bool`

HasRiskLevels returns a boolean if a field has been set.

### GetMetadata

`func (o *GetRiskLevels200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetRiskLevels200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetRiskLevels200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetRiskLevels200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


