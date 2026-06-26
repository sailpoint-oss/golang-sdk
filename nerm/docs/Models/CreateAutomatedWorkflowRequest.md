---
id: nerm-create-automated-workflow-request
title: CreateAutomatedWorkflowRequest
pagination_label: CreateAutomatedWorkflowRequest
sidebar_label: CreateAutomatedWorkflowRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateAutomatedWorkflowRequest', 'NERMCreateAutomatedWorkflowRequest'] 
slug: /tools/sdk/go/nerm/models/create-automated-workflow-request
tags: ['SDK', 'Software Development Kit', 'CreateAutomatedWorkflowRequest', 'NERMCreateAutomatedWorkflowRequest']
---

# CreateAutomatedWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | Pointer to [**AutomatedWorkflow**](automated-workflow) |  | [optional] 

## Methods

### NewCreateAutomatedWorkflowRequest

`func NewCreateAutomatedWorkflowRequest() *CreateAutomatedWorkflowRequest`

NewCreateAutomatedWorkflowRequest instantiates a new CreateAutomatedWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomatedWorkflowRequestWithDefaults

`func NewCreateAutomatedWorkflowRequestWithDefaults() *CreateAutomatedWorkflowRequest`

NewCreateAutomatedWorkflowRequestWithDefaults instantiates a new CreateAutomatedWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *CreateAutomatedWorkflowRequest) GetWorkflow() AutomatedWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CreateAutomatedWorkflowRequest) GetWorkflowOk() (*AutomatedWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CreateAutomatedWorkflowRequest) SetWorkflow(v AutomatedWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CreateAutomatedWorkflowRequest) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


