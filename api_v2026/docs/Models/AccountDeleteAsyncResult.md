---
id: v2026-account-delete-async-result
title: AccountDeleteAsyncResult
pagination_label: AccountDeleteAsyncResult
sidebar_label: AccountDeleteAsyncResult
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AccountDeleteAsyncResult', 'V2026AccountDeleteAsyncResult'] 
slug: /tools/sdk/go/v2026/models/account-delete-async-result
tags: ['SDK', 'Software Development Kit', 'AccountDeleteAsyncResult', 'V2026AccountDeleteAsyncResult']
---

# AccountDeleteAsyncResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRequestId** | **string** | Id of the deletion request | 

## Methods

### NewAccountDeleteAsyncResult

`func NewAccountDeleteAsyncResult(accountRequestId string, ) *AccountDeleteAsyncResult`

NewAccountDeleteAsyncResult instantiates a new AccountDeleteAsyncResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDeleteAsyncResultWithDefaults

`func NewAccountDeleteAsyncResultWithDefaults() *AccountDeleteAsyncResult`

NewAccountDeleteAsyncResultWithDefaults instantiates a new AccountDeleteAsyncResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRequestId

`func (o *AccountDeleteAsyncResult) GetAccountRequestId() string`

GetAccountRequestId returns the AccountRequestId field if non-nil, zero value otherwise.

### GetAccountRequestIdOk

`func (o *AccountDeleteAsyncResult) GetAccountRequestIdOk() (*string, bool)`

GetAccountRequestIdOk returns a tuple with the AccountRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRequestId

`func (o *AccountDeleteAsyncResult) SetAccountRequestId(v string)`

SetAccountRequestId sets AccountRequestId field to given value.



