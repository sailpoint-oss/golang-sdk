---
id: nerm-profile-select-action
title: ProfileSelectAction
pagination_label: ProfileSelectAction
sidebar_label: ProfileSelectAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileSelectAction', 'NERMProfileSelectAction'] 
slug: /tools/sdk/go/nerm/models/profile-select-action
tags: ['SDK', 'Software Development Kit', 'ProfileSelectAction', 'NERMProfileSelectAction']
---

# ProfileSelectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewProfileSelectAction

`func NewProfileSelectAction(workflowId string, description string, ) *ProfileSelectAction`

NewProfileSelectAction instantiates a new ProfileSelectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileSelectActionWithDefaults

`func NewProfileSelectActionWithDefaults() *ProfileSelectAction`

NewProfileSelectActionWithDefaults instantiates a new ProfileSelectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *ProfileSelectAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ProfileSelectAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ProfileSelectAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *ProfileSelectAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProfileSelectAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProfileSelectAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *ProfileSelectAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ProfileSelectAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ProfileSelectAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ProfileSelectAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


