---
id: v1-identity-v2-manager
title: IdentityV2Manager
pagination_label: IdentityV2Manager
sidebar_label: IdentityV2Manager
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IdentityV2Manager', 'V1IdentityV2Manager'] 
slug: /tools/sdk/go/identities/models/identity-v2-manager
tags: ['SDK', 'Software Development Kit', 'IdentityV2Manager', 'V1IdentityV2Manager']
---

# IdentityV2Manager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | ID of the identity's manager. | [optional] 
**Name** | Pointer to **NullableString** | Display name of the identity's manager. | [optional] 
**Type** | Pointer to **NullableString** | DTO type of the identity's manager. | [optional] 

## Methods

### NewIdentityV2Manager

`func NewIdentityV2Manager() *IdentityV2Manager`

NewIdentityV2Manager instantiates a new IdentityV2Manager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityV2ManagerWithDefaults

`func NewIdentityV2ManagerWithDefaults() *IdentityV2Manager`

NewIdentityV2ManagerWithDefaults instantiates a new IdentityV2Manager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityV2Manager) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityV2Manager) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityV2Manager) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityV2Manager) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IdentityV2Manager) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IdentityV2Manager) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *IdentityV2Manager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityV2Manager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityV2Manager) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityV2Manager) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IdentityV2Manager) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdentityV2Manager) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *IdentityV2Manager) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityV2Manager) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityV2Manager) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityV2Manager) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *IdentityV2Manager) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *IdentityV2Manager) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

