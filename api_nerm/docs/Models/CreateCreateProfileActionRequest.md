---
id: nerm-create-create-profile-action-request
title: CreateCreateProfileActionRequest
pagination_label: CreateCreateProfileActionRequest
sidebar_label: CreateCreateProfileActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateCreateProfileActionRequest', 'NERMCreateCreateProfileActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-create-profile-action-request
tags: ['SDK', 'Software Development Kit', 'CreateCreateProfileActionRequest', 'NERMCreateCreateProfileActionRequest']
---

# CreateCreateProfileActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**CreateProfileAction**](create-profile-action) |  | [optional] 

## Methods

### NewCreateCreateProfileActionRequest

`func NewCreateCreateProfileActionRequest() *CreateCreateProfileActionRequest`

NewCreateCreateProfileActionRequest instantiates a new CreateCreateProfileActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCreateProfileActionRequestWithDefaults

`func NewCreateCreateProfileActionRequestWithDefaults() *CreateCreateProfileActionRequest`

NewCreateCreateProfileActionRequestWithDefaults instantiates a new CreateCreateProfileActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateCreateProfileActionRequest) GetWorkflowAction() CreateProfileAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateCreateProfileActionRequest) GetWorkflowActionOk() (*CreateProfileAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateCreateProfileActionRequest) SetWorkflowAction(v CreateProfileAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateCreateProfileActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


