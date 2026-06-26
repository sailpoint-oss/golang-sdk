---
id: nerm-run-workflow-action-configuration-attributes
title: RunWorkflowActionConfigurationAttributes
pagination_label: RunWorkflowActionConfigurationAttributes
sidebar_label: RunWorkflowActionConfigurationAttributes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RunWorkflowActionConfigurationAttributes', 'NERMRunWorkflowActionConfigurationAttributes'] 
slug: /tools/sdk/go/nerm/models/run-workflow-action-configuration-attributes
tags: ['SDK', 'Software Development Kit', 'RunWorkflowActionConfigurationAttributes', 'NERMRunWorkflowActionConfigurationAttributes']
---

# RunWorkflowActionConfigurationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id. | [optional] 
**WorkflowId** | Pointer to **string** | the id of the workflow. | [optional] 
**WaitForCompletion** | Pointer to **bool** | If the parent workflow should wait for the child to complete. | [optional] 
**ProfileToSend** | Pointer to **string** | the profile the parent should send to the child workflow. | [optional] 
**ReturnProfile** | Pointer to **bool** | if the child workflow should return a profile. | [optional] 
**RunWorkflowActionId** | Pointer to **string** | the id of the child workflow to run. | [optional] 
**ProfileTypeId** | Pointer to **string** | the id of the profile type. | [optional] 

## Methods

### NewRunWorkflowActionConfigurationAttributes

`func NewRunWorkflowActionConfigurationAttributes() *RunWorkflowActionConfigurationAttributes`

NewRunWorkflowActionConfigurationAttributes instantiates a new RunWorkflowActionConfigurationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWorkflowActionConfigurationAttributesWithDefaults

`func NewRunWorkflowActionConfigurationAttributesWithDefaults() *RunWorkflowActionConfigurationAttributes`

NewRunWorkflowActionConfigurationAttributesWithDefaults instantiates a new RunWorkflowActionConfigurationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunWorkflowActionConfigurationAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunWorkflowActionConfigurationAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunWorkflowActionConfigurationAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RunWorkflowActionConfigurationAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *RunWorkflowActionConfigurationAttributes) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *RunWorkflowActionConfigurationAttributes) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *RunWorkflowActionConfigurationAttributes) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *RunWorkflowActionConfigurationAttributes) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWaitForCompletion

`func (o *RunWorkflowActionConfigurationAttributes) GetWaitForCompletion() bool`

GetWaitForCompletion returns the WaitForCompletion field if non-nil, zero value otherwise.

### GetWaitForCompletionOk

`func (o *RunWorkflowActionConfigurationAttributes) GetWaitForCompletionOk() (*bool, bool)`

GetWaitForCompletionOk returns a tuple with the WaitForCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletion

`func (o *RunWorkflowActionConfigurationAttributes) SetWaitForCompletion(v bool)`

SetWaitForCompletion sets WaitForCompletion field to given value.

### HasWaitForCompletion

`func (o *RunWorkflowActionConfigurationAttributes) HasWaitForCompletion() bool`

HasWaitForCompletion returns a boolean if a field has been set.

### GetProfileToSend

`func (o *RunWorkflowActionConfigurationAttributes) GetProfileToSend() string`

GetProfileToSend returns the ProfileToSend field if non-nil, zero value otherwise.

### GetProfileToSendOk

`func (o *RunWorkflowActionConfigurationAttributes) GetProfileToSendOk() (*string, bool)`

GetProfileToSendOk returns a tuple with the ProfileToSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileToSend

`func (o *RunWorkflowActionConfigurationAttributes) SetProfileToSend(v string)`

SetProfileToSend sets ProfileToSend field to given value.

### HasProfileToSend

`func (o *RunWorkflowActionConfigurationAttributes) HasProfileToSend() bool`

HasProfileToSend returns a boolean if a field has been set.

### GetReturnProfile

`func (o *RunWorkflowActionConfigurationAttributes) GetReturnProfile() bool`

GetReturnProfile returns the ReturnProfile field if non-nil, zero value otherwise.

### GetReturnProfileOk

`func (o *RunWorkflowActionConfigurationAttributes) GetReturnProfileOk() (*bool, bool)`

GetReturnProfileOk returns a tuple with the ReturnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnProfile

`func (o *RunWorkflowActionConfigurationAttributes) SetReturnProfile(v bool)`

SetReturnProfile sets ReturnProfile field to given value.

### HasReturnProfile

`func (o *RunWorkflowActionConfigurationAttributes) HasReturnProfile() bool`

HasReturnProfile returns a boolean if a field has been set.

### GetRunWorkflowActionId

`func (o *RunWorkflowActionConfigurationAttributes) GetRunWorkflowActionId() string`

GetRunWorkflowActionId returns the RunWorkflowActionId field if non-nil, zero value otherwise.

### GetRunWorkflowActionIdOk

`func (o *RunWorkflowActionConfigurationAttributes) GetRunWorkflowActionIdOk() (*string, bool)`

GetRunWorkflowActionIdOk returns a tuple with the RunWorkflowActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunWorkflowActionId

`func (o *RunWorkflowActionConfigurationAttributes) SetRunWorkflowActionId(v string)`

SetRunWorkflowActionId sets RunWorkflowActionId field to given value.

### HasRunWorkflowActionId

`func (o *RunWorkflowActionConfigurationAttributes) HasRunWorkflowActionId() bool`

HasRunWorkflowActionId returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *RunWorkflowActionConfigurationAttributes) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *RunWorkflowActionConfigurationAttributes) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *RunWorkflowActionConfigurationAttributes) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *RunWorkflowActionConfigurationAttributes) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.


