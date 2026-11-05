---
id: v1-accessprofilev2-privilege-level
title: Accessprofilev2PrivilegeLevel
pagination_label: Accessprofilev2PrivilegeLevel
sidebar_label: Accessprofilev2PrivilegeLevel
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Accessprofilev2PrivilegeLevel', 'V1Accessprofilev2PrivilegeLevel'] 
slug: /tools/sdk/go/accessprofiles/models/accessprofilev2-privilege-level
tags: ['SDK', 'Software Development Kit', 'Accessprofilev2PrivilegeLevel', 'V1Accessprofilev2PrivilegeLevel']
---

# Accessprofilev2PrivilegeLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effective** | Pointer to **string** | Effective privilege level of the access profile, derived from the highest privilege level of the entitlements it contains. `UNDETERMINED` indicates the level could not be calculated. | [optional] 

## Methods

### NewAccessprofilev2PrivilegeLevel

`func NewAccessprofilev2PrivilegeLevel() *Accessprofilev2PrivilegeLevel`

NewAccessprofilev2PrivilegeLevel instantiates a new Accessprofilev2PrivilegeLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessprofilev2PrivilegeLevelWithDefaults

`func NewAccessprofilev2PrivilegeLevelWithDefaults() *Accessprofilev2PrivilegeLevel`

NewAccessprofilev2PrivilegeLevelWithDefaults instantiates a new Accessprofilev2PrivilegeLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffective

`func (o *Accessprofilev2PrivilegeLevel) GetEffective() string`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *Accessprofilev2PrivilegeLevel) GetEffectiveOk() (*string, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *Accessprofilev2PrivilegeLevel) SetEffective(v string)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *Accessprofilev2PrivilegeLevel) HasEffective() bool`

HasEffective returns a boolean if a field has been set.


