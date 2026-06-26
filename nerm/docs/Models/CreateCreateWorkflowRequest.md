---
id: nerm-create-create-workflow-request
title: CreateCreateWorkflowRequest
pagination_label: CreateCreateWorkflowRequest
sidebar_label: CreateCreateWorkflowRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateCreateWorkflowRequest', 'NERMCreateCreateWorkflowRequest'] 
slug: /tools/sdk/go/nerm/models/create-create-workflow-request
tags: ['SDK', 'Software Development Kit', 'CreateCreateWorkflowRequest', 'NERMCreateCreateWorkflowRequest']
---

# CreateCreateWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | Pointer to [**CreateWorkflow**](create-workflow) |  | [optional] 

## Methods

### NewCreateCreateWorkflowRequest

`func NewCreateCreateWorkflowRequest() *CreateCreateWorkflowRequest`

NewCreateCreateWorkflowRequest instantiates a new CreateCreateWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCreateWorkflowRequestWithDefaults

`func NewCreateCreateWorkflowRequestWithDefaults() *CreateCreateWorkflowRequest`

NewCreateCreateWorkflowRequestWithDefaults instantiates a new CreateCreateWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *CreateCreateWorkflowRequest) GetWorkflow() CreateWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CreateCreateWorkflowRequest) GetWorkflowOk() (*CreateWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CreateCreateWorkflowRequest) SetWorkflow(v CreateWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CreateCreateWorkflowRequest) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


