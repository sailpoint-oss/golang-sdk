---
id: v2026-machine-account-sub-type-config-dto-machine-account-create
title: MachineAccountSubTypeConfigDtoMachineAccountCreate
pagination_label: MachineAccountSubTypeConfigDtoMachineAccountCreate
sidebar_label: MachineAccountSubTypeConfigDtoMachineAccountCreate
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccountSubTypeConfigDtoMachineAccountCreate', 'V2026MachineAccountSubTypeConfigDtoMachineAccountCreate'] 
slug: /tools/sdk/go/v2026/models/machine-account-sub-type-config-dto-machine-account-create
tags: ['SDK', 'Software Development Kit', 'MachineAccountSubTypeConfigDtoMachineAccountCreate', 'V2026MachineAccountSubTypeConfigDtoMachineAccountCreate']
---

# MachineAccountSubTypeConfigDtoMachineAccountCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCreateEnabled** | Pointer to **bool** | Specifies if the creation of machine accounts is allowed for this subtype. | [optional] [default to false]
**ApprovalRequired** | Pointer to **bool** | Specifies if approval is needed before a machine account can be created for this subtype. | [optional] [default to false]
**FormId** | Pointer to **string** | formId | [optional] 
**EntitlementId** | Pointer to **string** | Configuration details specifying who can approve machine account creation requests and the types of comments allowed during the approval process. | [optional] 
**ApprovalConfig** | Pointer to [**ApprovalConfig**](approval-config) |  | [optional] 

## Methods

### NewMachineAccountSubTypeConfigDtoMachineAccountCreate

`func NewMachineAccountSubTypeConfigDtoMachineAccountCreate() *MachineAccountSubTypeConfigDtoMachineAccountCreate`

NewMachineAccountSubTypeConfigDtoMachineAccountCreate instantiates a new MachineAccountSubTypeConfigDtoMachineAccountCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountSubTypeConfigDtoMachineAccountCreateWithDefaults

`func NewMachineAccountSubTypeConfigDtoMachineAccountCreateWithDefaults() *MachineAccountSubTypeConfigDtoMachineAccountCreate`

NewMachineAccountSubTypeConfigDtoMachineAccountCreateWithDefaults instantiates a new MachineAccountSubTypeConfigDtoMachineAccountCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCreateEnabled

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetAccountCreateEnabled() bool`

GetAccountCreateEnabled returns the AccountCreateEnabled field if non-nil, zero value otherwise.

### GetAccountCreateEnabledOk

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetAccountCreateEnabledOk() (*bool, bool)`

GetAccountCreateEnabledOk returns a tuple with the AccountCreateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCreateEnabled

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) SetAccountCreateEnabled(v bool)`

SetAccountCreateEnabled sets AccountCreateEnabled field to given value.

### HasAccountCreateEnabled

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) HasAccountCreateEnabled() bool`

HasAccountCreateEnabled returns a boolean if a field has been set.

### GetApprovalRequired

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.

### HasApprovalRequired

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) HasApprovalRequired() bool`

HasApprovalRequired returns a boolean if a field has been set.

### GetFormId

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetEntitlementId

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.

### HasEntitlementId

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) HasEntitlementId() bool`

HasEntitlementId returns a boolean if a field has been set.

### GetApprovalConfig

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetApprovalConfig() ApprovalConfig`

GetApprovalConfig returns the ApprovalConfig field if non-nil, zero value otherwise.

### GetApprovalConfigOk

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) GetApprovalConfigOk() (*ApprovalConfig, bool)`

GetApprovalConfigOk returns a tuple with the ApprovalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfig

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) SetApprovalConfig(v ApprovalConfig)`

SetApprovalConfig sets ApprovalConfig field to given value.

### HasApprovalConfig

`func (o *MachineAccountSubTypeConfigDtoMachineAccountCreate) HasApprovalConfig() bool`

HasApprovalConfig returns a boolean if a field has been set.


