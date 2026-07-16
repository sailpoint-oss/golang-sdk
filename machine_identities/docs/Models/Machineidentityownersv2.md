---
id: v1-machineidentityownersv2
title: Machineidentityownersv2
pagination_label: Machineidentityownersv2
sidebar_label: Machineidentityownersv2
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Machineidentityownersv2', 'V1Machineidentityownersv2'] 
slug: /tools/sdk/go/machineidentities/models/machineidentityownersv2
tags: ['SDK', 'Software Development Kit', 'Machineidentityownersv2', 'V1Machineidentityownersv2']
---

# Machineidentityownersv2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to [**Machineidentityownersv2Primary**](machineidentityownersv2-primary) |  | [optional] 
**Secondary** | Pointer to [**[]Basereferencedto**](basereferencedto) | Additional owners. Entries are either up to ten human (IDENTITY) references or exactly one GOVERNANCE_GROUP reference - not both. Governance-group owners appear here with type GOVERNANCE_GROUP. | [optional] 

## Methods

### NewMachineidentityownersv2

`func NewMachineidentityownersv2() *Machineidentityownersv2`

NewMachineidentityownersv2 instantiates a new Machineidentityownersv2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineidentityownersv2WithDefaults

`func NewMachineidentityownersv2WithDefaults() *Machineidentityownersv2`

NewMachineidentityownersv2WithDefaults instantiates a new Machineidentityownersv2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *Machineidentityownersv2) GetPrimary() Machineidentityownersv2Primary`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Machineidentityownersv2) GetPrimaryOk() (*Machineidentityownersv2Primary, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Machineidentityownersv2) SetPrimary(v Machineidentityownersv2Primary)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *Machineidentityownersv2) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSecondary

`func (o *Machineidentityownersv2) GetSecondary() []Basereferencedto`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *Machineidentityownersv2) GetSecondaryOk() (*[]Basereferencedto, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *Machineidentityownersv2) SetSecondary(v []Basereferencedto)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *Machineidentityownersv2) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


