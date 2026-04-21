---
id: nerm-unassign-action
title: UnassignAction
pagination_label: UnassignAction
sidebar_label: UnassignAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UnassignAction', 'NERMUnassignAction'] 
slug: /tools/sdk/go/nerm/models/unassign-action
tags: ['SDK', 'Software Development Kit', 'UnassignAction', 'NERMUnassignAction']
---

# UnassignAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewUnassignAction

`func NewUnassignAction(workflowId string, description string, ) *UnassignAction`

NewUnassignAction instantiates a new UnassignAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnassignActionWithDefaults

`func NewUnassignActionWithDefaults() *UnassignAction`

NewUnassignActionWithDefaults instantiates a new UnassignAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *UnassignAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *UnassignAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *UnassignAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *UnassignAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnassignAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnassignAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *UnassignAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UnassignAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UnassignAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UnassignAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


