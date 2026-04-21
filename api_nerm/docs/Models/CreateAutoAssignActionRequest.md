---
id: nerm-create-auto-assign-action-request
title: CreateAutoAssignActionRequest
pagination_label: CreateAutoAssignActionRequest
sidebar_label: CreateAutoAssignActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateAutoAssignActionRequest', 'NERMCreateAutoAssignActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-auto-assign-action-request
tags: ['SDK', 'Software Development Kit', 'CreateAutoAssignActionRequest', 'NERMCreateAutoAssignActionRequest']
---

# CreateAutoAssignActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**AutoAssignAction**](auto-assign-action) |  | [optional] 

## Methods

### NewCreateAutoAssignActionRequest

`func NewCreateAutoAssignActionRequest() *CreateAutoAssignActionRequest`

NewCreateAutoAssignActionRequest instantiates a new CreateAutoAssignActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoAssignActionRequestWithDefaults

`func NewCreateAutoAssignActionRequestWithDefaults() *CreateAutoAssignActionRequest`

NewCreateAutoAssignActionRequestWithDefaults instantiates a new CreateAutoAssignActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateAutoAssignActionRequest) GetWorkflowAction() AutoAssignAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateAutoAssignActionRequest) GetWorkflowActionOk() (*AutoAssignAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateAutoAssignActionRequest) SetWorkflowAction(v AutoAssignAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateAutoAssignActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


