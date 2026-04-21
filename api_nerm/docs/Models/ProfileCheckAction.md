---
id: nerm-profile-check-action
title: ProfileCheckAction
pagination_label: ProfileCheckAction
sidebar_label: ProfileCheckAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileCheckAction', 'NERMProfileCheckAction'] 
slug: /tools/sdk/go/nerm/models/profile-check-action
tags: ['SDK', 'Software Development Kit', 'ProfileCheckAction', 'NERMProfileCheckAction']
---

# ProfileCheckAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**NeAttributeIds** | Pointer to **[]string** | An array of ne_attribute_ids. | [optional] 
**HandleType** | Pointer to **string** | The handle type for what should happen if an existing profile is found. | [optional] 
**HandleId** | Pointer to **string** | The handle id.  When handle type is session, this is the session id. | [optional] 

## Methods

### NewProfileCheckAction

`func NewProfileCheckAction(workflowId string, description string, ) *ProfileCheckAction`

NewProfileCheckAction instantiates a new ProfileCheckAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileCheckActionWithDefaults

`func NewProfileCheckActionWithDefaults() *ProfileCheckAction`

NewProfileCheckActionWithDefaults instantiates a new ProfileCheckAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *ProfileCheckAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ProfileCheckAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ProfileCheckAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *ProfileCheckAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProfileCheckAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProfileCheckAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *ProfileCheckAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ProfileCheckAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ProfileCheckAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ProfileCheckAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetNeAttributeIds

`func (o *ProfileCheckAction) GetNeAttributeIds() []string`

GetNeAttributeIds returns the NeAttributeIds field if non-nil, zero value otherwise.

### GetNeAttributeIdsOk

`func (o *ProfileCheckAction) GetNeAttributeIdsOk() (*[]string, bool)`

GetNeAttributeIdsOk returns a tuple with the NeAttributeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeIds

`func (o *ProfileCheckAction) SetNeAttributeIds(v []string)`

SetNeAttributeIds sets NeAttributeIds field to given value.

### HasNeAttributeIds

`func (o *ProfileCheckAction) HasNeAttributeIds() bool`

HasNeAttributeIds returns a boolean if a field has been set.

### GetHandleType

`func (o *ProfileCheckAction) GetHandleType() string`

GetHandleType returns the HandleType field if non-nil, zero value otherwise.

### GetHandleTypeOk

`func (o *ProfileCheckAction) GetHandleTypeOk() (*string, bool)`

GetHandleTypeOk returns a tuple with the HandleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleType

`func (o *ProfileCheckAction) SetHandleType(v string)`

SetHandleType sets HandleType field to given value.

### HasHandleType

`func (o *ProfileCheckAction) HasHandleType() bool`

HasHandleType returns a boolean if a field has been set.

### GetHandleId

`func (o *ProfileCheckAction) GetHandleId() string`

GetHandleId returns the HandleId field if non-nil, zero value otherwise.

### GetHandleIdOk

`func (o *ProfileCheckAction) GetHandleIdOk() (*string, bool)`

GetHandleIdOk returns a tuple with the HandleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleId

`func (o *ProfileCheckAction) SetHandleId(v string)`

SetHandleId sets HandleId field to given value.

### HasHandleId

`func (o *ProfileCheckAction) HasHandleId() bool`

HasHandleId returns a boolean if a field has been set.


