---
id: v2026-machine-account-sub-type-config-dto
title: MachineAccountSubTypeConfigDto
pagination_label: MachineAccountSubTypeConfigDto
sidebar_label: MachineAccountSubTypeConfigDto
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccountSubTypeConfigDto', 'V2026MachineAccountSubTypeConfigDto'] 
slug: /tools/sdk/go/v2026/models/machine-account-sub-type-config-dto
tags: ['SDK', 'Software Development Kit', 'MachineAccountSubTypeConfigDto', 'V2026MachineAccountSubTypeConfigDto']
---

# MachineAccountSubTypeConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubtypeId** | Pointer to **string** | Unique identifier representing the specific subtype of the machine account, used to distinguish between different machine account categories. | [optional] 
**MachineAccountCreate** | Pointer to [**MachineAccountSubTypeConfigDtoMachineAccountCreate**](machine-account-sub-type-config-dto-machine-account-create) |  | [optional] 
**MachineAccountDelete** | Pointer to [**MachineAccountSubTypeConfigDtoMachineAccountDelete**](machine-account-sub-type-config-dto-machine-account-delete) |  | [optional] 

## Methods

### NewMachineAccountSubTypeConfigDto

`func NewMachineAccountSubTypeConfigDto() *MachineAccountSubTypeConfigDto`

NewMachineAccountSubTypeConfigDto instantiates a new MachineAccountSubTypeConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountSubTypeConfigDtoWithDefaults

`func NewMachineAccountSubTypeConfigDtoWithDefaults() *MachineAccountSubTypeConfigDto`

NewMachineAccountSubTypeConfigDtoWithDefaults instantiates a new MachineAccountSubTypeConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtypeId

`func (o *MachineAccountSubTypeConfigDto) GetSubtypeId() string`

GetSubtypeId returns the SubtypeId field if non-nil, zero value otherwise.

### GetSubtypeIdOk

`func (o *MachineAccountSubTypeConfigDto) GetSubtypeIdOk() (*string, bool)`

GetSubtypeIdOk returns a tuple with the SubtypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtypeId

`func (o *MachineAccountSubTypeConfigDto) SetSubtypeId(v string)`

SetSubtypeId sets SubtypeId field to given value.

### HasSubtypeId

`func (o *MachineAccountSubTypeConfigDto) HasSubtypeId() bool`

HasSubtypeId returns a boolean if a field has been set.

### GetMachineAccountCreate

`func (o *MachineAccountSubTypeConfigDto) GetMachineAccountCreate() MachineAccountSubTypeConfigDtoMachineAccountCreate`

GetMachineAccountCreate returns the MachineAccountCreate field if non-nil, zero value otherwise.

### GetMachineAccountCreateOk

`func (o *MachineAccountSubTypeConfigDto) GetMachineAccountCreateOk() (*MachineAccountSubTypeConfigDtoMachineAccountCreate, bool)`

GetMachineAccountCreateOk returns a tuple with the MachineAccountCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountCreate

`func (o *MachineAccountSubTypeConfigDto) SetMachineAccountCreate(v MachineAccountSubTypeConfigDtoMachineAccountCreate)`

SetMachineAccountCreate sets MachineAccountCreate field to given value.

### HasMachineAccountCreate

`func (o *MachineAccountSubTypeConfigDto) HasMachineAccountCreate() bool`

HasMachineAccountCreate returns a boolean if a field has been set.

### GetMachineAccountDelete

`func (o *MachineAccountSubTypeConfigDto) GetMachineAccountDelete() MachineAccountSubTypeConfigDtoMachineAccountDelete`

GetMachineAccountDelete returns the MachineAccountDelete field if non-nil, zero value otherwise.

### GetMachineAccountDeleteOk

`func (o *MachineAccountSubTypeConfigDto) GetMachineAccountDeleteOk() (*MachineAccountSubTypeConfigDtoMachineAccountDelete, bool)`

GetMachineAccountDeleteOk returns a tuple with the MachineAccountDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAccountDelete

`func (o *MachineAccountSubTypeConfigDto) SetMachineAccountDelete(v MachineAccountSubTypeConfigDtoMachineAccountDelete)`

SetMachineAccountDelete sets MachineAccountDelete field to given value.

### HasMachineAccountDelete

`func (o *MachineAccountSubTypeConfigDto) HasMachineAccountDelete() bool`

HasMachineAccountDelete returns a boolean if a field has been set.


