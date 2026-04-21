---
id: nerm-create-duplicate-prevention-action-request
title: CreateDuplicatePreventionActionRequest
pagination_label: CreateDuplicatePreventionActionRequest
sidebar_label: CreateDuplicatePreventionActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateDuplicatePreventionActionRequest', 'NERMCreateDuplicatePreventionActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-duplicate-prevention-action-request
tags: ['SDK', 'Software Development Kit', 'CreateDuplicatePreventionActionRequest', 'NERMCreateDuplicatePreventionActionRequest']
---

# CreateDuplicatePreventionActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**DuplicatePreventionAction**](duplicate-prevention-action) |  | [optional] 

## Methods

### NewCreateDuplicatePreventionActionRequest

`func NewCreateDuplicatePreventionActionRequest() *CreateDuplicatePreventionActionRequest`

NewCreateDuplicatePreventionActionRequest instantiates a new CreateDuplicatePreventionActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDuplicatePreventionActionRequestWithDefaults

`func NewCreateDuplicatePreventionActionRequestWithDefaults() *CreateDuplicatePreventionActionRequest`

NewCreateDuplicatePreventionActionRequestWithDefaults instantiates a new CreateDuplicatePreventionActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateDuplicatePreventionActionRequest) GetWorkflowAction() DuplicatePreventionAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateDuplicatePreventionActionRequest) GetWorkflowActionOk() (*DuplicatePreventionAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateDuplicatePreventionActionRequest) SetWorkflowAction(v DuplicatePreventionAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateDuplicatePreventionActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


