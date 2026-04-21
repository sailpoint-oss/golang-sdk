---
id: nerm-rest-api-action
title: RestApiAction
pagination_label: RestApiAction
sidebar_label: RestApiAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RestApiAction', 'NERMRestApiAction'] 
slug: /tools/sdk/go/nerm/models/rest-api-action
tags: ['SDK', 'Software Development Kit', 'RestApiAction', 'NERMRestApiAction']
---

# RestApiAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewRestApiAction

`func NewRestApiAction(workflowId string, description string, ) *RestApiAction`

NewRestApiAction instantiates a new RestApiAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestApiActionWithDefaults

`func NewRestApiActionWithDefaults() *RestApiAction`

NewRestApiActionWithDefaults instantiates a new RestApiAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *RestApiAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *RestApiAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *RestApiAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *RestApiAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestApiAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestApiAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *RestApiAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *RestApiAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *RestApiAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *RestApiAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


