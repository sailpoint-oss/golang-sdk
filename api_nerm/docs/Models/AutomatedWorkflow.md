---
id: nerm-automated-workflow
title: AutomatedWorkflow
pagination_label: AutomatedWorkflow
sidebar_label: AutomatedWorkflow
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AutomatedWorkflow', 'NERMAutomatedWorkflow'] 
slug: /tools/sdk/go/nerm/models/automated-workflow
tags: ['SDK', 'Software Development Kit', 'AutomatedWorkflow', 'NERMAutomatedWorkflow']
---

# AutomatedWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypeId** | **string** | The profile type the workflow effects. | 
**Status** | **string** | Whether or not the workflow is enabled or disabled. | 
**Uid** | **string** | The user-specified identifier of the workflow. | 
**Name** | **string** | Name of the workflow | 
**DisableFailureEmailNotifications** | Pointer to **NullableBool** | When honored at runtime, suppresses failure email notifications for this workflow's sessions. | [optional] 
**ConditionRulesAttributes** | [**[]AutomatedWorkflowConditionRulesAttributesInner**](automated-workflow-condition-rules-attributes-inner) | The ProfileTypeRule this workflow will be working with. | 

## Methods

### NewAutomatedWorkflow

`func NewAutomatedWorkflow(profileTypeId string, status string, uid string, name string, conditionRulesAttributes []AutomatedWorkflowConditionRulesAttributesInner, ) *AutomatedWorkflow`

NewAutomatedWorkflow instantiates a new AutomatedWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomatedWorkflowWithDefaults

`func NewAutomatedWorkflowWithDefaults() *AutomatedWorkflow`

NewAutomatedWorkflowWithDefaults instantiates a new AutomatedWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypeId

`func (o *AutomatedWorkflow) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *AutomatedWorkflow) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *AutomatedWorkflow) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.


### GetStatus

`func (o *AutomatedWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomatedWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomatedWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUid

`func (o *AutomatedWorkflow) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AutomatedWorkflow) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AutomatedWorkflow) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *AutomatedWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomatedWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomatedWorkflow) SetName(v string)`

SetName sets Name field to given value.


### GetDisableFailureEmailNotifications

`func (o *AutomatedWorkflow) GetDisableFailureEmailNotifications() bool`

GetDisableFailureEmailNotifications returns the DisableFailureEmailNotifications field if non-nil, zero value otherwise.

### GetDisableFailureEmailNotificationsOk

`func (o *AutomatedWorkflow) GetDisableFailureEmailNotificationsOk() (*bool, bool)`

GetDisableFailureEmailNotificationsOk returns a tuple with the DisableFailureEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFailureEmailNotifications

`func (o *AutomatedWorkflow) SetDisableFailureEmailNotifications(v bool)`

SetDisableFailureEmailNotifications sets DisableFailureEmailNotifications field to given value.

### HasDisableFailureEmailNotifications

`func (o *AutomatedWorkflow) HasDisableFailureEmailNotifications() bool`

HasDisableFailureEmailNotifications returns a boolean if a field has been set.

### SetDisableFailureEmailNotificationsNil

`func (o *AutomatedWorkflow) SetDisableFailureEmailNotificationsNil(b bool)`

 SetDisableFailureEmailNotificationsNil sets the value for DisableFailureEmailNotifications to be an explicit nil

### UnsetDisableFailureEmailNotifications
`func (o *AutomatedWorkflow) UnsetDisableFailureEmailNotifications()`

UnsetDisableFailureEmailNotifications ensures that no value is present for DisableFailureEmailNotifications, not even an explicit nil
### GetConditionRulesAttributes

`func (o *AutomatedWorkflow) GetConditionRulesAttributes() []AutomatedWorkflowConditionRulesAttributesInner`

GetConditionRulesAttributes returns the ConditionRulesAttributes field if non-nil, zero value otherwise.

### GetConditionRulesAttributesOk

`func (o *AutomatedWorkflow) GetConditionRulesAttributesOk() (*[]AutomatedWorkflowConditionRulesAttributesInner, bool)`

GetConditionRulesAttributesOk returns a tuple with the ConditionRulesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionRulesAttributes

`func (o *AutomatedWorkflow) SetConditionRulesAttributes(v []AutomatedWorkflowConditionRulesAttributesInner)`

SetConditionRulesAttributes sets ConditionRulesAttributes field to given value.



