---
id: nerm-login-workflow
title: LoginWorkflow
pagination_label: LoginWorkflow
sidebar_label: LoginWorkflow
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'LoginWorkflow', 'NERMLoginWorkflow'] 
slug: /tools/sdk/go/nerm/models/login-workflow
tags: ['SDK', 'Software Development Kit', 'LoginWorkflow', 'NERMLoginWorkflow']
---

# LoginWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypeId** | **string** | The profile type the workflow effects. | 
**Status** | **string** | Whether or not the workflow is enabled or disabled. | 
**Uid** | **string** | The user-specified identifier of the workflow. | 
**Name** | **string** | Name of the workflow | 
**Options** | Pointer to [**LoginWorkflowOptions**](login-workflow-options) |  | [optional] 
**DisableFailureEmailNotifications** | Pointer to **NullableBool** | When honored at runtime, suppresses failure email notifications for this workflow's sessions. | [optional] 

## Methods

### NewLoginWorkflow

`func NewLoginWorkflow(profileTypeId string, status string, uid string, name string, ) *LoginWorkflow`

NewLoginWorkflow instantiates a new LoginWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWorkflowWithDefaults

`func NewLoginWorkflowWithDefaults() *LoginWorkflow`

NewLoginWorkflowWithDefaults instantiates a new LoginWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypeId

`func (o *LoginWorkflow) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *LoginWorkflow) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *LoginWorkflow) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.


### GetStatus

`func (o *LoginWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoginWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoginWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUid

`func (o *LoginWorkflow) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *LoginWorkflow) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *LoginWorkflow) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *LoginWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoginWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoginWorkflow) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *LoginWorkflow) GetOptions() LoginWorkflowOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *LoginWorkflow) GetOptionsOk() (*LoginWorkflowOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *LoginWorkflow) SetOptions(v LoginWorkflowOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *LoginWorkflow) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDisableFailureEmailNotifications

`func (o *LoginWorkflow) GetDisableFailureEmailNotifications() bool`

GetDisableFailureEmailNotifications returns the DisableFailureEmailNotifications field if non-nil, zero value otherwise.

### GetDisableFailureEmailNotificationsOk

`func (o *LoginWorkflow) GetDisableFailureEmailNotificationsOk() (*bool, bool)`

GetDisableFailureEmailNotificationsOk returns a tuple with the DisableFailureEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFailureEmailNotifications

`func (o *LoginWorkflow) SetDisableFailureEmailNotifications(v bool)`

SetDisableFailureEmailNotifications sets DisableFailureEmailNotifications field to given value.

### HasDisableFailureEmailNotifications

`func (o *LoginWorkflow) HasDisableFailureEmailNotifications() bool`

HasDisableFailureEmailNotifications returns a boolean if a field has been set.

### SetDisableFailureEmailNotificationsNil

`func (o *LoginWorkflow) SetDisableFailureEmailNotificationsNil(b bool)`

 SetDisableFailureEmailNotificationsNil sets the value for DisableFailureEmailNotifications to be an explicit nil

### UnsetDisableFailureEmailNotifications
`func (o *LoginWorkflow) UnsetDisableFailureEmailNotifications()`

UnsetDisableFailureEmailNotifications ensures that no value is present for DisableFailureEmailNotifications, not even an explicit nil

