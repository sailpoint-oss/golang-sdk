---
id: nerm-contributors-action
title: ContributorsAction
pagination_label: ContributorsAction
sidebar_label: ContributorsAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ContributorsAction', 'NERMContributorsAction'] 
slug: /tools/sdk/go/nerm/models/contributors-action
tags: ['SDK', 'Software Development Kit', 'ContributorsAction', 'NERMContributorsAction']
---

# ContributorsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Owner** | Pointer to **string** | If the assignment option should be an owner.  Owner, Contributors, or Roles must be chosen. | [optional] 
**Contributors** | Pointer to **string** | If the assignment option should be contributors. Owner, Contributors, or Roles must be chosen. | [optional] 
**Roles** | Pointer to **string** | If the assignment option should be roles. Owner, Contributors, or Roles must be chosen. | [optional] 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewContributorsAction

`func NewContributorsAction(workflowId string, description string, ) *ContributorsAction`

NewContributorsAction instantiates a new ContributorsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributorsActionWithDefaults

`func NewContributorsActionWithDefaults() *ContributorsAction`

NewContributorsActionWithDefaults instantiates a new ContributorsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *ContributorsAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ContributorsAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ContributorsAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *ContributorsAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContributorsAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContributorsAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOwner

`func (o *ContributorsAction) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ContributorsAction) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ContributorsAction) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ContributorsAction) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetContributors

`func (o *ContributorsAction) GetContributors() string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *ContributorsAction) GetContributorsOk() (*string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *ContributorsAction) SetContributors(v string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *ContributorsAction) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetRoles

`func (o *ContributorsAction) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ContributorsAction) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ContributorsAction) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ContributorsAction) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetArchived

`func (o *ContributorsAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ContributorsAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ContributorsAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ContributorsAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


