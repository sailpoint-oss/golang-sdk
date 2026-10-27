---
id: v1-bulk-classify-machine-accounts-request
title: BulkClassifyMachineAccountsRequest
pagination_label: BulkClassifyMachineAccountsRequest
sidebar_label: BulkClassifyMachineAccountsRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'BulkClassifyMachineAccountsRequest', 'V1BulkClassifyMachineAccountsRequest'] 
slug: /tools/sdk/go/machineaccountclassify/models/bulk-classify-machine-accounts-request
tags: ['SDK', 'Software Development Kit', 'BulkClassifyMachineAccountsRequest', 'V1BulkClassifyMachineAccountsRequest']
---

# BulkClassifyMachineAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | IDs of the accounts to classify. Provide between 1 and 250 account IDs. | 

## Methods

### NewBulkClassifyMachineAccountsRequest

`func NewBulkClassifyMachineAccountsRequest(ids []string, ) *BulkClassifyMachineAccountsRequest`

NewBulkClassifyMachineAccountsRequest instantiates a new BulkClassifyMachineAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkClassifyMachineAccountsRequestWithDefaults

`func NewBulkClassifyMachineAccountsRequestWithDefaults() *BulkClassifyMachineAccountsRequest`

NewBulkClassifyMachineAccountsRequestWithDefaults instantiates a new BulkClassifyMachineAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BulkClassifyMachineAccountsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkClassifyMachineAccountsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkClassifyMachineAccountsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.



