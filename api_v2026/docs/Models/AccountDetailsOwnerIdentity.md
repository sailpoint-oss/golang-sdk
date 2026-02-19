---
id: v2026-account-details-owner-identity
title: AccountDetailsOwnerIdentity
pagination_label: AccountDetailsOwnerIdentity
sidebar_label: AccountDetailsOwnerIdentity
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AccountDetailsOwnerIdentity', 'V2026AccountDetailsOwnerIdentity'] 
slug: /tools/sdk/go/v2026/models/account-details-owner-identity
tags: ['SDK', 'Software Development Kit', 'AccountDetailsOwnerIdentity', 'V2026AccountDetailsOwnerIdentity']
---

# AccountDetailsOwnerIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DtoType**](dto-type) |  | [optional] 
**Id** | Pointer to **string** | ID of the object to which this reference applies | [optional] 
**Name** | Pointer to **string** | Human-readable display name of the object to which this reference applies | [optional] 

## Methods

### NewAccountDetailsOwnerIdentity

`func NewAccountDetailsOwnerIdentity() *AccountDetailsOwnerIdentity`

NewAccountDetailsOwnerIdentity instantiates a new AccountDetailsOwnerIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsOwnerIdentityWithDefaults

`func NewAccountDetailsOwnerIdentityWithDefaults() *AccountDetailsOwnerIdentity`

NewAccountDetailsOwnerIdentityWithDefaults instantiates a new AccountDetailsOwnerIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountDetailsOwnerIdentity) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountDetailsOwnerIdentity) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountDetailsOwnerIdentity) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *AccountDetailsOwnerIdentity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AccountDetailsOwnerIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountDetailsOwnerIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountDetailsOwnerIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountDetailsOwnerIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountDetailsOwnerIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountDetailsOwnerIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountDetailsOwnerIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountDetailsOwnerIdentity) HasName() bool`

HasName returns a boolean if a field has been set.


