---
id: nerm-create-approval-action-request
title: CreateApprovalActionRequest
pagination_label: CreateApprovalActionRequest
sidebar_label: CreateApprovalActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateApprovalActionRequest', 'NERMCreateApprovalActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-approval-action-request
tags: ['SDK', 'Software Development Kit', 'CreateApprovalActionRequest', 'NERMCreateApprovalActionRequest']
---

# CreateApprovalActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**ApprovalAction**](approval-action) |  | [optional] 

## Methods

### NewCreateApprovalActionRequest

`func NewCreateApprovalActionRequest() *CreateApprovalActionRequest`

NewCreateApprovalActionRequest instantiates a new CreateApprovalActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApprovalActionRequestWithDefaults

`func NewCreateApprovalActionRequestWithDefaults() *CreateApprovalActionRequest`

NewCreateApprovalActionRequestWithDefaults instantiates a new CreateApprovalActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateApprovalActionRequest) GetWorkflowAction() ApprovalAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateApprovalActionRequest) GetWorkflowActionOk() (*ApprovalAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateApprovalActionRequest) SetWorkflowAction(v ApprovalAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateApprovalActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


