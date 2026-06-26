---
id: nerm-create-batch-update-action-request
title: CreateBatchUpdateActionRequest
pagination_label: CreateBatchUpdateActionRequest
sidebar_label: CreateBatchUpdateActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateBatchUpdateActionRequest', 'NERMCreateBatchUpdateActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-batch-update-action-request
tags: ['SDK', 'Software Development Kit', 'CreateBatchUpdateActionRequest', 'NERMCreateBatchUpdateActionRequest']
---

# CreateBatchUpdateActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**BatchUpdateAction**](batch-update-action) |  | [optional] 

## Methods

### NewCreateBatchUpdateActionRequest

`func NewCreateBatchUpdateActionRequest() *CreateBatchUpdateActionRequest`

NewCreateBatchUpdateActionRequest instantiates a new CreateBatchUpdateActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchUpdateActionRequestWithDefaults

`func NewCreateBatchUpdateActionRequestWithDefaults() *CreateBatchUpdateActionRequest`

NewCreateBatchUpdateActionRequestWithDefaults instantiates a new CreateBatchUpdateActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateBatchUpdateActionRequest) GetWorkflowAction() BatchUpdateAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateBatchUpdateActionRequest) GetWorkflowActionOk() (*BatchUpdateAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateBatchUpdateActionRequest) SetWorkflowAction(v BatchUpdateAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateBatchUpdateActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


