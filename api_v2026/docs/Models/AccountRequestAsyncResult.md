---
id: v2026-account-request-async-result
title: AccountRequestAsyncResult
pagination_label: AccountRequestAsyncResult
sidebar_label: AccountRequestAsyncResult
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AccountRequestAsyncResult', 'V2026AccountRequestAsyncResult'] 
slug: /tools/sdk/go/v2026/models/account-request-async-result
tags: ['SDK', 'Software Development Kit', 'AccountRequestAsyncResult', 'V2026AccountRequestAsyncResult']
---

# AccountRequestAsyncResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRequestId** | **string** | Id of the account request | 

## Methods

### NewAccountRequestAsyncResult

`func NewAccountRequestAsyncResult(accountRequestId string, ) *AccountRequestAsyncResult`

NewAccountRequestAsyncResult instantiates a new AccountRequestAsyncResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRequestAsyncResultWithDefaults

`func NewAccountRequestAsyncResultWithDefaults() *AccountRequestAsyncResult`

NewAccountRequestAsyncResultWithDefaults instantiates a new AccountRequestAsyncResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRequestId

`func (o *AccountRequestAsyncResult) GetAccountRequestId() string`

GetAccountRequestId returns the AccountRequestId field if non-nil, zero value otherwise.

### GetAccountRequestIdOk

`func (o *AccountRequestAsyncResult) GetAccountRequestIdOk() (*string, bool)`

GetAccountRequestIdOk returns a tuple with the AccountRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRequestId

`func (o *AccountRequestAsyncResult) SetAccountRequestId(v string)`

SetAccountRequestId sets AccountRequestId field to given value.



