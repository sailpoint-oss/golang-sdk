---
id: nerm-workflow-session1
title: WorkflowSession1
pagination_label: WorkflowSession1
sidebar_label: WorkflowSession1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowSession1', 'NERMWorkflowSession1'] 
slug: /tools/sdk/go/nerm/models/workflow-session1
tags: ['SDK', 'Software Development Kit', 'WorkflowSession1', 'NERMWorkflowSession1']
---

# WorkflowSession1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow id. | 
**RequesterId** | **string** | The requester's id. | 
**RequesterType** | **string** | The requester type. | 
**ProfileId** | Pointer to **string** | The profile this workflow session will be working with. Only Applicable for Update workflows | [optional] 
**ProfileIds** | Pointer to **[]string** | The profiles this workflow session will be working with. Only Applicable for Batch workflows | [optional] 
**Attributes** | Pointer to **map[string]string** | The attributes associated with the workflow session. | [optional] 

## Methods

### NewWorkflowSession1

`func NewWorkflowSession1(workflowId string, requesterId string, requesterType string, ) *WorkflowSession1`

NewWorkflowSession1 instantiates a new WorkflowSession1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSession1WithDefaults

`func NewWorkflowSession1WithDefaults() *WorkflowSession1`

NewWorkflowSession1WithDefaults instantiates a new WorkflowSession1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *WorkflowSession1) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowSession1) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowSession1) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetRequesterId

`func (o *WorkflowSession1) GetRequesterId() string`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *WorkflowSession1) GetRequesterIdOk() (*string, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *WorkflowSession1) SetRequesterId(v string)`

SetRequesterId sets RequesterId field to given value.


### GetRequesterType

`func (o *WorkflowSession1) GetRequesterType() string`

GetRequesterType returns the RequesterType field if non-nil, zero value otherwise.

### GetRequesterTypeOk

`func (o *WorkflowSession1) GetRequesterTypeOk() (*string, bool)`

GetRequesterTypeOk returns a tuple with the RequesterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterType

`func (o *WorkflowSession1) SetRequesterType(v string)`

SetRequesterType sets RequesterType field to given value.


### GetProfileId

`func (o *WorkflowSession1) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *WorkflowSession1) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *WorkflowSession1) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *WorkflowSession1) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileIds

`func (o *WorkflowSession1) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *WorkflowSession1) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *WorkflowSession1) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.

### HasProfileIds

`func (o *WorkflowSession1) HasProfileIds() bool`

HasProfileIds returns a boolean if a field has been set.

### GetAttributes

`func (o *WorkflowSession1) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowSession1) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowSession1) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WorkflowSession1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


