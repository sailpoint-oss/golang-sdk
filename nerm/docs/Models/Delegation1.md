---
id: nerm-delegation1
title: Delegation1
pagination_label: Delegation1
sidebar_label: Delegation1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Delegation1', 'NERMDelegation1'] 
slug: /tools/sdk/go/nerm/models/delegation1
tags: ['SDK', 'Software Development Kit', 'Delegation1', 'NERMDelegation1']
---

# Delegation1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatorId** | Pointer to **string** | The id of the delegator | [optional] 
**DelegateId** | Pointer to **string** | The id of the delegate | [optional] 
**Expiration** | Pointer to **SailPointTime** | The expiration date of the delegation | [optional] 

## Methods

### NewDelegation1

`func NewDelegation1() *Delegation1`

NewDelegation1 instantiates a new Delegation1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegation1WithDefaults

`func NewDelegation1WithDefaults() *Delegation1`

NewDelegation1WithDefaults instantiates a new Delegation1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatorId

`func (o *Delegation1) GetDelegatorId() string`

GetDelegatorId returns the DelegatorId field if non-nil, zero value otherwise.

### GetDelegatorIdOk

`func (o *Delegation1) GetDelegatorIdOk() (*string, bool)`

GetDelegatorIdOk returns a tuple with the DelegatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorId

`func (o *Delegation1) SetDelegatorId(v string)`

SetDelegatorId sets DelegatorId field to given value.

### HasDelegatorId

`func (o *Delegation1) HasDelegatorId() bool`

HasDelegatorId returns a boolean if a field has been set.

### GetDelegateId

`func (o *Delegation1) GetDelegateId() string`

GetDelegateId returns the DelegateId field if non-nil, zero value otherwise.

### GetDelegateIdOk

`func (o *Delegation1) GetDelegateIdOk() (*string, bool)`

GetDelegateIdOk returns a tuple with the DelegateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateId

`func (o *Delegation1) SetDelegateId(v string)`

SetDelegateId sets DelegateId field to given value.

### HasDelegateId

`func (o *Delegation1) HasDelegateId() bool`

HasDelegateId returns a boolean if a field has been set.

### GetExpiration

`func (o *Delegation1) GetExpiration() SailPointTime`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Delegation1) GetExpirationOk() (*SailPointTime, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Delegation1) SetExpiration(v SailPointTime)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Delegation1) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


