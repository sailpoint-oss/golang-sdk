---
id: nerm-create-login-workflow-request
title: CreateLoginWorkflowRequest
pagination_label: CreateLoginWorkflowRequest
sidebar_label: CreateLoginWorkflowRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateLoginWorkflowRequest', 'NERMCreateLoginWorkflowRequest'] 
slug: /tools/sdk/go/nerm/models/create-login-workflow-request
tags: ['SDK', 'Software Development Kit', 'CreateLoginWorkflowRequest', 'NERMCreateLoginWorkflowRequest']
---

# CreateLoginWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | Pointer to [**LoginWorkflow**](login-workflow) |  | [optional] 

## Methods

### NewCreateLoginWorkflowRequest

`func NewCreateLoginWorkflowRequest() *CreateLoginWorkflowRequest`

NewCreateLoginWorkflowRequest instantiates a new CreateLoginWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoginWorkflowRequestWithDefaults

`func NewCreateLoginWorkflowRequestWithDefaults() *CreateLoginWorkflowRequest`

NewCreateLoginWorkflowRequestWithDefaults instantiates a new CreateLoginWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *CreateLoginWorkflowRequest) GetWorkflow() LoginWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CreateLoginWorkflowRequest) GetWorkflowOk() (*LoginWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CreateLoginWorkflowRequest) SetWorkflow(v LoginWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CreateLoginWorkflowRequest) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


