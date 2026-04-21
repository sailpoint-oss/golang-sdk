---
id: nerm-run-workflow-action
title: RunWorkflowAction
pagination_label: RunWorkflowAction
sidebar_label: RunWorkflowAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RunWorkflowAction', 'NERMRunWorkflowAction'] 
slug: /tools/sdk/go/nerm/models/run-workflow-action
tags: ['SDK', 'Software Development Kit', 'RunWorkflowAction', 'NERMRunWorkflowAction']
---

# RunWorkflowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**ConfigurationAttributes** | Pointer to [**RunWorkflowActionConfigurationAttributes**](run-workflow-action-configuration-attributes) |  | [optional] 

## Methods

### NewRunWorkflowAction

`func NewRunWorkflowAction(workflowId string, description string, ) *RunWorkflowAction`

NewRunWorkflowAction instantiates a new RunWorkflowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWorkflowActionWithDefaults

`func NewRunWorkflowActionWithDefaults() *RunWorkflowAction`

NewRunWorkflowActionWithDefaults instantiates a new RunWorkflowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *RunWorkflowAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *RunWorkflowAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *RunWorkflowAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *RunWorkflowAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RunWorkflowAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RunWorkflowAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *RunWorkflowAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *RunWorkflowAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *RunWorkflowAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *RunWorkflowAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetConfigurationAttributes

`func (o *RunWorkflowAction) GetConfigurationAttributes() RunWorkflowActionConfigurationAttributes`

GetConfigurationAttributes returns the ConfigurationAttributes field if non-nil, zero value otherwise.

### GetConfigurationAttributesOk

`func (o *RunWorkflowAction) GetConfigurationAttributesOk() (*RunWorkflowActionConfigurationAttributes, bool)`

GetConfigurationAttributesOk returns a tuple with the ConfigurationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationAttributes

`func (o *RunWorkflowAction) SetConfigurationAttributes(v RunWorkflowActionConfigurationAttributes)`

SetConfigurationAttributes sets ConfigurationAttributes field to given value.

### HasConfigurationAttributes

`func (o *RunWorkflowAction) HasConfigurationAttributes() bool`

HasConfigurationAttributes returns a boolean if a field has been set.


