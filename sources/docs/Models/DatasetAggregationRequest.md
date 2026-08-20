---
id: v1-dataset-aggregation-request
title: DatasetAggregationRequest
pagination_label: DatasetAggregationRequest
sidebar_label: DatasetAggregationRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DatasetAggregationRequest', 'V1DatasetAggregationRequest'] 
slug: /tools/sdk/go/sources/models/dataset-aggregation-request
tags: ['SDK', 'Software Development Kit', 'DatasetAggregationRequest', 'V1DatasetAggregationRequest']
---

# DatasetAggregationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Connector-specific aggregation configuration. | [optional] 

## Methods

### NewDatasetAggregationRequest

`func NewDatasetAggregationRequest() *DatasetAggregationRequest`

NewDatasetAggregationRequest instantiates a new DatasetAggregationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetAggregationRequestWithDefaults

`func NewDatasetAggregationRequestWithDefaults() *DatasetAggregationRequest`

NewDatasetAggregationRequestWithDefaults instantiates a new DatasetAggregationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *DatasetAggregationRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DatasetAggregationRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DatasetAggregationRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DatasetAggregationRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


