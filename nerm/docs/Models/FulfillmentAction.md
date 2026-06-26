---
id: nerm-fulfillment-action
title: FulfillmentAction
pagination_label: FulfillmentAction
sidebar_label: FulfillmentAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'FulfillmentAction', 'NERMFulfillmentAction'] 
slug: /tools/sdk/go/nerm/models/fulfillment-action
tags: ['SDK', 'Software Development Kit', 'FulfillmentAction', 'NERMFulfillmentAction']
---

# FulfillmentAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**PageId** | Pointer to **string** | The page the workflow action should render. | [optional] 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**RequiresComment** | Pointer to **bool** | If the workflow action requires a comment or not. | [optional] [default to false]

## Methods

### NewFulfillmentAction

`func NewFulfillmentAction(workflowId string, description string, ) *FulfillmentAction`

NewFulfillmentAction instantiates a new FulfillmentAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentActionWithDefaults

`func NewFulfillmentActionWithDefaults() *FulfillmentAction`

NewFulfillmentActionWithDefaults instantiates a new FulfillmentAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *FulfillmentAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *FulfillmentAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *FulfillmentAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *FulfillmentAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FulfillmentAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FulfillmentAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPageId

`func (o *FulfillmentAction) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *FulfillmentAction) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *FulfillmentAction) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *FulfillmentAction) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetArchived

`func (o *FulfillmentAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FulfillmentAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FulfillmentAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *FulfillmentAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetRequiresComment

`func (o *FulfillmentAction) GetRequiresComment() bool`

GetRequiresComment returns the RequiresComment field if non-nil, zero value otherwise.

### GetRequiresCommentOk

`func (o *FulfillmentAction) GetRequiresCommentOk() (*bool, bool)`

GetRequiresCommentOk returns a tuple with the RequiresComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresComment

`func (o *FulfillmentAction) SetRequiresComment(v bool)`

SetRequiresComment sets RequiresComment field to given value.

### HasRequiresComment

`func (o *FulfillmentAction) HasRequiresComment() bool`

HasRequiresComment returns a boolean if a field has been set.


