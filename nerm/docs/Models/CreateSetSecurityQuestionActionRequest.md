---
id: nerm-create-set-security-question-action-request
title: CreateSetSecurityQuestionActionRequest
pagination_label: CreateSetSecurityQuestionActionRequest
sidebar_label: CreateSetSecurityQuestionActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateSetSecurityQuestionActionRequest', 'NERMCreateSetSecurityQuestionActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-set-security-question-action-request
tags: ['SDK', 'Software Development Kit', 'CreateSetSecurityQuestionActionRequest', 'NERMCreateSetSecurityQuestionActionRequest']
---

# CreateSetSecurityQuestionActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**SetSecurityQuestionAction**](set-security-question-action) |  | [optional] 

## Methods

### NewCreateSetSecurityQuestionActionRequest

`func NewCreateSetSecurityQuestionActionRequest() *CreateSetSecurityQuestionActionRequest`

NewCreateSetSecurityQuestionActionRequest instantiates a new CreateSetSecurityQuestionActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSetSecurityQuestionActionRequestWithDefaults

`func NewCreateSetSecurityQuestionActionRequestWithDefaults() *CreateSetSecurityQuestionActionRequest`

NewCreateSetSecurityQuestionActionRequestWithDefaults instantiates a new CreateSetSecurityQuestionActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateSetSecurityQuestionActionRequest) GetWorkflowAction() SetSecurityQuestionAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateSetSecurityQuestionActionRequest) GetWorkflowActionOk() (*SetSecurityQuestionAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateSetSecurityQuestionActionRequest) SetWorkflowAction(v SetSecurityQuestionAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateSetSecurityQuestionActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


