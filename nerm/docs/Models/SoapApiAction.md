---
id: nerm-soap-api-action
title: SoapApiAction
pagination_label: SoapApiAction
sidebar_label: SoapApiAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SoapApiAction', 'NERMSoapApiAction'] 
slug: /tools/sdk/go/nerm/models/soap-api-action
tags: ['SDK', 'Software Development Kit', 'SoapApiAction', 'NERMSoapApiAction']
---

# SoapApiAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewSoapApiAction

`func NewSoapApiAction(workflowId string, description string, ) *SoapApiAction`

NewSoapApiAction instantiates a new SoapApiAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoapApiActionWithDefaults

`func NewSoapApiActionWithDefaults() *SoapApiAction`

NewSoapApiActionWithDefaults instantiates a new SoapApiAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *SoapApiAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *SoapApiAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *SoapApiAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *SoapApiAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoapApiAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoapApiAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *SoapApiAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SoapApiAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SoapApiAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *SoapApiAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


