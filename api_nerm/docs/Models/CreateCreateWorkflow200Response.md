---
id: nerm-create-create-workflow200-response
title: CreateCreateWorkflow200Response
pagination_label: CreateCreateWorkflow200Response
sidebar_label: CreateCreateWorkflow200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateCreateWorkflow200Response', 'NERMCreateCreateWorkflow200Response'] 
slug: /tools/sdk/go/nerm/models/create-create-workflow200-response
tags: ['SDK', 'Software Development Kit', 'CreateCreateWorkflow200Response', 'NERMCreateCreateWorkflow200Response']
---

# CreateCreateWorkflow200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | Pointer to [**WorkflowSession**](workflow-session) |  | [optional] 

## Methods

### NewCreateCreateWorkflow200Response

`func NewCreateCreateWorkflow200Response() *CreateCreateWorkflow200Response`

NewCreateCreateWorkflow200Response instantiates a new CreateCreateWorkflow200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCreateWorkflow200ResponseWithDefaults

`func NewCreateCreateWorkflow200ResponseWithDefaults() *CreateCreateWorkflow200Response`

NewCreateCreateWorkflow200ResponseWithDefaults instantiates a new CreateCreateWorkflow200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *CreateCreateWorkflow200Response) GetWorkflow() WorkflowSession`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CreateCreateWorkflow200Response) GetWorkflowOk() (*WorkflowSession, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CreateCreateWorkflow200Response) SetWorkflow(v WorkflowSession)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CreateCreateWorkflow200Response) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


