---
id: nerm-username-password-action
title: UsernamePasswordAction
pagination_label: UsernamePasswordAction
sidebar_label: UsernamePasswordAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UsernamePasswordAction', 'NERMUsernamePasswordAction'] 
slug: /tools/sdk/go/nerm/models/username-password-action
tags: ['SDK', 'Software Development Kit', 'UsernamePasswordAction', 'NERMUsernamePasswordAction']
---

# UsernamePasswordAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewUsernamePasswordAction

`func NewUsernamePasswordAction(workflowId string, description string, ) *UsernamePasswordAction`

NewUsernamePasswordAction instantiates a new UsernamePasswordAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsernamePasswordActionWithDefaults

`func NewUsernamePasswordActionWithDefaults() *UsernamePasswordAction`

NewUsernamePasswordActionWithDefaults instantiates a new UsernamePasswordAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *UsernamePasswordAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *UsernamePasswordAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *UsernamePasswordAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *UsernamePasswordAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UsernamePasswordAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UsernamePasswordAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *UsernamePasswordAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UsernamePasswordAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UsernamePasswordAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UsernamePasswordAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


