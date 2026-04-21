---
id: nerm-auto-assign-action
title: AutoAssignAction
pagination_label: AutoAssignAction
sidebar_label: AutoAssignAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AutoAssignAction', 'NERMAutoAssignAction'] 
slug: /tools/sdk/go/nerm/models/auto-assign-action
tags: ['SDK', 'Software Development Kit', 'AutoAssignAction', 'NERMAutoAssignAction']
---

# AutoAssignAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**ContributorAttr** | Pointer to **string** | The id of the contributor attribute for contributors from another profile. If workflow_action_roles are not associated to this object, this is required. | [optional] 

## Methods

### NewAutoAssignAction

`func NewAutoAssignAction(workflowId string, description string, ) *AutoAssignAction`

NewAutoAssignAction instantiates a new AutoAssignAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoAssignActionWithDefaults

`func NewAutoAssignActionWithDefaults() *AutoAssignAction`

NewAutoAssignActionWithDefaults instantiates a new AutoAssignAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *AutoAssignAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AutoAssignAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AutoAssignAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *AutoAssignAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoAssignAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoAssignAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *AutoAssignAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AutoAssignAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AutoAssignAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *AutoAssignAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetContributorAttr

`func (o *AutoAssignAction) GetContributorAttr() string`

GetContributorAttr returns the ContributorAttr field if non-nil, zero value otherwise.

### GetContributorAttrOk

`func (o *AutoAssignAction) GetContributorAttrOk() (*string, bool)`

GetContributorAttrOk returns a tuple with the ContributorAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorAttr

`func (o *AutoAssignAction) SetContributorAttr(v string)`

SetContributorAttr sets ContributorAttr field to given value.

### HasContributorAttr

`func (o *AutoAssignAction) HasContributorAttr() bool`

HasContributorAttr returns a boolean if a field has been set.


