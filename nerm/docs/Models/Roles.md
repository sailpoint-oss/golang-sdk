---
id: nerm-roles
title: Roles
pagination_label: Roles
sidebar_label: Roles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Roles', 'NERMRoles'] 
slug: /tools/sdk/go/nerm/models/roles
tags: ['SDK', 'Software Development Kit', 'Roles', 'NERMRoles']
---

# Roles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]Role**](role) |  | [optional] 

## Methods

### NewRoles

`func NewRoles() *Roles`

NewRoles instantiates a new Roles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesWithDefaults

`func NewRolesWithDefaults() *Roles`

NewRolesWithDefaults instantiates a new Roles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *Roles) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Roles) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Roles) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Roles) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


