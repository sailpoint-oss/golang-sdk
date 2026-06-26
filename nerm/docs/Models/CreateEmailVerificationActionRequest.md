---
id: nerm-create-email-verification-action-request
title: CreateEmailVerificationActionRequest
pagination_label: CreateEmailVerificationActionRequest
sidebar_label: CreateEmailVerificationActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateEmailVerificationActionRequest', 'NERMCreateEmailVerificationActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-email-verification-action-request
tags: ['SDK', 'Software Development Kit', 'CreateEmailVerificationActionRequest', 'NERMCreateEmailVerificationActionRequest']
---

# CreateEmailVerificationActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**EmailVerificationAction**](email-verification-action) |  | [optional] 

## Methods

### NewCreateEmailVerificationActionRequest

`func NewCreateEmailVerificationActionRequest() *CreateEmailVerificationActionRequest`

NewCreateEmailVerificationActionRequest instantiates a new CreateEmailVerificationActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailVerificationActionRequestWithDefaults

`func NewCreateEmailVerificationActionRequestWithDefaults() *CreateEmailVerificationActionRequest`

NewCreateEmailVerificationActionRequestWithDefaults instantiates a new CreateEmailVerificationActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateEmailVerificationActionRequest) GetWorkflowAction() EmailVerificationAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateEmailVerificationActionRequest) GetWorkflowActionOk() (*EmailVerificationAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateEmailVerificationActionRequest) SetWorkflowAction(v EmailVerificationAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateEmailVerificationActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


