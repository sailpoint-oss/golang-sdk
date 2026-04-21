---
id: nerm-create-notification-action-request
title: CreateNotificationActionRequest
pagination_label: CreateNotificationActionRequest
sidebar_label: CreateNotificationActionRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateNotificationActionRequest', 'NERMCreateNotificationActionRequest'] 
slug: /tools/sdk/go/nerm/models/create-notification-action-request
tags: ['SDK', 'Software Development Kit', 'CreateNotificationActionRequest', 'NERMCreateNotificationActionRequest']
---

# CreateNotificationActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowAction** | Pointer to [**NotificationAction**](notification-action) |  | [optional] 

## Methods

### NewCreateNotificationActionRequest

`func NewCreateNotificationActionRequest() *CreateNotificationActionRequest`

NewCreateNotificationActionRequest instantiates a new CreateNotificationActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationActionRequestWithDefaults

`func NewCreateNotificationActionRequestWithDefaults() *CreateNotificationActionRequest`

NewCreateNotificationActionRequestWithDefaults instantiates a new CreateNotificationActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowAction

`func (o *CreateNotificationActionRequest) GetWorkflowAction() NotificationAction`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *CreateNotificationActionRequest) GetWorkflowActionOk() (*NotificationAction, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *CreateNotificationActionRequest) SetWorkflowAction(v NotificationAction)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *CreateNotificationActionRequest) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.


