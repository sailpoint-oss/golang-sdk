---
id: nerm-ask-security-question-action
title: AskSecurityQuestionAction
pagination_label: AskSecurityQuestionAction
sidebar_label: AskSecurityQuestionAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AskSecurityQuestionAction', 'NERMAskSecurityQuestionAction'] 
slug: /tools/sdk/go/nerm/models/ask-security-question-action
tags: ['SDK', 'Software Development Kit', 'AskSecurityQuestionAction', 'NERMAskSecurityQuestionAction']
---

# AskSecurityQuestionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**NumberOfQuestions** | **int32** | The number of questions the user should answer upon login. | 
**TokenExpiration** | **int32** | The token expiration time, coordinates with token_expiration_type. | 
**TokenExpirationType** | **string** | The token expiration type, coordinates with token_expiration. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewAskSecurityQuestionAction

`func NewAskSecurityQuestionAction(workflowId string, description string, numberOfQuestions int32, tokenExpiration int32, tokenExpirationType string, ) *AskSecurityQuestionAction`

NewAskSecurityQuestionAction instantiates a new AskSecurityQuestionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskSecurityQuestionActionWithDefaults

`func NewAskSecurityQuestionActionWithDefaults() *AskSecurityQuestionAction`

NewAskSecurityQuestionActionWithDefaults instantiates a new AskSecurityQuestionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *AskSecurityQuestionAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AskSecurityQuestionAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AskSecurityQuestionAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *AskSecurityQuestionAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AskSecurityQuestionAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AskSecurityQuestionAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNumberOfQuestions

`func (o *AskSecurityQuestionAction) GetNumberOfQuestions() int32`

GetNumberOfQuestions returns the NumberOfQuestions field if non-nil, zero value otherwise.

### GetNumberOfQuestionsOk

`func (o *AskSecurityQuestionAction) GetNumberOfQuestionsOk() (*int32, bool)`

GetNumberOfQuestionsOk returns a tuple with the NumberOfQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfQuestions

`func (o *AskSecurityQuestionAction) SetNumberOfQuestions(v int32)`

SetNumberOfQuestions sets NumberOfQuestions field to given value.


### GetTokenExpiration

`func (o *AskSecurityQuestionAction) GetTokenExpiration() int32`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *AskSecurityQuestionAction) GetTokenExpirationOk() (*int32, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *AskSecurityQuestionAction) SetTokenExpiration(v int32)`

SetTokenExpiration sets TokenExpiration field to given value.


### GetTokenExpirationType

`func (o *AskSecurityQuestionAction) GetTokenExpirationType() string`

GetTokenExpirationType returns the TokenExpirationType field if non-nil, zero value otherwise.

### GetTokenExpirationTypeOk

`func (o *AskSecurityQuestionAction) GetTokenExpirationTypeOk() (*string, bool)`

GetTokenExpirationTypeOk returns a tuple with the TokenExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationType

`func (o *AskSecurityQuestionAction) SetTokenExpirationType(v string)`

SetTokenExpirationType sets TokenExpirationType field to given value.


### GetArchived

`func (o *AskSecurityQuestionAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AskSecurityQuestionAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AskSecurityQuestionAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *AskSecurityQuestionAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


