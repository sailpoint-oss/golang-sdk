---
id: nerm-invitation-action-configuration-attributes
title: InvitationActionConfigurationAttributes
pagination_label: InvitationActionConfigurationAttributes
sidebar_label: InvitationActionConfigurationAttributes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'InvitationActionConfigurationAttributes', 'NERMInvitationActionConfigurationAttributes'] 
slug: /tools/sdk/go/nerm/models/invitation-action-configuration-attributes
tags: ['SDK', 'Software Development Kit', 'InvitationActionConfigurationAttributes', 'NERMInvitationActionConfigurationAttributes']
---

# InvitationActionConfigurationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | the id. | [optional] 
**WaitForCompletion** | Pointer to **bool** | If the invitation action should pause the parent workflow to wait for completion. | [optional] 
**ReturnProfile** | Pointer to **bool** | If the invitation action should return a profile. | [optional] 
**PortalId** | Pointer to **string** | the id of the portal. | [optional] 
**RegistrationWorkflowId** | Pointer to **string** | the id of the registration workflow. | [optional] 
**EmailAttributeId** | Pointer to **string** | the id of the email attribute. | [optional] 
**ValidateCompletedRegistration** | Pointer to **bool** | If the action should validate against completed registrations by email address. | [optional] 
**ValidateOpenRegistration** | Pointer to **bool** | If the action should validate against open registrations by email address. | [optional] 

## Methods

### NewInvitationActionConfigurationAttributes

`func NewInvitationActionConfigurationAttributes() *InvitationActionConfigurationAttributes`

NewInvitationActionConfigurationAttributes instantiates a new InvitationActionConfigurationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationActionConfigurationAttributesWithDefaults

`func NewInvitationActionConfigurationAttributesWithDefaults() *InvitationActionConfigurationAttributes`

NewInvitationActionConfigurationAttributesWithDefaults instantiates a new InvitationActionConfigurationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvitationActionConfigurationAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitationActionConfigurationAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitationActionConfigurationAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvitationActionConfigurationAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWaitForCompletion

`func (o *InvitationActionConfigurationAttributes) GetWaitForCompletion() bool`

GetWaitForCompletion returns the WaitForCompletion field if non-nil, zero value otherwise.

### GetWaitForCompletionOk

`func (o *InvitationActionConfigurationAttributes) GetWaitForCompletionOk() (*bool, bool)`

GetWaitForCompletionOk returns a tuple with the WaitForCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCompletion

`func (o *InvitationActionConfigurationAttributes) SetWaitForCompletion(v bool)`

SetWaitForCompletion sets WaitForCompletion field to given value.

### HasWaitForCompletion

`func (o *InvitationActionConfigurationAttributes) HasWaitForCompletion() bool`

HasWaitForCompletion returns a boolean if a field has been set.

### GetReturnProfile

`func (o *InvitationActionConfigurationAttributes) GetReturnProfile() bool`

GetReturnProfile returns the ReturnProfile field if non-nil, zero value otherwise.

### GetReturnProfileOk

`func (o *InvitationActionConfigurationAttributes) GetReturnProfileOk() (*bool, bool)`

GetReturnProfileOk returns a tuple with the ReturnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnProfile

`func (o *InvitationActionConfigurationAttributes) SetReturnProfile(v bool)`

SetReturnProfile sets ReturnProfile field to given value.

### HasReturnProfile

`func (o *InvitationActionConfigurationAttributes) HasReturnProfile() bool`

HasReturnProfile returns a boolean if a field has been set.

### GetPortalId

`func (o *InvitationActionConfigurationAttributes) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *InvitationActionConfigurationAttributes) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *InvitationActionConfigurationAttributes) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *InvitationActionConfigurationAttributes) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetRegistrationWorkflowId

`func (o *InvitationActionConfigurationAttributes) GetRegistrationWorkflowId() string`

GetRegistrationWorkflowId returns the RegistrationWorkflowId field if non-nil, zero value otherwise.

### GetRegistrationWorkflowIdOk

`func (o *InvitationActionConfigurationAttributes) GetRegistrationWorkflowIdOk() (*string, bool)`

GetRegistrationWorkflowIdOk returns a tuple with the RegistrationWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationWorkflowId

`func (o *InvitationActionConfigurationAttributes) SetRegistrationWorkflowId(v string)`

SetRegistrationWorkflowId sets RegistrationWorkflowId field to given value.

### HasRegistrationWorkflowId

`func (o *InvitationActionConfigurationAttributes) HasRegistrationWorkflowId() bool`

HasRegistrationWorkflowId returns a boolean if a field has been set.

### GetEmailAttributeId

`func (o *InvitationActionConfigurationAttributes) GetEmailAttributeId() string`

GetEmailAttributeId returns the EmailAttributeId field if non-nil, zero value otherwise.

### GetEmailAttributeIdOk

`func (o *InvitationActionConfigurationAttributes) GetEmailAttributeIdOk() (*string, bool)`

GetEmailAttributeIdOk returns a tuple with the EmailAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttributeId

`func (o *InvitationActionConfigurationAttributes) SetEmailAttributeId(v string)`

SetEmailAttributeId sets EmailAttributeId field to given value.

### HasEmailAttributeId

`func (o *InvitationActionConfigurationAttributes) HasEmailAttributeId() bool`

HasEmailAttributeId returns a boolean if a field has been set.

### GetValidateCompletedRegistration

`func (o *InvitationActionConfigurationAttributes) GetValidateCompletedRegistration() bool`

GetValidateCompletedRegistration returns the ValidateCompletedRegistration field if non-nil, zero value otherwise.

### GetValidateCompletedRegistrationOk

`func (o *InvitationActionConfigurationAttributes) GetValidateCompletedRegistrationOk() (*bool, bool)`

GetValidateCompletedRegistrationOk returns a tuple with the ValidateCompletedRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCompletedRegistration

`func (o *InvitationActionConfigurationAttributes) SetValidateCompletedRegistration(v bool)`

SetValidateCompletedRegistration sets ValidateCompletedRegistration field to given value.

### HasValidateCompletedRegistration

`func (o *InvitationActionConfigurationAttributes) HasValidateCompletedRegistration() bool`

HasValidateCompletedRegistration returns a boolean if a field has been set.

### GetValidateOpenRegistration

`func (o *InvitationActionConfigurationAttributes) GetValidateOpenRegistration() bool`

GetValidateOpenRegistration returns the ValidateOpenRegistration field if non-nil, zero value otherwise.

### GetValidateOpenRegistrationOk

`func (o *InvitationActionConfigurationAttributes) GetValidateOpenRegistrationOk() (*bool, bool)`

GetValidateOpenRegistrationOk returns a tuple with the ValidateOpenRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOpenRegistration

`func (o *InvitationActionConfigurationAttributes) SetValidateOpenRegistration(v bool)`

SetValidateOpenRegistration sets ValidateOpenRegistration field to given value.

### HasValidateOpenRegistration

`func (o *InvitationActionConfigurationAttributes) HasValidateOpenRegistration() bool`

HasValidateOpenRegistration returns a boolean if a field has been set.


