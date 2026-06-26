---
id: nerm-create-unassign-action-request
title: CreateUnassignActionRequest
pagination_label: CreateUnassignActionRequest
sidebar_label: CreateUnassignActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateUnassignActionRequest', 'NERMCreateUnassignActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-unassign-action-request
tags: ['SDK', 'Software Development Kit', 'CreateUnassignActionRequest', 'NERMCreateUnassignActionRequest']
---

# CreateUnassignActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**UnassignAction**](unassign-action) |  | [optional] 

## Methods

### NewCreateUnassignActionRequest

`func NewCreateUnassignActionRequest() *CreateUnassignActionRequest`

NewCreateUnassignActionRequest instantiates a new CreateUnassignActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUnassignActionRequestWithDefaults

`func NewCreateUnassignActionRequestWithDefaults() *CreateUnassignActionRequest`

NewCreateUnassignActionRequestWithDefaults instantiates a new CreateUnassignActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateUnassignActionRequest) GetWorkflowAction() UnassignAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateUnassignActionRequest) GetWorkflowActionOk() (*UnassignAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateUnassignActionRequest) SetWorkflowAction(v UnassignAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateUnassignActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


