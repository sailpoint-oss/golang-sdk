---
id: nerm-create-set-attributes-action-request
title: CreateSetAttributesActionRequest
pagination_label: CreateSetAttributesActionRequest
sidebar_label: CreateSetAttributesActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateSetAttributesActionRequest', 'NERMCreateSetAttributesActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-set-attributes-action-request
tags: ['SDK', 'Software Development Kit', 'CreateSetAttributesActionRequest', 'NERMCreateSetAttributesActionRequest']
---

# CreateSetAttributesActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**SetAttributesAction**](set-attributes-action) |  | [optional] 

## Methods

### NewCreateSetAttributesActionRequest

`func NewCreateSetAttributesActionRequest() *CreateSetAttributesActionRequest`

NewCreateSetAttributesActionRequest instantiates a new CreateSetAttributesActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSetAttributesActionRequestWithDefaults

`func NewCreateSetAttributesActionRequestWithDefaults() *CreateSetAttributesActionRequest`

NewCreateSetAttributesActionRequestWithDefaults instantiates a new CreateSetAttributesActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateSetAttributesActionRequest) GetWorkflowAction() SetAttributesAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateSetAttributesActionRequest) GetWorkflowActionOk() (*SetAttributesAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateSetAttributesActionRequest) SetWorkflowAction(v SetAttributesAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateSetAttributesActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


