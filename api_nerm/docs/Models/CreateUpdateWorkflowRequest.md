---
id: nerm-create-update-workflow-request
title: CreateUpdateWorkflowRequest
pagination_label: CreateUpdateWorkflowRequest
sidebar_label: CreateUpdateWorkflowRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateUpdateWorkflowRequest', 'NERMCreateUpdateWorkflowRequest'] 
slug: /tools/sdk/go/nerm/models/create-update-workflow-request
tags: ['SDK', 'Software Development Kit', 'CreateUpdateWorkflowRequest', 'NERMCreateUpdateWorkflowRequest']
---

# CreateUpdateWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | Pointer to [**UpdateWorkflow**](update-workflow) |  | [optional] 

## Methods

### NewCreateUpdateWorkflowRequest

`func NewCreateUpdateWorkflowRequest() *CreateUpdateWorkflowRequest`

NewCreateUpdateWorkflowRequest instantiates a new CreateUpdateWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateWorkflowRequestWithDefaults

`func NewCreateUpdateWorkflowRequestWithDefaults() *CreateUpdateWorkflowRequest`

NewCreateUpdateWorkflowRequestWithDefaults instantiates a new CreateUpdateWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *CreateUpdateWorkflowRequest) GetWorkflow() UpdateWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CreateUpdateWorkflowRequest) GetWorkflowOk() (*UpdateWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CreateUpdateWorkflowRequest) SetWorkflow(v UpdateWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CreateUpdateWorkflowRequest) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


