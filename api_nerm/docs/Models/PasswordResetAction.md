---
id: nerm-password-reset-action
title: PasswordResetAction
pagination_label: PasswordResetAction
sidebar_label: PasswordResetAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PasswordResetAction', 'NERMPasswordResetAction'] 
slug: /tools/sdk/go/nerm/models/password-reset-action
tags: ['SDK', 'Software Development Kit', 'PasswordResetAction', 'NERMPasswordResetAction']
---

# PasswordResetAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewPasswordResetAction

`func NewPasswordResetAction(workflowId string, description string, ) *PasswordResetAction`

NewPasswordResetAction instantiates a new PasswordResetAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordResetActionWithDefaults

`func NewPasswordResetActionWithDefaults() *PasswordResetAction`

NewPasswordResetActionWithDefaults instantiates a new PasswordResetAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *PasswordResetAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *PasswordResetAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *PasswordResetAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *PasswordResetAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PasswordResetAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PasswordResetAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *PasswordResetAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PasswordResetAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PasswordResetAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PasswordResetAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


