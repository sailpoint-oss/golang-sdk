---
id: nerm-create-ask-security-question-action-request
title: CreateAskSecurityQuestionActionRequest
pagination_label: CreateAskSecurityQuestionActionRequest
sidebar_label: CreateAskSecurityQuestionActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateAskSecurityQuestionActionRequest', 'NERMCreateAskSecurityQuestionActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-ask-security-question-action-request
tags: ['SDK', 'Software Development Kit', 'CreateAskSecurityQuestionActionRequest', 'NERMCreateAskSecurityQuestionActionRequest']
---

# CreateAskSecurityQuestionActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**AskSecurityQuestionAction**](ask-security-question-action) |  | [optional] 

## Methods

### NewCreateAskSecurityQuestionActionRequest

`func NewCreateAskSecurityQuestionActionRequest() *CreateAskSecurityQuestionActionRequest`

NewCreateAskSecurityQuestionActionRequest instantiates a new CreateAskSecurityQuestionActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAskSecurityQuestionActionRequestWithDefaults

`func NewCreateAskSecurityQuestionActionRequestWithDefaults() *CreateAskSecurityQuestionActionRequest`

NewCreateAskSecurityQuestionActionRequestWithDefaults instantiates a new CreateAskSecurityQuestionActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateAskSecurityQuestionActionRequest) GetWorkflowAction() AskSecurityQuestionAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateAskSecurityQuestionActionRequest) GetWorkflowActionOk() (*AskSecurityQuestionAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateAskSecurityQuestionActionRequest) SetWorkflowAction(v AskSecurityQuestionAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateAskSecurityQuestionActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


