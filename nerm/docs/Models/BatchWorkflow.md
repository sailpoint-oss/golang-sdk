---
id: nerm-batch-workflow
title: BatchWorkflow
pagination_label: BatchWorkflow
sidebar_label: BatchWorkflow
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'BatchWorkflow', 'NERMBatchWorkflow'] 
slug: /tools/sdk/go/nerm/models/batch-workflow
tags: ['SDK', 'Software Development Kit', 'BatchWorkflow', 'NERMBatchWorkflow']
---

# BatchWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypeId** | **string** | The profile type the workflow effects. | 
**Status** | **string** | Whether or not the workflow is enabled or disabled. | 
**Uid** | **string** | The user-specified identifier of the workflow. | 
**Name** | **string** | Name of the workflow | 
**Options** | Pointer to [**BatchWorkflowOptions**](batch-workflow-options) |  | [optional] 
**DisableFailureEmailNotifications** | Pointer to **NullableBool** | When honored at runtime, suppresses failure email notifications for this workflow's sessions. | [optional] 

## Methods

### NewBatchWorkflow

`func NewBatchWorkflow(profileTypeId string, status string, uid string, name string, ) *BatchWorkflow`

NewBatchWorkflow instantiates a new BatchWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWorkflowWithDefaults

`func NewBatchWorkflowWithDefaults() *BatchWorkflow`

NewBatchWorkflowWithDefaults instantiates a new BatchWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypeId

`func (o *BatchWorkflow) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *BatchWorkflow) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *BatchWorkflow) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.


### GetStatus

`func (o *BatchWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUid

`func (o *BatchWorkflow) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *BatchWorkflow) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *BatchWorkflow) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *BatchWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchWorkflow) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *BatchWorkflow) GetOptions() BatchWorkflowOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BatchWorkflow) GetOptionsOk() (*BatchWorkflowOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BatchWorkflow) SetOptions(v BatchWorkflowOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BatchWorkflow) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDisableFailureEmailNotifications

`func (o *BatchWorkflow) GetDisableFailureEmailNotifications() bool`

GetDisableFailureEmailNotifications returns the DisableFailureEmailNotifications field if non-nil, zero value otherwise.

### GetDisableFailureEmailNotificationsOk

`func (o *BatchWorkflow) GetDisableFailureEmailNotificationsOk() (*bool, bool)`

GetDisableFailureEmailNotificationsOk returns a tuple with the DisableFailureEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFailureEmailNotifications

`func (o *BatchWorkflow) SetDisableFailureEmailNotifications(v bool)`

SetDisableFailureEmailNotifications sets DisableFailureEmailNotifications field to given value.

### HasDisableFailureEmailNotifications

`func (o *BatchWorkflow) HasDisableFailureEmailNotifications() bool`

HasDisableFailureEmailNotifications returns a boolean if a field has been set.

### SetDisableFailureEmailNotificationsNil

`func (o *BatchWorkflow) SetDisableFailureEmailNotificationsNil(b bool)`

 SetDisableFailureEmailNotificationsNil sets the value for DisableFailureEmailNotifications to be an explicit nil

### UnsetDisableFailureEmailNotifications
`func (o *BatchWorkflow) UnsetDisableFailureEmailNotifications()`

UnsetDisableFailureEmailNotifications ensures that no value is present for DisableFailureEmailNotifications, not even an explicit nil

