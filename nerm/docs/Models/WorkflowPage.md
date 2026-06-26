---
id: nerm-workflow-page
title: WorkflowPage
pagination_label: WorkflowPage
sidebar_label: WorkflowPage
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowPage', 'NERMWorkflowPage'] 
slug: /tools/sdk/go/nerm/models/workflow-page
tags: ['SDK', 'Software Development Kit', 'WorkflowPage', 'NERMWorkflowPage']
---

# WorkflowPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The user-specified identifier of the page | [optional] 
**Description** | Pointer to **string** | The description of the page | [optional] 
**Name** | Pointer to **string** | The name of the page | [optional] 
**Archived** | Pointer to **bool** | Determines whether the page is archived | [optional] 
**Id** | Pointer to **string** | The id of the form | [optional] [readonly] 

## Methods

### NewWorkflowPage

`func NewWorkflowPage() *WorkflowPage`

NewWorkflowPage instantiates a new WorkflowPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPageWithDefaults

`func NewWorkflowPageWithDefaults() *WorkflowPage`

NewWorkflowPageWithDefaults instantiates a new WorkflowPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *WorkflowPage) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *WorkflowPage) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *WorkflowPage) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *WorkflowPage) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowPage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowPage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowPage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowPage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowPage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *WorkflowPage) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *WorkflowPage) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *WorkflowPage) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *WorkflowPage) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetId

`func (o *WorkflowPage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowPage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowPage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowPage) HasId() bool`

HasId returns a boolean if a field has been set.


