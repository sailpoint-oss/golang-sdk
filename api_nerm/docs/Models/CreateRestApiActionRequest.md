---
id: nerm-create-rest-api-action-request
title: CreateRestApiActionRequest
pagination_label: CreateRestApiActionRequest
sidebar_label: CreateRestApiActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateRestApiActionRequest', 'NERMCreateRestApiActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-rest-api-action-request
tags: ['SDK', 'Software Development Kit', 'CreateRestApiActionRequest', 'NERMCreateRestApiActionRequest']
---

# CreateRestApiActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**RestApiAction**](rest-api-action) |  | [optional] 

## Methods

### NewCreateRestApiActionRequest

`func NewCreateRestApiActionRequest() *CreateRestApiActionRequest`

NewCreateRestApiActionRequest instantiates a new CreateRestApiActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRestApiActionRequestWithDefaults

`func NewCreateRestApiActionRequestWithDefaults() *CreateRestApiActionRequest`

NewCreateRestApiActionRequestWithDefaults instantiates a new CreateRestApiActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateRestApiActionRequest) GetWorkflowAction() RestApiAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateRestApiActionRequest) GetWorkflowActionOk() (*RestApiAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateRestApiActionRequest) SetWorkflowAction(v RestApiAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateRestApiActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


