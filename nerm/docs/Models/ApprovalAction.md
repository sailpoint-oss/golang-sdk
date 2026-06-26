---
id: nerm-approval-action
title: ApprovalAction
pagination_label: ApprovalAction
sidebar_label: ApprovalAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ApprovalAction', 'NERMApprovalAction'] 
slug: /tools/sdk/go/nerm/models/approval-action
tags: ['SDK', 'Software Development Kit', 'ApprovalAction', 'NERMApprovalAction']
---

# ApprovalAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | Pointer to **string** | The description of the workflow action. | [optional] 
**PageId** | **string** | The page the workflow action should render. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**Skippable** | Pointer to **bool** | If the workflow action is skippable or not. | [optional] [default to false]
**RequiresComment** | Pointer to **bool** | If the workflow action requires a comment or not. | [optional] [default to false]

## Methods

### NewApprovalAction

`func NewApprovalAction(workflowId string, pageId string, ) *ApprovalAction`

NewApprovalAction instantiates a new ApprovalAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalActionWithDefaults

`func NewApprovalActionWithDefaults() *ApprovalAction`

NewApprovalActionWithDefaults instantiates a new ApprovalAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *ApprovalAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ApprovalAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ApprovalAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *ApprovalAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApprovalAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPageId

`func (o *ApprovalAction) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *ApprovalAction) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *ApprovalAction) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetArchived

`func (o *ApprovalAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ApprovalAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ApprovalAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ApprovalAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetSkippable

`func (o *ApprovalAction) GetSkippable() bool`

GetSkippable returns the Skippable field if non-nil, zero value otherwise.

### GetSkippableOk

`func (o *ApprovalAction) GetSkippableOk() (*bool, bool)`

GetSkippableOk returns a tuple with the Skippable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippable

`func (o *ApprovalAction) SetSkippable(v bool)`

SetSkippable sets Skippable field to given value.

### HasSkippable

`func (o *ApprovalAction) HasSkippable() bool`

HasSkippable returns a boolean if a field has been set.

### GetRequiresComment

`func (o *ApprovalAction) GetRequiresComment() bool`

GetRequiresComment returns the RequiresComment field if non-nil, zero value otherwise.

### GetRequiresCommentOk

`func (o *ApprovalAction) GetRequiresCommentOk() (*bool, bool)`

GetRequiresCommentOk returns a tuple with the RequiresComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresComment

`func (o *ApprovalAction) SetRequiresComment(v bool)`

SetRequiresComment sets RequiresComment field to given value.

### HasRequiresComment

`func (o *ApprovalAction) HasRequiresComment() bool`

HasRequiresComment returns a boolean if a field has been set.


