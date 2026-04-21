---
id: nerm-identity-proofing-action
title: IdentityProofingAction
pagination_label: IdentityProofingAction
sidebar_label: IdentityProofingAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IdentityProofingAction', 'NERMIdentityProofingAction'] 
slug: /tools/sdk/go/nerm/models/identity-proofing-action
tags: ['SDK', 'Software Development Kit', 'IdentityProofingAction', 'NERMIdentityProofingAction']
---

# IdentityProofingAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewIdentityProofingAction

`func NewIdentityProofingAction(workflowId string, description string, ) *IdentityProofingAction`

NewIdentityProofingAction instantiates a new IdentityProofingAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProofingActionWithDefaults

`func NewIdentityProofingActionWithDefaults() *IdentityProofingAction`

NewIdentityProofingActionWithDefaults instantiates a new IdentityProofingAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *IdentityProofingAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *IdentityProofingAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *IdentityProofingAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *IdentityProofingAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProofingAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProofingAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *IdentityProofingAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *IdentityProofingAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *IdentityProofingAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *IdentityProofingAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


