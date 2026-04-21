---
id: nerm-set-security-question-action
title: SetSecurityQuestionAction
pagination_label: SetSecurityQuestionAction
sidebar_label: SetSecurityQuestionAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SetSecurityQuestionAction', 'NERMSetSecurityQuestionAction'] 
slug: /tools/sdk/go/nerm/models/set-security-question-action
tags: ['SDK', 'Software Development Kit', 'SetSecurityQuestionAction', 'NERMSetSecurityQuestionAction']
---

# SetSecurityQuestionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**NumberOfQuestions** | **int32** | The number of questions the user should answer upon login. | 

## Methods

### NewSetSecurityQuestionAction

`func NewSetSecurityQuestionAction(workflowId string, description string, numberOfQuestions int32, ) *SetSecurityQuestionAction`

NewSetSecurityQuestionAction instantiates a new SetSecurityQuestionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetSecurityQuestionActionWithDefaults

`func NewSetSecurityQuestionActionWithDefaults() *SetSecurityQuestionAction`

NewSetSecurityQuestionActionWithDefaults instantiates a new SetSecurityQuestionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *SetSecurityQuestionAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *SetSecurityQuestionAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *SetSecurityQuestionAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *SetSecurityQuestionAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetSecurityQuestionAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetSecurityQuestionAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *SetSecurityQuestionAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SetSecurityQuestionAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SetSecurityQuestionAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *SetSecurityQuestionAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetNumberOfQuestions

`func (o *SetSecurityQuestionAction) GetNumberOfQuestions() int32`

GetNumberOfQuestions returns the NumberOfQuestions field if non-nil, zero value otherwise.

### GetNumberOfQuestionsOk

`func (o *SetSecurityQuestionAction) GetNumberOfQuestionsOk() (*int32, bool)`

GetNumberOfQuestionsOk returns a tuple with the NumberOfQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfQuestions

`func (o *SetSecurityQuestionAction) SetNumberOfQuestions(v int32)`

SetNumberOfQuestions sets NumberOfQuestions field to given value.



