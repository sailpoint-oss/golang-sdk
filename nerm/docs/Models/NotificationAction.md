---
id: nerm-notification-action
title: NotificationAction
pagination_label: NotificationAction
sidebar_label: NotificationAction
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'NotificationAction', 'NERMNotificationAction'] 
slug: /tools/sdk/go/nerm/models/notification-action
tags: ['SDK', 'Software Development Kit', 'NotificationAction', 'NERMNotificationAction']
---

# NotificationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The workflow the workflow action belongs to. | 
**Description** | **string** | The description of the workflow action. | 
**EmailAttributeId** | Pointer to **string** | The attribute storing the email address for the workflow action. | [optional] 
**EmailAddresses** | Pointer to **[]string** | The email addresses for the workflow action. | [optional] 
**Archived** | Pointer to **bool** | If the workflow action is archived or not. | [optional] [default to false]
**WorkflowActionEmailAttributes** | Pointer to [**NotificationActionWorkflowActionEmailAttributes**](notification-action-workflow-action-email-attributes) |  | [optional] 

## Methods

### NewNotificationAction

`func NewNotificationAction(workflowId string, description string, ) *NotificationAction`

NewNotificationAction instantiates a new NotificationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationActionWithDefaults

`func NewNotificationActionWithDefaults() *NotificationAction`

NewNotificationActionWithDefaults instantiates a new NotificationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *NotificationAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *NotificationAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *NotificationAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetDescription

`func (o *NotificationAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationAction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmailAttributeId

`func (o *NotificationAction) GetEmailAttributeId() string`

GetEmailAttributeId returns the EmailAttributeId field if non-nil, zero value otherwise.

### GetEmailAttributeIdOk

`func (o *NotificationAction) GetEmailAttributeIdOk() (*string, bool)`

GetEmailAttributeIdOk returns a tuple with the EmailAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttributeId

`func (o *NotificationAction) SetEmailAttributeId(v string)`

SetEmailAttributeId sets EmailAttributeId field to given value.

### HasEmailAttributeId

`func (o *NotificationAction) HasEmailAttributeId() bool`

HasEmailAttributeId returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *NotificationAction) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *NotificationAction) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *NotificationAction) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *NotificationAction) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetArchived

`func (o *NotificationAction) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *NotificationAction) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *NotificationAction) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *NotificationAction) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetWorkflowActionEmailAttributes

`func (o *NotificationAction) GetWorkflowActionEmailAttributes() NotificationActionWorkflowActionEmailAttributes`

GetWorkflowActionEmailAttributes returns the WorkflowActionEmailAttributes field if non-nil, zero value otherwise.

### GetWorkflowActionEmailAttributesOk

`func (o *NotificationAction) GetWorkflowActionEmailAttributesOk() (*NotificationActionWorkflowActionEmailAttributes, bool)`

GetWorkflowActionEmailAttributesOk returns a tuple with the WorkflowActionEmailAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionEmailAttributes

`func (o *NotificationAction) SetWorkflowActionEmailAttributes(v NotificationActionWorkflowActionEmailAttributes)`

SetWorkflowActionEmailAttributes sets WorkflowActionEmailAttributes field to given value.

### HasWorkflowActionEmailAttributes

`func (o *NotificationAction) HasWorkflowActionEmailAttributes() bool`

HasWorkflowActionEmailAttributes returns a boolean if a field has been set.


