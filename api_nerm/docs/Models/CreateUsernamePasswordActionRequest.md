---
id: nerm-create-username-password-action-request
title: CreateUsernamePasswordActionRequest
pagination_label: CreateUsernamePasswordActionRequest
sidebar_label: CreateUsernamePasswordActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateUsernamePasswordActionRequest', 'NERMCreateUsernamePasswordActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-username-password-action-request
tags: ['SDK', 'Software Development Kit', 'CreateUsernamePasswordActionRequest', 'NERMCreateUsernamePasswordActionRequest']
---

# CreateUsernamePasswordActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**UsernamePasswordAction**](username-password-action) |  | [optional] 

## Methods

### NewCreateUsernamePasswordActionRequest

`func NewCreateUsernamePasswordActionRequest() *CreateUsernamePasswordActionRequest`

NewCreateUsernamePasswordActionRequest instantiates a new CreateUsernamePasswordActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUsernamePasswordActionRequestWithDefaults

`func NewCreateUsernamePasswordActionRequestWithDefaults() *CreateUsernamePasswordActionRequest`

NewCreateUsernamePasswordActionRequestWithDefaults instantiates a new CreateUsernamePasswordActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateUsernamePasswordActionRequest) GetWorkflowAction() UsernamePasswordAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateUsernamePasswordActionRequest) GetWorkflowActionOk() (*UsernamePasswordAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateUsernamePasswordActionRequest) SetWorkflowAction(v UsernamePasswordAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateUsernamePasswordActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


