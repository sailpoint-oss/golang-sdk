---
id: nerm-workflow-session
title: WorkflowSession
pagination_label: WorkflowSession
sidebar_label: WorkflowSession
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowSession', 'NERMWorkflowSession'] 
slug: /tools/sdk/go/nerm/models/workflow-session
tags: ['SDK', 'Software Development Kit', 'WorkflowSession', 'NERMWorkflowSession']
---

# WorkflowSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The objects ID. | [optional] [readonly] 
**Uid** | Pointer to **string** | The objects UID. | [optional] [readonly] 
**WorkflowId** | Pointer to **string** | The workflow id. | [optional] 
**RequesterId** | Pointer to **string** | The requester's id. | [optional] 
**RequesterType** | Pointer to **string** | The requester type. | [optional] 
**ProfileId** | Pointer to **string** | The profile this workflow session will be working with. Only Applicable for Update workflows | [optional] 
**ProfileIds** | Pointer to **[]string** | The profiles this workflow session will be working with. Only Applicable for Batch workflows | [optional] 
**Status** | Pointer to **string** | The status of the workflow session. | [optional] 
**Attributes** | Pointer to **map[string]string** | The attributes asscoiated with the workflow session. | [optional] 

## Methods

### NewWorkflowSession

`func NewWorkflowSession() *WorkflowSession`

NewWorkflowSession instantiates a new WorkflowSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSessionWithDefaults

`func NewWorkflowSessionWithDefaults() *WorkflowSession`

NewWorkflowSessionWithDefaults instantiates a new WorkflowSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *WorkflowSession) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *WorkflowSession) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *WorkflowSession) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *WorkflowSession) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowSession) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowSession) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowSession) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowSession) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetRequesterId

`func (o *WorkflowSession) GetRequesterId() string`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *WorkflowSession) GetRequesterIdOk() (*string, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *WorkflowSession) SetRequesterId(v string)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *WorkflowSession) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### GetRequesterType

`func (o *WorkflowSession) GetRequesterType() string`

GetRequesterType returns the RequesterType field if non-nil, zero value otherwise.

### GetRequesterTypeOk

`func (o *WorkflowSession) GetRequesterTypeOk() (*string, bool)`

GetRequesterTypeOk returns a tuple with the RequesterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterType

`func (o *WorkflowSession) SetRequesterType(v string)`

SetRequesterType sets RequesterType field to given value.

### HasRequesterType

`func (o *WorkflowSession) HasRequesterType() bool`

HasRequesterType returns a boolean if a field has been set.

### GetProfileId

`func (o *WorkflowSession) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *WorkflowSession) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *WorkflowSession) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *WorkflowSession) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileIds

`func (o *WorkflowSession) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *WorkflowSession) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *WorkflowSession) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.

### HasProfileIds

`func (o *WorkflowSession) HasProfileIds() bool`

HasProfileIds returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttributes

`func (o *WorkflowSession) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowSession) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowSession) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkflowSession) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


