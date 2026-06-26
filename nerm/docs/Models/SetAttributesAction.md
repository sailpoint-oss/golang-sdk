---
id: nerm-set-attributes-action
title: SetAttributesAction
pagination_label: SetAttributesAction
sidebar_label: SetAttributesAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SetAttributesAction', 'NERMSetAttributesAction'] 
slug: /tools/sdk/go/nerm/models/set-attributes-action
tags: ['SDK', 'Software Development Kit', 'SetAttributesAction', 'NERMSetAttributesAction']
---

# SetAttributesAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewSetAttributesAction

`func NewSetAttributesAction(workflowId string, description string, ) *SetAttributesAction`

NewSetAttributesAction instantiates a new SetAttributesAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAttributesActionWithDefaults

`func NewSetAttributesActionWithDefaults() *SetAttributesAction`

NewSetAttributesActionWithDefaults instantiates a new SetAttributesAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *SetAttributesAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *SetAttributesAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *SetAttributesAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *SetAttributesAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetAttributesAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetAttributesAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *SetAttributesAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SetAttributesAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SetAttributesAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *SetAttributesAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


