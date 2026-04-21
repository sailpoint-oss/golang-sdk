---
id: nerm-create-update-profile-action-request
title: CreateUpdateProfileActionRequest
pagination_label: CreateUpdateProfileActionRequest
sidebar_label: CreateUpdateProfileActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateUpdateProfileActionRequest', 'NERMCreateUpdateProfileActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-update-profile-action-request
tags: ['SDK', 'Software Development Kit', 'CreateUpdateProfileActionRequest', 'NERMCreateUpdateProfileActionRequest']
---

# CreateUpdateProfileActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**UpdateProfileAction**](update-profile-action) |  | [optional] 

## Methods

### NewCreateUpdateProfileActionRequest

`func NewCreateUpdateProfileActionRequest() *CreateUpdateProfileActionRequest`

NewCreateUpdateProfileActionRequest instantiates a new CreateUpdateProfileActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateProfileActionRequestWithDefaults

`func NewCreateUpdateProfileActionRequestWithDefaults() *CreateUpdateProfileActionRequest`

NewCreateUpdateProfileActionRequestWithDefaults instantiates a new CreateUpdateProfileActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateUpdateProfileActionRequest) GetWorkflowAction() UpdateProfileAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateUpdateProfileActionRequest) GetWorkflowActionOk() (*UpdateProfileAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateUpdateProfileActionRequest) SetWorkflowAction(v UpdateProfileAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateUpdateProfileActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


