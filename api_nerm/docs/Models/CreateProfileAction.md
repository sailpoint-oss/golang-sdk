---
id: nerm-create-profile-action
title: CreateProfileAction
pagination_label: CreateProfileAction
sidebar_label: CreateProfileAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateProfileAction', 'NERMCreateProfileAction'] 
slug: /tools/sdk/go/nerm/models/create-profile-action
tags: ['SDK', 'Software Development Kit', 'CreateProfileAction', 'NERMCreateProfileAction']
---

# CreateProfileAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**AddRequesterAsOwner** | Pointer to **bool** | If the requester should be added as the owner of the profile to be created. | [optional] [default to true]
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewCreateProfileAction

`func NewCreateProfileAction(workflowId string, description string, ) *CreateProfileAction`

NewCreateProfileAction instantiates a new CreateProfileAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfileActionWithDefaults

`func NewCreateProfileActionWithDefaults() *CreateProfileAction`

NewCreateProfileActionWithDefaults instantiates a new CreateProfileAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *CreateProfileAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CreateProfileAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CreateProfileAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *CreateProfileAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProfileAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProfileAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAddRequesterAsOwner

`func (o *CreateProfileAction) GetAddRequesterAsOwner() bool`

GetAddRequesterAsOwner returns the AddRequesterAsOwner field if non-nil, zero value otherwise.

### GetAddRequesterAsOwnerOk

`func (o *CreateProfileAction) GetAddRequesterAsOwnerOk() (*bool, bool)`

GetAddRequesterAsOwnerOk returns a tuple with the AddRequesterAsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRequesterAsOwner

`func (o *CreateProfileAction) SetAddRequesterAsOwner(v bool)`

SetAddRequesterAsOwner sets AddRequesterAsOwner field to given value.

### HasAddRequesterAsOwner

`func (o *CreateProfileAction) HasAddRequesterAsOwner() bool`

HasAddRequesterAsOwner returns a boolean if a field has been set.

### GetArchived

`func (o *CreateProfileAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CreateProfileAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CreateProfileAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *CreateProfileAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


