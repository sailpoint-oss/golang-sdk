---
id: nerm-update-workflow
title: UpdateWorkflow
pagination_label: UpdateWorkflow
sidebar_label: UpdateWorkflow
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UpdateWorkflow', 'NERMUpdateWorkflow'] 
slug: /tools/sdk/go/nerm/models/update-workflow
tags: ['SDK', 'Software Development Kit', 'UpdateWorkflow', 'NERMUpdateWorkflow']
---

# UpdateWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypeId** | **string** | The profile type the workflow effects. | 
**Status** | **string** | Whether or not the workflow is enabled or disabled. | 
**Uid** | **string** | The user-specified identifier of the workflow. | 
**Name** | **string** | Name of the workflow | 
**ProfileStatus** | **string** | The status of the profiles the workflow will effect. | 
**DisableFailureEmailNotifications** | Pointer to **NullableBool** | When honored at runtime, suppresses failure email notifications for this workflow's sessions. | [optional] 

## Methods

### NewUpdateWorkflow

`func NewUpdateWorkflow(profileTypeId string, status string, uid string, name string, profileStatus string, ) *UpdateWorkflow`

NewUpdateWorkflow instantiates a new UpdateWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkflowWithDefaults

`func NewUpdateWorkflowWithDefaults() *UpdateWorkflow`

NewUpdateWorkflowWithDefaults instantiates a new UpdateWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypeId

`func (o *UpdateWorkflow) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *UpdateWorkflow) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *UpdateWorkflow) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.


### GetStatus

`func (o *UpdateWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUid

`func (o *UpdateWorkflow) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UpdateWorkflow) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UpdateWorkflow) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *UpdateWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWorkflow) SetName(v string)`

SetName sets Name field to given value.


### GetProfileStatus

`func (o *UpdateWorkflow) GetProfileStatus() string`

GetProfileStatus returns the ProfileStatus field if non-nil, zero value otherwise.

### GetProfileStatusOk

`func (o *UpdateWorkflow) GetProfileStatusOk() (*string, bool)`

GetProfileStatusOk returns a tuple with the ProfileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileStatus

`func (o *UpdateWorkflow) SetProfileStatus(v string)`

SetProfileStatus sets ProfileStatus field to given value.


### GetDisableFailureEmailNotifications

`func (o *UpdateWorkflow) GetDisableFailureEmailNotifications() bool`

GetDisableFailureEmailNotifications returns the DisableFailureEmailNotifications field if non-nil, zero value otherwise.

### GetDisableFailureEmailNotificationsOk

`func (o *UpdateWorkflow) GetDisableFailureEmailNotificationsOk() (*bool, bool)`

GetDisableFailureEmailNotificationsOk returns a tuple with the DisableFailureEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFailureEmailNotifications

`func (o *UpdateWorkflow) SetDisableFailureEmailNotifications(v bool)`

SetDisableFailureEmailNotifications sets DisableFailureEmailNotifications field to given value.

### HasDisableFailureEmailNotifications

`func (o *UpdateWorkflow) HasDisableFailureEmailNotifications() bool`

HasDisableFailureEmailNotifications returns a boolean if a field has been set.

### SetDisableFailureEmailNotificationsNil

`func (o *UpdateWorkflow) SetDisableFailureEmailNotificationsNil(b bool)`

 SetDisableFailureEmailNotificationsNil sets the value for DisableFailureEmailNotifications to be an explicit nil

### UnsetDisableFailureEmailNotifications
`func (o *UpdateWorkflow) UnsetDisableFailureEmailNotifications()`

UnsetDisableFailureEmailNotifications ensures that no value is present for DisableFailureEmailNotifications, not even an explicit nil

