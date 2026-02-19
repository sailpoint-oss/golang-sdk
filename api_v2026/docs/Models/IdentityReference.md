---
id: v2026-identity-reference
title: IdentityReference
pagination_label: IdentityReference
sidebar_label: IdentityReference
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IdentityReference', 'V2026IdentityReference'] 
slug: /tools/sdk/go/v2026/models/identity-reference
tags: ['SDK', 'Software Development Kit', 'IdentityReference', 'V2026IdentityReference']
---

# IdentityReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of identity | [optional] 
**Name** | Pointer to **string** | Name of Identity | [optional] 
**Email** | Pointer to **NullableString** | mail id of identity | [optional] 
**Status** | Pointer to **NullableString** | status of identity UNREGISTERED/REGISTERED | [optional] 

## Methods

### NewIdentityReference

`func NewIdentityReference() *IdentityReference`

NewIdentityReference instantiates a new IdentityReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityReferenceWithDefaults

`func NewIdentityReferenceWithDefaults() *IdentityReference`

NewIdentityReferenceWithDefaults instantiates a new IdentityReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *IdentityReference) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IdentityReference) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IdentityReference) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IdentityReference) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IdentityReference) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IdentityReference) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetStatus

`func (o *IdentityReference) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityReference) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityReference) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityReference) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *IdentityReference) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *IdentityReference) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

