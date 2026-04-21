---
id: nerm-workflow-action
title: WorkflowAction
pagination_label: WorkflowAction
sidebar_label: WorkflowAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowAction', 'NERMWorkflowAction'] 
slug: /tools/sdk/go/nerm/models/workflow-action
tags: ['SDK', 'Software Development Kit', 'WorkflowAction', 'NERMWorkflowAction']
---

# WorkflowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | Pointer to **string** | The description of the workflow action. | [optional] 
**PageId** | **string** | The page the workflow action should render. | 
**AddRequesterAsOwner** | Pointer to **bool** | If the requester should be added as the owner of the profile to be created. | [optional] [default to true]
**EmailAttributeId** | Pointer to **string** | The attribute storing the email address for the workflow action. | [optional] 
**EmailAddresses** | Pointer to **[]string** | The email addresses for the workflow action. | [optional] 
**NewStatus** | Pointer to **string** | The new status for the Status Change workflow action. | [optional] 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**Skippable** | Pointer to **bool** | If the workflow action is skippable or not. | [optional] [default to false]
**RequiresComment** | Pointer to **bool** | If the workflow action requires a comment or not. | [optional] [default to false]

## Methods

### NewWorkflowAction

`func NewWorkflowAction(workflowId string, pageId string, ) *WorkflowAction`

NewWorkflowAction instantiates a new WorkflowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionWithDefaults

`func NewWorkflowActionWithDefaults() *WorkflowAction`

NewWorkflowActionWithDefaults instantiates a new WorkflowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *WorkflowAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *WorkflowAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPageId

`func (o *WorkflowAction) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *WorkflowAction) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *WorkflowAction) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetAddRequesterAsOwner

`func (o *WorkflowAction) GetAddRequesterAsOwner() bool`

GetAddRequesterAsOwner returns the AddRequesterAsOwner field if non-nil, zero value otherwise.

### GetAddRequesterAsOwnerOk

`func (o *WorkflowAction) GetAddRequesterAsOwnerOk() (*bool, bool)`

GetAddRequesterAsOwnerOk returns a tuple with the AddRequesterAsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddRequesterAsOwner

`func (o *WorkflowAction) SetAddRequesterAsOwner(v bool)`

SetAddRequesterAsOwner sets AddRequesterAsOwner field to given value.

### HasAddRequesterAsOwner

`func (o *WorkflowAction) HasAddRequesterAsOwner() bool`

HasAddRequesterAsOwner returns a boolean if a field has been set.

### GetEmailAttributeId

`func (o *WorkflowAction) GetEmailAttributeId() string`

GetEmailAttributeId returns the EmailAttributeId field if non-nil, zero value otherwise.

### GetEmailAttributeIdOk

`func (o *WorkflowAction) GetEmailAttributeIdOk() (*string, bool)`

GetEmailAttributeIdOk returns a tuple with the EmailAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttributeId

`func (o *WorkflowAction) SetEmailAttributeId(v string)`

SetEmailAttributeId sets EmailAttributeId field to given value.

### HasEmailAttributeId

`func (o *WorkflowAction) HasEmailAttributeId() bool`

HasEmailAttributeId returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *WorkflowAction) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *WorkflowAction) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *WorkflowAction) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *WorkflowAction) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetNewStatus

`func (o *WorkflowAction) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *WorkflowAction) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *WorkflowAction) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.

### HasNewStatus

`func (o *WorkflowAction) HasNewStatus() bool`

HasNewStatus returns a boolean if a field has been set.

### GetArchived

`func (o *WorkflowAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *WorkflowAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *WorkflowAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *WorkflowAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetSkippable

`func (o *WorkflowAction) GetSkippable() bool`

GetSkippable returns the Skippable field if non-nil, zero value otherwise.

### GetSkippableOk

`func (o *WorkflowAction) GetSkippableOk() (*bool, bool)`

GetSkippableOk returns a tuple with the Skippable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippable

`func (o *WorkflowAction) SetSkippable(v bool)`

SetSkippable sets Skippable field to given value.

### HasSkippable

`func (o *WorkflowAction) HasSkippable() bool`

HasSkippable returns a boolean if a field has been set.

### GetRequiresComment

`func (o *WorkflowAction) GetRequiresComment() bool`

GetRequiresComment returns the RequiresComment field if non-nil, zero value otherwise.

### GetRequiresCommentOk

`func (o *WorkflowAction) GetRequiresCommentOk() (*bool, bool)`

GetRequiresCommentOk returns a tuple with the RequiresComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresComment

`func (o *WorkflowAction) SetRequiresComment(v bool)`

SetRequiresComment sets RequiresComment field to given value.

### HasRequiresComment

`func (o *WorkflowAction) HasRequiresComment() bool`

HasRequiresComment returns a boolean if a field has been set.


