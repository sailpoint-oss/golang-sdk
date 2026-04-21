---
id: nerm-create-workflow-action-performer-request
title: CreateWorkflowActionPerformerRequest
pagination_label: CreateWorkflowActionPerformerRequest
sidebar_label: CreateWorkflowActionPerformerRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateWorkflowActionPerformerRequest', 'NERMCreateWorkflowActionPerformerRequest'] 
slug: /tools/sdk/go/nerm/models/create-workflow-action-performer-request
tags: ['SDK', 'Software Development Kit', 'CreateWorkflowActionPerformerRequest', 'NERMCreateWorkflowActionPerformerRequest']
---

# CreateWorkflowActionPerformerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowActionPerformers** | Pointer to [**WorkflowActionPerformers1**](workflow-action-performers1) |  | [optional] 

## Methods

### NewCreateWorkflowActionPerformerRequest

`func NewCreateWorkflowActionPerformerRequest() *CreateWorkflowActionPerformerRequest`

NewCreateWorkflowActionPerformerRequest instantiates a new CreateWorkflowActionPerformerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowActionPerformerRequestWithDefaults

`func NewCreateWorkflowActionPerformerRequestWithDefaults() *CreateWorkflowActionPerformerRequest`

NewCreateWorkflowActionPerformerRequestWithDefaults instantiates a new CreateWorkflowActionPerformerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowActionPerformers

`func (o *CreateWorkflowActionPerformerRequest) GetWorkflowActionPerformers() WorkflowActionPerformers1`

GetWorkflowActionPerformers returns the WorkflowActionPerformers field if non-nil, zero value otherwise.

### GetWorkflowActionPerformersOk

`func (o *CreateWorkflowActionPerformerRequest) GetWorkflowActionPerformersOk() (*WorkflowActionPerformers1, bool)`

GetWorkflowActionPerformersOk returns a tuple with the WorkflowActionPerformers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionPerformers

`func (o *CreateWorkflowActionPerformerRequest) SetWorkflowActionPerformers(v WorkflowActionPerformers1)`

SetWorkflowActionPerformers sets WorkflowActionPerformers field to given value.

### HasWorkflowActionPerformers

`func (o *CreateWorkflowActionPerformerRequest) HasWorkflowActionPerformers() bool`

HasWorkflowActionPerformers returns a boolean if a field has been set.


