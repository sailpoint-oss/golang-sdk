---
id: nerm-create-run-workflow-action-request
title: CreateRunWorkflowActionRequest
pagination_label: CreateRunWorkflowActionRequest
sidebar_label: CreateRunWorkflowActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateRunWorkflowActionRequest', 'NERMCreateRunWorkflowActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-run-workflow-action-request
tags: ['SDK', 'Software Development Kit', 'CreateRunWorkflowActionRequest', 'NERMCreateRunWorkflowActionRequest']
---

# CreateRunWorkflowActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**RunWorkflowAction**](run-workflow-action) |  | [optional] 

## Methods

### NewCreateRunWorkflowActionRequest

`func NewCreateRunWorkflowActionRequest() *CreateRunWorkflowActionRequest`

NewCreateRunWorkflowActionRequest instantiates a new CreateRunWorkflowActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRunWorkflowActionRequestWithDefaults

`func NewCreateRunWorkflowActionRequestWithDefaults() *CreateRunWorkflowActionRequest`

NewCreateRunWorkflowActionRequestWithDefaults instantiates a new CreateRunWorkflowActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateRunWorkflowActionRequest) GetWorkflowAction() RunWorkflowAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateRunWorkflowActionRequest) GetWorkflowActionOk() (*RunWorkflowAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateRunWorkflowActionRequest) SetWorkflowAction(v RunWorkflowAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateRunWorkflowActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


