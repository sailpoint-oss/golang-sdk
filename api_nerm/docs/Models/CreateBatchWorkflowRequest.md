---
id: nerm-create-batch-workflow-request
title: CreateBatchWorkflowRequest
pagination_label: CreateBatchWorkflowRequest
sidebar_label: CreateBatchWorkflowRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateBatchWorkflowRequest', 'NERMCreateBatchWorkflowRequest'] 
slug: /tools/sdk/go/nerm/models/create-batch-workflow-request
tags: ['SDK', 'Software Development Kit', 'CreateBatchWorkflowRequest', 'NERMCreateBatchWorkflowRequest']
---

# CreateBatchWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | Pointer to [**BatchWorkflow**](batch-workflow) |  | [optional] 

## Methods

### NewCreateBatchWorkflowRequest

`func NewCreateBatchWorkflowRequest() *CreateBatchWorkflowRequest`

NewCreateBatchWorkflowRequest instantiates a new CreateBatchWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchWorkflowRequestWithDefaults

`func NewCreateBatchWorkflowRequestWithDefaults() *CreateBatchWorkflowRequest`

NewCreateBatchWorkflowRequestWithDefaults instantiates a new CreateBatchWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *CreateBatchWorkflowRequest) GetWorkflow() BatchWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CreateBatchWorkflowRequest) GetWorkflowOk() (*BatchWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CreateBatchWorkflowRequest) SetWorkflow(v BatchWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CreateBatchWorkflowRequest) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


