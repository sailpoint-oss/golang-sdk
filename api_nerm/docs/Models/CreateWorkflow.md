---
id: nerm-create-workflow
title: CreateWorkflow
pagination_label: CreateWorkflow
sidebar_label: CreateWorkflow
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateWorkflow', 'NERMCreateWorkflow'] 
slug: /tools/sdk/go/nerm/models/create-workflow
tags: ['SDK', 'Software Development Kit', 'CreateWorkflow', 'NERMCreateWorkflow']
---

# CreateWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypeId** | **string** | The profile type the workflow effects. | 
**Status** | **string** | Whether or not the workflow is enabled or disabled. | 
**Uid** | **string** | The user-specified identifier of the workflow. | 
**Name** | **string** | Name of the workflow | 
**DisableFailureEmailNotifications** | Pointer to **NullableBool** | When honored at runtime, suppresses failure email notifications for this workflow's sessions. | [optional] 

## Methods

### NewCreateWorkflow

`func NewCreateWorkflow(profileTypeId string, status string, uid string, name string, ) *CreateWorkflow`

NewCreateWorkflow instantiates a new CreateWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowWithDefaults

`func NewCreateWorkflowWithDefaults() *CreateWorkflow`

NewCreateWorkflowWithDefaults instantiates a new CreateWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypeId

`func (o *CreateWorkflow) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *CreateWorkflow) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *CreateWorkflow) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.


### GetStatus

`func (o *CreateWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUid

`func (o *CreateWorkflow) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CreateWorkflow) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CreateWorkflow) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *CreateWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkflow) SetName(v string)`

SetName sets Name field to given value.


### GetDisableFailureEmailNotifications

`func (o *CreateWorkflow) GetDisableFailureEmailNotifications() bool`

GetDisableFailureEmailNotifications returns the DisableFailureEmailNotifications field if non-nil, zero value otherwise.

### GetDisableFailureEmailNotificationsOk

`func (o *CreateWorkflow) GetDisableFailureEmailNotificationsOk() (*bool, bool)`

GetDisableFailureEmailNotificationsOk returns a tuple with the DisableFailureEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFailureEmailNotifications

`func (o *CreateWorkflow) SetDisableFailureEmailNotifications(v bool)`

SetDisableFailureEmailNotifications sets DisableFailureEmailNotifications field to given value.

### HasDisableFailureEmailNotifications

`func (o *CreateWorkflow) HasDisableFailureEmailNotifications() bool`

HasDisableFailureEmailNotifications returns a boolean if a field has been set.

### SetDisableFailureEmailNotificationsNil

`func (o *CreateWorkflow) SetDisableFailureEmailNotificationsNil(b bool)`

 SetDisableFailureEmailNotificationsNil sets the value for DisableFailureEmailNotifications to be an explicit nil

### UnsetDisableFailureEmailNotifications
`func (o *CreateWorkflow) UnsetDisableFailureEmailNotifications()`

UnsetDisableFailureEmailNotifications ensures that no value is present for DisableFailureEmailNotifications, not even an explicit nil

