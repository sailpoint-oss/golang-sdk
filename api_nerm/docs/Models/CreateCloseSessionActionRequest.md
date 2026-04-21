---
id: nerm-create-close-session-action-request
title: CreateCloseSessionActionRequest
pagination_label: CreateCloseSessionActionRequest
sidebar_label: CreateCloseSessionActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateCloseSessionActionRequest', 'NERMCreateCloseSessionActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-close-session-action-request
tags: ['SDK', 'Software Development Kit', 'CreateCloseSessionActionRequest', 'NERMCreateCloseSessionActionRequest']
---

# CreateCloseSessionActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**CloseSessionAction**](close-session-action) |  | [optional] 

## Methods

### NewCreateCloseSessionActionRequest

`func NewCreateCloseSessionActionRequest() *CreateCloseSessionActionRequest`

NewCreateCloseSessionActionRequest instantiates a new CreateCloseSessionActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCloseSessionActionRequestWithDefaults

`func NewCreateCloseSessionActionRequestWithDefaults() *CreateCloseSessionActionRequest`

NewCreateCloseSessionActionRequestWithDefaults instantiates a new CreateCloseSessionActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateCloseSessionActionRequest) GetWorkflowAction() CloseSessionAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateCloseSessionActionRequest) GetWorkflowActionOk() (*CloseSessionAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateCloseSessionActionRequest) SetWorkflowAction(v CloseSessionAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateCloseSessionActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


