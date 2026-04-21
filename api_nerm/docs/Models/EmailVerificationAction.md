---
id: nerm-email-verification-action
title: EmailVerificationAction
pagination_label: EmailVerificationAction
sidebar_label: EmailVerificationAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'EmailVerificationAction', 'NERMEmailVerificationAction'] 
slug: /tools/sdk/go/nerm/models/email-verification-action
tags: ['SDK', 'Software Development Kit', 'EmailVerificationAction', 'NERMEmailVerificationAction']
---

# EmailVerificationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**EmailExpiration** | **int32** | The email expiration time, in minutes. | 
**TokenExpiration** | **int32** | The token expiration time, coordinates with token_expiration_type. | 
**TokenExpirationType** | **string** | The token expiration type, coordinates with token_expiration. | 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]

## Methods

### NewEmailVerificationAction

`func NewEmailVerificationAction(workflowId string, description string, emailExpiration int32, tokenExpiration int32, tokenExpirationType string, ) *EmailVerificationAction`

NewEmailVerificationAction instantiates a new EmailVerificationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailVerificationActionWithDefaults

`func NewEmailVerificationActionWithDefaults() *EmailVerificationAction`

NewEmailVerificationActionWithDefaults instantiates a new EmailVerificationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *EmailVerificationAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *EmailVerificationAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *EmailVerificationAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *EmailVerificationAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailVerificationAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailVerificationAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmailExpiration

`func (o *EmailVerificationAction) GetEmailExpiration() int32`

GetEmailExpiration returns the EmailExpiration field if non-nil, zero value otherwise.

### GetEmailExpirationOk

`func (o *EmailVerificationAction) GetEmailExpirationOk() (*int32, bool)`

GetEmailExpirationOk returns a tuple with the EmailExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailExpiration

`func (o *EmailVerificationAction) SetEmailExpiration(v int32)`

SetEmailExpiration sets EmailExpiration field to given value.


### GetTokenExpiration

`func (o *EmailVerificationAction) GetTokenExpiration() int32`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *EmailVerificationAction) GetTokenExpirationOk() (*int32, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *EmailVerificationAction) SetTokenExpiration(v int32)`

SetTokenExpiration sets TokenExpiration field to given value.


### GetTokenExpirationType

`func (o *EmailVerificationAction) GetTokenExpirationType() string`

GetTokenExpirationType returns the TokenExpirationType field if non-nil, zero value otherwise.

### GetTokenExpirationTypeOk

`func (o *EmailVerificationAction) GetTokenExpirationTypeOk() (*string, bool)`

GetTokenExpirationTypeOk returns a tuple with the TokenExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationType

`func (o *EmailVerificationAction) SetTokenExpirationType(v string)`

SetTokenExpirationType sets TokenExpirationType field to given value.


### GetArchived

`func (o *EmailVerificationAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *EmailVerificationAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *EmailVerificationAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *EmailVerificationAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


