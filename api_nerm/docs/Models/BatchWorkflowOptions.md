---
id: nerm-batch-workflow-options
title: BatchWorkflowOptions
pagination_label: BatchWorkflowOptions
sidebar_label: BatchWorkflowOptions
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'BatchWorkflowOptions', 'NERMBatchWorkflowOptions'] 
slug: /tools/sdk/go/nerm/models/batch-workflow-options
tags: ['SDK', 'Software Development Kit', 'BatchWorkflowOptions', 'NERMBatchWorkflowOptions']
---

# BatchWorkflowOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllProfiles** | Pointer to **string** | Used for batch workflows. | [optional] 
**MultipleRequests** | Pointer to **string** | Used for batch workflows for if multiple requests. | [optional] 

## Methods

### NewBatchWorkflowOptions

`func NewBatchWorkflowOptions() *BatchWorkflowOptions`

NewBatchWorkflowOptions instantiates a new BatchWorkflowOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWorkflowOptionsWithDefaults

`func NewBatchWorkflowOptionsWithDefaults() *BatchWorkflowOptions`

NewBatchWorkflowOptionsWithDefaults instantiates a new BatchWorkflowOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllProfiles

`func (o *BatchWorkflowOptions) GetAllProfiles() string`

GetAllProfiles returns the AllProfiles field if non-nil, zero value otherwise.

### GetAllProfilesOk

`func (o *BatchWorkflowOptions) GetAllProfilesOk() (*string, bool)`

GetAllProfilesOk returns a tuple with the AllProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllProfiles

`func (o *BatchWorkflowOptions) SetAllProfiles(v string)`

SetAllProfiles sets AllProfiles field to given value.

### HasAllProfiles

`func (o *BatchWorkflowOptions) HasAllProfiles() bool`

HasAllProfiles returns a boolean if a field has been set.

### GetMultipleRequests

`func (o *BatchWorkflowOptions) GetMultipleRequests() string`

GetMultipleRequests returns the MultipleRequests field if non-nil, zero value otherwise.

### GetMultipleRequestsOk

`func (o *BatchWorkflowOptions) GetMultipleRequestsOk() (*string, bool)`

GetMultipleRequestsOk returns a tuple with the MultipleRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleRequests

`func (o *BatchWorkflowOptions) SetMultipleRequests(v string)`

SetMultipleRequests sets MultipleRequests field to given value.

### HasMultipleRequests

`func (o *BatchWorkflowOptions) HasMultipleRequests() bool`

HasMultipleRequests returns a boolean if a field has been set.


