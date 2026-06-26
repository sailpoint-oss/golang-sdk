---
id: nerm-delegations-post-request
title: DelegationsPostRequest
pagination_label: DelegationsPostRequest
sidebar_label: DelegationsPostRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DelegationsPostRequest', 'NERMDelegationsPostRequest'] 
slug: /tools/sdk/go/nerm/models/delegations-post-request
tags: ['SDK', 'Software Development Kit', 'DelegationsPostRequest', 'NERMDelegationsPostRequest']
---

# DelegationsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegation** | Pointer to [**Delegation1**](delegation1) |  | [optional] 

## Methods

### NewDelegationsPostRequest

`func NewDelegationsPostRequest() *DelegationsPostRequest`

NewDelegationsPostRequest instantiates a new DelegationsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationsPostRequestWithDefaults

`func NewDelegationsPostRequestWithDefaults() *DelegationsPostRequest`

NewDelegationsPostRequestWithDefaults instantiates a new DelegationsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegation

`func (o *DelegationsPostRequest) GetDelegation() Delegation1`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *DelegationsPostRequest) GetDelegationOk() (*Delegation1, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *DelegationsPostRequest) SetDelegation(v Delegation1)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *DelegationsPostRequest) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.


