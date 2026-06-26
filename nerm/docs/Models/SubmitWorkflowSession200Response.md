---
id: nerm-submit-workflow-session200-response
title: SubmitWorkflowSession200Response
pagination_label: SubmitWorkflowSession200Response
sidebar_label: SubmitWorkflowSession200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitWorkflowSession200Response', 'NERMSubmitWorkflowSession200Response'] 
slug: /tools/sdk/go/nerm/models/submit-workflow-session200-response
tags: ['SDK', 'Software Development Kit', 'SubmitWorkflowSession200Response', 'NERMSubmitWorkflowSession200Response']
---

# SubmitWorkflowSession200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowSession** | Pointer to [**WorkflowSession**](workflow-session) |  | [optional] 

## Methods

### NewSubmitWorkflowSession200Response

`func NewSubmitWorkflowSession200Response() *SubmitWorkflowSession200Response`

NewSubmitWorkflowSession200Response instantiates a new SubmitWorkflowSession200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitWorkflowSession200ResponseWithDefaults

`func NewSubmitWorkflowSession200ResponseWithDefaults() *SubmitWorkflowSession200Response`

NewSubmitWorkflowSession200ResponseWithDefaults instantiates a new SubmitWorkflowSession200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowSession

`func (o *SubmitWorkflowSession200Response) GetWorkflowSession() WorkflowSession`

GetWorkflowSession returns the WorkflowSession field if non-nil, zero value otherwise.

### GetWorkflowSessionOk

`func (o *SubmitWorkflowSession200Response) GetWorkflowSessionOk() (*WorkflowSession, bool)`

GetWorkflowSessionOk returns a tuple with the WorkflowSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSession

`func (o *SubmitWorkflowSession200Response) SetWorkflowSession(v WorkflowSession)`

SetWorkflowSession sets WorkflowSession field to given value.

### HasWorkflowSession

`func (o *SubmitWorkflowSession200Response) HasWorkflowSession() bool`

HasWorkflowSession returns a boolean if a field has been set.


