---
id: nerm-duplicate-prevention-action
title: DuplicatePreventionAction
pagination_label: DuplicatePreventionAction
sidebar_label: DuplicatePreventionAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DuplicatePreventionAction', 'NERMDuplicatePreventionAction'] 
slug: /tools/sdk/go/nerm/models/duplicate-prevention-action
tags: ['SDK', 'Software Development Kit', 'DuplicatePreventionAction', 'NERMDuplicatePreventionAction']
---

# DuplicatePreventionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**SearchScope** | **string** | The search scope of the profiles to check against. | 
**NeAttributeIds** | Pointer to **[]string** | An array of ne_attribute_ids. | [optional] 
**HandleId** | Pointer to **string** | The handle id. | [optional] 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewDuplicatePreventionAction

`func NewDuplicatePreventionAction(workflowId string, description string, searchScope string, ) *DuplicatePreventionAction`

NewDuplicatePreventionAction instantiates a new DuplicatePreventionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicatePreventionActionWithDefaults

`func NewDuplicatePreventionActionWithDefaults() *DuplicatePreventionAction`

NewDuplicatePreventionActionWithDefaults instantiates a new DuplicatePreventionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *DuplicatePreventionAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *DuplicatePreventionAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *DuplicatePreventionAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *DuplicatePreventionAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DuplicatePreventionAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DuplicatePreventionAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSearchScope

`func (o *DuplicatePreventionAction) GetSearchScope() string`

GetSearchScope returns the SearchScope field if non-nil, zero value otherwise.

### GetSearchScopeOk

`func (o *DuplicatePreventionAction) GetSearchScopeOk() (*string, bool)`

GetSearchScopeOk returns a tuple with the SearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchScope

`func (o *DuplicatePreventionAction) SetSearchScope(v string)`

SetSearchScope sets SearchScope field to given value.


### GetNeAttributeIds

`func (o *DuplicatePreventionAction) GetNeAttributeIds() []string`

GetNeAttributeIds returns the NeAttributeIds field if non-nil, zero value otherwise.

### GetNeAttributeIdsOk

`func (o *DuplicatePreventionAction) GetNeAttributeIdsOk() (*[]string, bool)`

GetNeAttributeIdsOk returns a tuple with the NeAttributeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeIds

`func (o *DuplicatePreventionAction) SetNeAttributeIds(v []string)`

SetNeAttributeIds sets NeAttributeIds field to given value.

### HasNeAttributeIds

`func (o *DuplicatePreventionAction) HasNeAttributeIds() bool`

HasNeAttributeIds returns a boolean if a field has been set.

### GetHandleId

`func (o *DuplicatePreventionAction) GetHandleId() string`

GetHandleId returns the HandleId field if non-nil, zero value otherwise.

### GetHandleIdOk

`func (o *DuplicatePreventionAction) GetHandleIdOk() (*string, bool)`

GetHandleIdOk returns a tuple with the HandleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleId

`func (o *DuplicatePreventionAction) SetHandleId(v string)`

SetHandleId sets HandleId field to given value.

### HasHandleId

`func (o *DuplicatePreventionAction) HasHandleId() bool`

HasHandleId returns a boolean if a field has been set.

### GetArchived

`func (o *DuplicatePreventionAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *DuplicatePreventionAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *DuplicatePreventionAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *DuplicatePreventionAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


