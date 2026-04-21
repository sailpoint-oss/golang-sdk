---
id: nerm-create-soap-api-action-request
title: CreateSoapApiActionRequest
pagination_label: CreateSoapApiActionRequest
sidebar_label: CreateSoapApiActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateSoapApiActionRequest', 'NERMCreateSoapApiActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-soap-api-action-request
tags: ['SDK', 'Software Development Kit', 'CreateSoapApiActionRequest', 'NERMCreateSoapApiActionRequest']
---

# CreateSoapApiActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**SoapApiAction**](soap-api-action) |  | [optional] 

## Methods

### NewCreateSoapApiActionRequest

`func NewCreateSoapApiActionRequest() *CreateSoapApiActionRequest`

NewCreateSoapApiActionRequest instantiates a new CreateSoapApiActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSoapApiActionRequestWithDefaults

`func NewCreateSoapApiActionRequestWithDefaults() *CreateSoapApiActionRequest`

NewCreateSoapApiActionRequestWithDefaults instantiates a new CreateSoapApiActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateSoapApiActionRequest) GetWorkflowAction() SoapApiAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateSoapApiActionRequest) GetWorkflowActionOk() (*SoapApiAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateSoapApiActionRequest) SetWorkflowAction(v SoapApiAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateSoapApiActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


