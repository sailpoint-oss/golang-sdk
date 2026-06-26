---
id: nerm-create-approval-action200-response
title: CreateApprovalAction200Response
pagination_label: CreateApprovalAction200Response
sidebar_label: CreateApprovalAction200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateApprovalAction200Response', 'NERMCreateApprovalAction200Response'] 
slug: /tools/sdk/go/nerm/models/create-approval-action200-response
tags: ['SDK', 'Software Development Kit', 'CreateApprovalAction200Response', 'NERMCreateApprovalAction200Response']
---

# CreateApprovalAction200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**WorkflowAction**](workflow-action) |  | [optional] 

## Methods

### NewCreateApprovalAction200Response

`func NewCreateApprovalAction200Response() *CreateApprovalAction200Response`

NewCreateApprovalAction200Response instantiates a new CreateApprovalAction200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApprovalAction200ResponseWithDefaults

`func NewCreateApprovalAction200ResponseWithDefaults() *CreateApprovalAction200Response`

NewCreateApprovalAction200ResponseWithDefaults instantiates a new CreateApprovalAction200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateApprovalAction200Response) GetWorkflowAction() WorkflowAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateApprovalAction200Response) GetWorkflowActionOk() (*WorkflowAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateApprovalAction200Response) SetWorkflowAction(v WorkflowAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateApprovalAction200Response) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


