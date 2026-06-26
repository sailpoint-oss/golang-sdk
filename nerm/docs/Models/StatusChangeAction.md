---
id: nerm-status-change-action
title: StatusChangeAction
pagination_label: StatusChangeAction
sidebar_label: StatusChangeAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'StatusChangeAction', 'NERMStatusChangeAction'] 
slug: /tools/sdk/go/nerm/models/status-change-action
tags: ['SDK', 'Software Development Kit', 'StatusChangeAction', 'NERMStatusChangeAction']
---

# StatusChangeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**NewStatus** | Pointer to **string** | The new status for the Status Change workflow action. | [optional] 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewStatusChangeAction

`func NewStatusChangeAction(workflowId string, description string, ) *StatusChangeAction`

NewStatusChangeAction instantiates a new StatusChangeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusChangeActionWithDefaults

`func NewStatusChangeActionWithDefaults() *StatusChangeAction`

NewStatusChangeActionWithDefaults instantiates a new StatusChangeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *StatusChangeAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *StatusChangeAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *StatusChangeAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *StatusChangeAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatusChangeAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatusChangeAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNewStatus

`func (o *StatusChangeAction) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *StatusChangeAction) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *StatusChangeAction) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.

### HasNewStatus

`func (o *StatusChangeAction) HasNewStatus() bool`

HasNewStatus returns a boolean if a field has been set.

### GetArchived

`func (o *StatusChangeAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *StatusChangeAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *StatusChangeAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *StatusChangeAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


