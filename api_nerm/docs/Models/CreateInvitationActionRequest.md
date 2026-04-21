---
id: nerm-create-invitation-action-request
title: CreateInvitationActionRequest
pagination_label: CreateInvitationActionRequest
sidebar_label: CreateInvitationActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateInvitationActionRequest', 'NERMCreateInvitationActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-invitation-action-request
tags: ['SDK', 'Software Development Kit', 'CreateInvitationActionRequest', 'NERMCreateInvitationActionRequest']
---

# CreateInvitationActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**InvitationAction**](invitation-action) |  | [optional] 

## Methods

### NewCreateInvitationActionRequest

`func NewCreateInvitationActionRequest() *CreateInvitationActionRequest`

NewCreateInvitationActionRequest instantiates a new CreateInvitationActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvitationActionRequestWithDefaults

`func NewCreateInvitationActionRequestWithDefaults() *CreateInvitationActionRequest`

NewCreateInvitationActionRequestWithDefaults instantiates a new CreateInvitationActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateInvitationActionRequest) GetWorkflowAction() InvitationAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateInvitationActionRequest) GetWorkflowActionOk() (*InvitationAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateInvitationActionRequest) SetWorkflowAction(v InvitationAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateInvitationActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


