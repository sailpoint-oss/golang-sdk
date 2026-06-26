---
id: nerm-profile-type-roles
title: ProfileTypeRoles
pagination_label: ProfileTypeRoles
sidebar_label: ProfileTypeRoles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileTypeRoles', 'NERMProfileTypeRoles'] 
slug: /tools/sdk/go/nerm/models/profile-type-roles
tags: ['SDK', 'Software Development Kit', 'ProfileTypeRoles', 'NERMProfileTypeRoles']
---

# ProfileTypeRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileTypeId** | Pointer to **string** | The id of the profile type | [optional] 
**RoleId** | Pointer to **string** | The id of the role | [optional] 
**Id** | Pointer to **string** | The id of the profile type role | [optional] 

## Methods

### NewProfileTypeRoles

`func NewProfileTypeRoles() *ProfileTypeRoles`

NewProfileTypeRoles instantiates a new ProfileTypeRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileTypeRolesWithDefaults

`func NewProfileTypeRolesWithDefaults() *ProfileTypeRoles`

NewProfileTypeRolesWithDefaults instantiates a new ProfileTypeRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileTypeId

`func (o *ProfileTypeRoles) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *ProfileTypeRoles) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *ProfileTypeRoles) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *ProfileTypeRoles) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.

### GetRoleId

`func (o *ProfileTypeRoles) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ProfileTypeRoles) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ProfileTypeRoles) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ProfileTypeRoles) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetId

`func (o *ProfileTypeRoles) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileTypeRoles) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileTypeRoles) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileTypeRoles) HasId() bool`

HasId returns a boolean if a field has been set.


