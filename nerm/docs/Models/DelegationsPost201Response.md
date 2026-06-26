---
id: nerm-delegations-post201-response
title: DelegationsPost201Response
pagination_label: DelegationsPost201Response
sidebar_label: DelegationsPost201Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DelegationsPost201Response', 'NERMDelegationsPost201Response'] 
slug: /tools/sdk/go/nerm/models/delegations-post201-response
tags: ['SDK', 'Software Development Kit', 'DelegationsPost201Response', 'NERMDelegationsPost201Response']
---

# DelegationsPost201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegation** | Pointer to [**Delegation**](delegation) |  | [optional] 

## Methods

### NewDelegationsPost201Response

`func NewDelegationsPost201Response() *DelegationsPost201Response`

NewDelegationsPost201Response instantiates a new DelegationsPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationsPost201ResponseWithDefaults

`func NewDelegationsPost201ResponseWithDefaults() *DelegationsPost201Response`

NewDelegationsPost201ResponseWithDefaults instantiates a new DelegationsPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegation

`func (o *DelegationsPost201Response) GetDelegation() Delegation`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *DelegationsPost201Response) GetDelegationOk() (*Delegation, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *DelegationsPost201Response) SetDelegation(v Delegation)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *DelegationsPost201Response) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.


