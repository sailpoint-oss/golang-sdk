---
id: nerm-create-request-action-request
title: CreateRequestActionRequest
pagination_label: CreateRequestActionRequest
sidebar_label: CreateRequestActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateRequestActionRequest', 'NERMCreateRequestActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-request-action-request
tags: ['SDK', 'Software Development Kit', 'CreateRequestActionRequest', 'NERMCreateRequestActionRequest']
---

# CreateRequestActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**RequestAction**](request-action) |  | [optional] 

## Methods

### NewCreateRequestActionRequest

`func NewCreateRequestActionRequest() *CreateRequestActionRequest`

NewCreateRequestActionRequest instantiates a new CreateRequestActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestActionRequestWithDefaults

`func NewCreateRequestActionRequestWithDefaults() *CreateRequestActionRequest`

NewCreateRequestActionRequestWithDefaults instantiates a new CreateRequestActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateRequestActionRequest) GetWorkflowAction() RequestAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateRequestActionRequest) GetWorkflowActionOk() (*RequestAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateRequestActionRequest) SetWorkflowAction(v RequestAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateRequestActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


