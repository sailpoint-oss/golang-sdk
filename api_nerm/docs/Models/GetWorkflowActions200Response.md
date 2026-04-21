---
id: nerm-get-workflow-actions200-response
title: GetWorkflowActions200Response
pagination_label: GetWorkflowActions200Response
sidebar_label: GetWorkflowActions200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetWorkflowActions200Response', 'NERMGetWorkflowActions200Response'] 
slug: /tools/sdk/go/nerm/models/get-workflow-actions200-response
tags: ['SDK', 'Software Development Kit', 'GetWorkflowActions200Response', 'NERMGetWorkflowActions200Response']
---

# GetWorkflowActions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowActions** | Pointer to [**WorkflowAction**](workflow-action) |  | [optional] 

## Methods

### NewGetWorkflowActions200Response

`func NewGetWorkflowActions200Response() *GetWorkflowActions200Response`

NewGetWorkflowActions200Response instantiates a new GetWorkflowActions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflowActions200ResponseWithDefaults

`func NewGetWorkflowActions200ResponseWithDefaults() *GetWorkflowActions200Response`

NewGetWorkflowActions200ResponseWithDefaults instantiates a new GetWorkflowActions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowActions

`func (o *GetWorkflowActions200Response) GetWorkflowActions() WorkflowAction`

GetWorkflowActions returns the WorkflowActions field if non-nil, zero value otherwise.

### GetWorkflowActionsOk

`func (o *GetWorkflowActions200Response) GetWorkflowActionsOk() (*WorkflowAction, bool)`

GetWorkflowActionsOk returns a tuple with the WorkflowActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActions

`func (o *GetWorkflowActions200Response) SetWorkflowActions(v WorkflowAction)`

SetWorkflowActions sets WorkflowActions field to given value.

### HasWorkflowActions

`func (o *GetWorkflowActions200Response) HasWorkflowActions() bool`

HasWorkflowActions returns a boolean if a field has been set.


