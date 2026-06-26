---
id: nerm-submit-workflow-session-request
title: SubmitWorkflowSessionRequest
pagination_label: SubmitWorkflowSessionRequest
sidebar_label: SubmitWorkflowSessionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitWorkflowSessionRequest', 'NERMSubmitWorkflowSessionRequest'] 
slug: /tools/sdk/go/nerm/models/submit-workflow-session-request
tags: ['SDK', 'Software Development Kit', 'SubmitWorkflowSessionRequest', 'NERMSubmitWorkflowSessionRequest']
---

# SubmitWorkflowSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowSession** | Pointer to [**WorkflowSession1**](workflow-session1) |  | [optional] 

## Methods

### NewSubmitWorkflowSessionRequest

`func NewSubmitWorkflowSessionRequest() *SubmitWorkflowSessionRequest`

NewSubmitWorkflowSessionRequest instantiates a new SubmitWorkflowSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitWorkflowSessionRequestWithDefaults

`func NewSubmitWorkflowSessionRequestWithDefaults() *SubmitWorkflowSessionRequest`

NewSubmitWorkflowSessionRequestWithDefaults instantiates a new SubmitWorkflowSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowSession

`func (o *SubmitWorkflowSessionRequest) GetWorkflowSession() WorkflowSession1`

GetWorkflowSession returns the WorkflowSession field if non-nil, zero value otherwise.

### GetWorkflowSessionOk

`func (o *SubmitWorkflowSessionRequest) GetWorkflowSessionOk() (*WorkflowSession1, bool)`

GetWorkflowSessionOk returns a tuple with the WorkflowSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSession

`func (o *SubmitWorkflowSessionRequest) SetWorkflowSession(v WorkflowSession1)`

SetWorkflowSession sets WorkflowSession field to given value.

### HasWorkflowSession

`func (o *SubmitWorkflowSessionRequest) HasWorkflowSession() bool`

HasWorkflowSession returns a boolean if a field has been set.


