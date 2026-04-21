---
id: nerm-create-identity-proofing-action-request
title: CreateIdentityProofingActionRequest
pagination_label: CreateIdentityProofingActionRequest
sidebar_label: CreateIdentityProofingActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateIdentityProofingActionRequest', 'NERMCreateIdentityProofingActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-identity-proofing-action-request
tags: ['SDK', 'Software Development Kit', 'CreateIdentityProofingActionRequest', 'NERMCreateIdentityProofingActionRequest']
---

# CreateIdentityProofingActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**IdentityProofingAction**](identity-proofing-action) |  | [optional] 

## Methods

### NewCreateIdentityProofingActionRequest

`func NewCreateIdentityProofingActionRequest() *CreateIdentityProofingActionRequest`

NewCreateIdentityProofingActionRequest instantiates a new CreateIdentityProofingActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdentityProofingActionRequestWithDefaults

`func NewCreateIdentityProofingActionRequestWithDefaults() *CreateIdentityProofingActionRequest`

NewCreateIdentityProofingActionRequestWithDefaults instantiates a new CreateIdentityProofingActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateIdentityProofingActionRequest) GetWorkflowAction() IdentityProofingAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateIdentityProofingActionRequest) GetWorkflowActionOk() (*IdentityProofingAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateIdentityProofingActionRequest) SetWorkflowAction(v IdentityProofingAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateIdentityProofingActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


