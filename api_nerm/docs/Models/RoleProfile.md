---
id: nerm-role-profile
title: RoleProfile
pagination_label: RoleProfile
sidebar_label: RoleProfile
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RoleProfile', 'NERMRoleProfile'] 
slug: /tools/sdk/go/nerm/models/role-profile
tags: ['SDK', 'Software Development Kit', 'RoleProfile', 'NERMRoleProfile']
---

# RoleProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Uid** | Pointer to **string** |  | [optional] [readonly] 
**RoleId** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 

## Methods

### NewRoleProfile

`func NewRoleProfile() *RoleProfile`

NewRoleProfile instantiates a new RoleProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleProfileWithDefaults

`func NewRoleProfileWithDefaults() *RoleProfile`

NewRoleProfileWithDefaults instantiates a new RoleProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *RoleProfile) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RoleProfile) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RoleProfile) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *RoleProfile) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetRoleId

`func (o *RoleProfile) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RoleProfile) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RoleProfile) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *RoleProfile) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetProfileId

`func (o *RoleProfile) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *RoleProfile) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *RoleProfile) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *RoleProfile) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


