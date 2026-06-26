---
id: nerm-create-profile-check-action-request
title: CreateProfileCheckActionRequest
pagination_label: CreateProfileCheckActionRequest
sidebar_label: CreateProfileCheckActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateProfileCheckActionRequest', 'NERMCreateProfileCheckActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-profile-check-action-request
tags: ['SDK', 'Software Development Kit', 'CreateProfileCheckActionRequest', 'NERMCreateProfileCheckActionRequest']
---

# CreateProfileCheckActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**ProfileCheckAction**](profile-check-action) |  | [optional] 

## Methods

### NewCreateProfileCheckActionRequest

`func NewCreateProfileCheckActionRequest() *CreateProfileCheckActionRequest`

NewCreateProfileCheckActionRequest instantiates a new CreateProfileCheckActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileCheckActionRequestWithDefaults

`func NewCreateProfileCheckActionRequestWithDefaults() *CreateProfileCheckActionRequest`

NewCreateProfileCheckActionRequestWithDefaults instantiates a new CreateProfileCheckActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateProfileCheckActionRequest) GetWorkflowAction() ProfileCheckAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateProfileCheckActionRequest) GetWorkflowActionOk() (*ProfileCheckAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateProfileCheckActionRequest) SetWorkflowAction(v ProfileCheckAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateProfileCheckActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


