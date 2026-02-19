---
id: v2026-machine-account-sub-type-config-dto-machine-account-delete
title: MachineAccountSubTypeConfigDtoMachineAccountDelete
pagination_label: MachineAccountSubTypeConfigDtoMachineAccountDelete
sidebar_label: MachineAccountSubTypeConfigDtoMachineAccountDelete
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccountSubTypeConfigDtoMachineAccountDelete', 'V2026MachineAccountSubTypeConfigDtoMachineAccountDelete'] 
slug: /tools/sdk/go/v2026/models/machine-account-sub-type-config-dto-machine-account-delete
tags: ['SDK', 'Software Development Kit', 'MachineAccountSubTypeConfigDtoMachineAccountDelete', 'V2026MachineAccountSubTypeConfigDtoMachineAccountDelete']
---

# MachineAccountSubTypeConfigDtoMachineAccountDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalRequired** | Pointer to **bool** | Indicates whether approval is required for an account deletion request. | [optional] [default to false]
**ApprovalConfig** | Pointer to [**ApprovalConfig**](approval-config) |  | [optional] 

## Methods

### NewMachineAccountSubTypeConfigDtoMachineAccountDelete

`func NewMachineAccountSubTypeConfigDtoMachineAccountDelete() *MachineAccountSubTypeConfigDtoMachineAccountDelete`

NewMachineAccountSubTypeConfigDtoMachineAccountDelete instantiates a new MachineAccountSubTypeConfigDtoMachineAccountDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountSubTypeConfigDtoMachineAccountDeleteWithDefaults

`func NewMachineAccountSubTypeConfigDtoMachineAccountDeleteWithDefaults() *MachineAccountSubTypeConfigDtoMachineAccountDelete`

NewMachineAccountSubTypeConfigDtoMachineAccountDeleteWithDefaults instantiates a new MachineAccountSubTypeConfigDtoMachineAccountDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalRequired

`func (o *MachineAccountSubTypeConfigDtoMachineAccountDelete) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *MachineAccountSubTypeConfigDtoMachineAccountDelete) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *MachineAccountSubTypeConfigDtoMachineAccountDelete) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.

### HasApprovalRequired

`func (o *MachineAccountSubTypeConfigDtoMachineAccountDelete) HasApprovalRequired() bool`

HasApprovalRequired returns a boolean if a field has been set.

### GetApprovalConfig

`func (o *MachineAccountSubTypeConfigDtoMachineAccountDelete) GetApprovalConfig() ApprovalConfig`

GetApprovalConfig returns the ApprovalConfig field if non-nil, zero value otherwise.

### GetApprovalConfigOk

`func (o *MachineAccountSubTypeConfigDtoMachineAccountDelete) GetApprovalConfigOk() (*ApprovalConfig, bool)`

GetApprovalConfigOk returns a tuple with the ApprovalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfig

`func (o *MachineAccountSubTypeConfigDtoMachineAccountDelete) SetApprovalConfig(v ApprovalConfig)`

SetApprovalConfig sets ApprovalConfig field to given value.

### HasApprovalConfig

`func (o *MachineAccountSubTypeConfigDtoMachineAccountDelete) HasApprovalConfig() bool`

HasApprovalConfig returns a boolean if a field has been set.


