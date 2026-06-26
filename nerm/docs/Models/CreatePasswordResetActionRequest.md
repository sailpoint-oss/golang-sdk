---
id: nerm-create-password-reset-action-request
title: CreatePasswordResetActionRequest
pagination_label: CreatePasswordResetActionRequest
sidebar_label: CreatePasswordResetActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreatePasswordResetActionRequest', 'NERMCreatePasswordResetActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-password-reset-action-request
tags: ['SDK', 'Software Development Kit', 'CreatePasswordResetActionRequest', 'NERMCreatePasswordResetActionRequest']
---

# CreatePasswordResetActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**PasswordResetAction**](password-reset-action) |  | [optional] 

## Methods

### NewCreatePasswordResetActionRequest

`func NewCreatePasswordResetActionRequest() *CreatePasswordResetActionRequest`

NewCreatePasswordResetActionRequest instantiates a new CreatePasswordResetActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePasswordResetActionRequestWithDefaults

`func NewCreatePasswordResetActionRequestWithDefaults() *CreatePasswordResetActionRequest`

NewCreatePasswordResetActionRequestWithDefaults instantiates a new CreatePasswordResetActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreatePasswordResetActionRequest) GetWorkflowAction() PasswordResetAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreatePasswordResetActionRequest) GetWorkflowActionOk() (*PasswordResetAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreatePasswordResetActionRequest) SetWorkflowAction(v PasswordResetAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreatePasswordResetActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


