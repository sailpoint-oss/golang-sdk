---
id: v1-machine-account-request-async-result
title: MachineAccountRequestAsyncResult
pagination_label: MachineAccountRequestAsyncResult
sidebar_label: MachineAccountRequestAsyncResult
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccountRequestAsyncResult', 'V1MachineAccountRequestAsyncResult'] 
slug: /tools/sdk/go/machineaccounts/models/machine-account-request-async-result
tags: ['SDK', 'Software Development Kit', 'MachineAccountRequestAsyncResult', 'V1MachineAccountRequestAsyncResult']
---

# MachineAccountRequestAsyncResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRequestId** | **string** | ID of the task that tracks the account request. | 

## Methods

### NewMachineAccountRequestAsyncResult

`func NewMachineAccountRequestAsyncResult(accountRequestId string, ) *MachineAccountRequestAsyncResult`

NewMachineAccountRequestAsyncResult instantiates a new MachineAccountRequestAsyncResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountRequestAsyncResultWithDefaults

`func NewMachineAccountRequestAsyncResultWithDefaults() *MachineAccountRequestAsyncResult`

NewMachineAccountRequestAsyncResultWithDefaults instantiates a new MachineAccountRequestAsyncResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRequestId

`func (o *MachineAccountRequestAsyncResult) GetAccountRequestId() string`

GetAccountRequestId returns the AccountRequestId field if non-nil, zero value otherwise.

### GetAccountRequestIdOk

`func (o *MachineAccountRequestAsyncResult) GetAccountRequestIdOk() (*string, bool)`

GetAccountRequestIdOk returns a tuple with the AccountRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRequestId

`func (o *MachineAccountRequestAsyncResult) SetAccountRequestId(v string)`

SetAccountRequestId sets AccountRequestId field to given value.



