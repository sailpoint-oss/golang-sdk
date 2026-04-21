---
id: nerm-batch-update-action
title: BatchUpdateAction
pagination_label: BatchUpdateAction
sidebar_label: BatchUpdateAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'BatchUpdateAction', 'NERMBatchUpdateAction'] 
slug: /tools/sdk/go/nerm/models/batch-update-action
tags: ['SDK', 'Software Development Kit', 'BatchUpdateAction', 'NERMBatchUpdateAction']
---

# BatchUpdateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewBatchUpdateAction

`func NewBatchUpdateAction(workflowId string, description string, ) *BatchUpdateAction`

NewBatchUpdateAction instantiates a new BatchUpdateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateActionWithDefaults

`func NewBatchUpdateActionWithDefaults() *BatchUpdateAction`

NewBatchUpdateActionWithDefaults instantiates a new BatchUpdateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *BatchUpdateAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *BatchUpdateAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *BatchUpdateAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *BatchUpdateAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BatchUpdateAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BatchUpdateAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *BatchUpdateAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *BatchUpdateAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *BatchUpdateAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *BatchUpdateAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


