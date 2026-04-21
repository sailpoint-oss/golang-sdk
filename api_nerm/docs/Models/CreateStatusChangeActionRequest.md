---
id: nerm-create-status-change-action-request
title: CreateStatusChangeActionRequest
pagination_label: CreateStatusChangeActionRequest
sidebar_label: CreateStatusChangeActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateStatusChangeActionRequest', 'NERMCreateStatusChangeActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-status-change-action-request
tags: ['SDK', 'Software Development Kit', 'CreateStatusChangeActionRequest', 'NERMCreateStatusChangeActionRequest']
---

# CreateStatusChangeActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**StatusChangeAction**](status-change-action) |  | [optional] 

## Methods

### NewCreateStatusChangeActionRequest

`func NewCreateStatusChangeActionRequest() *CreateStatusChangeActionRequest`

NewCreateStatusChangeActionRequest instantiates a new CreateStatusChangeActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStatusChangeActionRequestWithDefaults

`func NewCreateStatusChangeActionRequestWithDefaults() *CreateStatusChangeActionRequest`

NewCreateStatusChangeActionRequestWithDefaults instantiates a new CreateStatusChangeActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateStatusChangeActionRequest) GetWorkflowAction() StatusChangeAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateStatusChangeActionRequest) GetWorkflowActionOk() (*StatusChangeAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateStatusChangeActionRequest) SetWorkflowAction(v StatusChangeAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateStatusChangeActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


