---
id: nerm-create-fulfillment-action-request
title: CreateFulfillmentActionRequest
pagination_label: CreateFulfillmentActionRequest
sidebar_label: CreateFulfillmentActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateFulfillmentActionRequest', 'NERMCreateFulfillmentActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-fulfillment-action-request
tags: ['SDK', 'Software Development Kit', 'CreateFulfillmentActionRequest', 'NERMCreateFulfillmentActionRequest']
---

# CreateFulfillmentActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**FulfillmentAction**](fulfillment-action) |  | [optional] 

## Methods

### NewCreateFulfillmentActionRequest

`func NewCreateFulfillmentActionRequest() *CreateFulfillmentActionRequest`

NewCreateFulfillmentActionRequest instantiates a new CreateFulfillmentActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFulfillmentActionRequestWithDefaults

`func NewCreateFulfillmentActionRequestWithDefaults() *CreateFulfillmentActionRequest`

NewCreateFulfillmentActionRequestWithDefaults instantiates a new CreateFulfillmentActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateFulfillmentActionRequest) GetWorkflowAction() FulfillmentAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateFulfillmentActionRequest) GetWorkflowActionOk() (*FulfillmentAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateFulfillmentActionRequest) SetWorkflowAction(v FulfillmentAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateFulfillmentActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


