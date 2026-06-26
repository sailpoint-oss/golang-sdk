---
id: nerm-invitation-action
title: InvitationAction
pagination_label: InvitationAction
sidebar_label: InvitationAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'InvitationAction', 'NERMInvitationAction'] 
slug: /tools/sdk/go/nerm/models/invitation-action
tags: ['SDK', 'Software Development Kit', 'InvitationAction', 'NERMInvitationAction']
---

# InvitationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**ConfigurationAttributes** | Pointer to [**InvitationActionConfigurationAttributes**](invitation-action-configuration-attributes) |  | [optional] 
**WorkflowActionEmailAttributes** | Pointer to [**InvitationActionWorkflowActionEmailAttributes**](invitation-action-workflow-action-email-attributes) |  | [optional] 

## Methods

### NewInvitationAction

`func NewInvitationAction(workflowId string, description string, ) *InvitationAction`

NewInvitationAction instantiates a new InvitationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationActionWithDefaults

`func NewInvitationActionWithDefaults() *InvitationAction`

NewInvitationActionWithDefaults instantiates a new InvitationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *InvitationAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *InvitationAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *InvitationAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *InvitationAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvitationAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvitationAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *InvitationAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *InvitationAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *InvitationAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *InvitationAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetConfigurationAttributes

`func (o *InvitationAction) GetConfigurationAttributes() InvitationActionConfigurationAttributes`

GetConfigurationAttributes returns the ConfigurationAttributes field if non-nil, zero value otherwise.

### GetConfigurationAttributesOk

`func (o *InvitationAction) GetConfigurationAttributesOk() (*InvitationActionConfigurationAttributes, bool)`

GetConfigurationAttributesOk returns a tuple with the ConfigurationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationAttributes

`func (o *InvitationAction) SetConfigurationAttributes(v InvitationActionConfigurationAttributes)`

SetConfigurationAttributes sets ConfigurationAttributes field to given value.

### HasConfigurationAttributes

`func (o *InvitationAction) HasConfigurationAttributes() bool`

HasConfigurationAttributes returns a boolean if a field has been set.

### GetWorkflowActionEmailAttributes

`func (o *InvitationAction) GetWorkflowActionEmailAttributes() InvitationActionWorkflowActionEmailAttributes`

GetWorkflowActionEmailAttributes returns the WorkflowActionEmailAttributes field if non-nil, zero value otherwise.

### GetWorkflowActionEmailAttributesOk

`func (o *InvitationAction) GetWorkflowActionEmailAttributesOk() (*InvitationActionWorkflowActionEmailAttributes, bool)`

GetWorkflowActionEmailAttributesOk returns a tuple with the WorkflowActionEmailAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionEmailAttributes

`func (o *InvitationAction) SetWorkflowActionEmailAttributes(v InvitationActionWorkflowActionEmailAttributes)`

SetWorkflowActionEmailAttributes sets WorkflowActionEmailAttributes field to given value.

### HasWorkflowActionEmailAttributes

`func (o *InvitationAction) HasWorkflowActionEmailAttributes() bool`

HasWorkflowActionEmailAttributes returns a boolean if a field has been set.


