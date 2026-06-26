---
id: nermv2025-delegations-post-request
title: DelegationsPostRequest
pagination_label: DelegationsPostRequest
sidebar_label: DelegationsPostRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DelegationsPostRequest', 'NERMV2025DelegationsPostRequest'] 
slug: /tools/sdk/go/nermv2025/models/delegations-post-request
tags: ['SDK', 'Software Development Kit', 'DelegationsPostRequest', 'NERMV2025DelegationsPostRequest']
---

# DelegationsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegation** | Pointer to [**Delegation2**](delegation2) |  | [optional] 

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

`func (o *DelegationsPostRequest) GetDelegation() Delegation2`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *DelegationsPostRequest) GetDelegationOk() (*Delegation2, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *DelegationsPostRequest) SetDelegation(v Delegation2)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *DelegationsPostRequest) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.


