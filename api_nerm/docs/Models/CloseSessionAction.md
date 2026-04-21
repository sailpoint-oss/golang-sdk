---
id: nerm-close-session-action
title: CloseSessionAction
pagination_label: CloseSessionAction
sidebar_label: CloseSessionAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CloseSessionAction', 'NERMCloseSessionAction'] 
slug: /tools/sdk/go/nerm/models/close-session-action
tags: ['SDK', 'Software Development Kit', 'CloseSessionAction', 'NERMCloseSessionAction']
---

# CloseSessionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewCloseSessionAction

`func NewCloseSessionAction(workflowId string, description string, ) *CloseSessionAction`

NewCloseSessionAction instantiates a new CloseSessionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloseSessionActionWithDefaults

`func NewCloseSessionActionWithDefaults() *CloseSessionAction`

NewCloseSessionActionWithDefaults instantiates a new CloseSessionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *CloseSessionAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CloseSessionAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CloseSessionAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *CloseSessionAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloseSessionAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloseSessionAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *CloseSessionAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CloseSessionAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CloseSessionAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *CloseSessionAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


