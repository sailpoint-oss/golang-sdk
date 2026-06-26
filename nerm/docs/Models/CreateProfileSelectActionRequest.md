---
id: nerm-create-profile-select-action-request
title: CreateProfileSelectActionRequest
pagination_label: CreateProfileSelectActionRequest
sidebar_label: CreateProfileSelectActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateProfileSelectActionRequest', 'NERMCreateProfileSelectActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-profile-select-action-request
tags: ['SDK', 'Software Development Kit', 'CreateProfileSelectActionRequest', 'NERMCreateProfileSelectActionRequest']
---

# CreateProfileSelectActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**ProfileSelectAction**](profile-select-action) |  | [optional] 

## Methods

### NewCreateProfileSelectActionRequest

`func NewCreateProfileSelectActionRequest() *CreateProfileSelectActionRequest`

NewCreateProfileSelectActionRequest instantiates a new CreateProfileSelectActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileSelectActionRequestWithDefaults

`func NewCreateProfileSelectActionRequestWithDefaults() *CreateProfileSelectActionRequest`

NewCreateProfileSelectActionRequestWithDefaults instantiates a new CreateProfileSelectActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateProfileSelectActionRequest) GetWorkflowAction() ProfileSelectAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateProfileSelectActionRequest) GetWorkflowActionOk() (*ProfileSelectAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateProfileSelectActionRequest) SetWorkflowAction(v ProfileSelectAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateProfileSelectActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


