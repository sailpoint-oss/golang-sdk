---
id: v2026-account-action-request-dto-correlated-identity
title: AccountActionRequestDtoCorrelatedIdentity
pagination_label: AccountActionRequestDtoCorrelatedIdentity
sidebar_label: AccountActionRequestDtoCorrelatedIdentity
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AccountActionRequestDtoCorrelatedIdentity', 'V2026AccountActionRequestDtoCorrelatedIdentity'] 
slug: /tools/sdk/go/v2026/models/account-action-request-dto-correlated-identity
tags: ['SDK', 'Software Development Kit', 'AccountActionRequestDtoCorrelatedIdentity', 'V2026AccountActionRequestDtoCorrelatedIdentity']
---

# AccountActionRequestDtoCorrelatedIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of identity | [optional] 
**Name** | Pointer to **string** | Name of Identity | [optional] 
**Email** | Pointer to **NullableString** | mail id of identity | [optional] 
**Status** | Pointer to **NullableString** | status of identity UNREGISTERED/REGISTERED | [optional] 

## Methods

### NewAccountActionRequestDtoCorrelatedIdentity

`func NewAccountActionRequestDtoCorrelatedIdentity() *AccountActionRequestDtoCorrelatedIdentity`

NewAccountActionRequestDtoCorrelatedIdentity instantiates a new AccountActionRequestDtoCorrelatedIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountActionRequestDtoCorrelatedIdentityWithDefaults

`func NewAccountActionRequestDtoCorrelatedIdentityWithDefaults() *AccountActionRequestDtoCorrelatedIdentity`

NewAccountActionRequestDtoCorrelatedIdentityWithDefaults instantiates a new AccountActionRequestDtoCorrelatedIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountActionRequestDtoCorrelatedIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountActionRequestDtoCorrelatedIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountActionRequestDtoCorrelatedIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountActionRequestDtoCorrelatedIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountActionRequestDtoCorrelatedIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountActionRequestDtoCorrelatedIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountActionRequestDtoCorrelatedIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountActionRequestDtoCorrelatedIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *AccountActionRequestDtoCorrelatedIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountActionRequestDtoCorrelatedIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountActionRequestDtoCorrelatedIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountActionRequestDtoCorrelatedIdentity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AccountActionRequestDtoCorrelatedIdentity) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AccountActionRequestDtoCorrelatedIdentity) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetStatus

`func (o *AccountActionRequestDtoCorrelatedIdentity) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountActionRequestDtoCorrelatedIdentity) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountActionRequestDtoCorrelatedIdentity) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountActionRequestDtoCorrelatedIdentity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AccountActionRequestDtoCorrelatedIdentity) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AccountActionRequestDtoCorrelatedIdentity) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

