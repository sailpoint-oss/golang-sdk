---
id: nermv2025-delegation2
title: Delegation2
pagination_label: Delegation2
sidebar_label: Delegation2
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Delegation2', 'NERMV2025Delegation2'] 
slug: /tools/sdk/go/nermv2025/models/delegation2
tags: ['SDK', 'Software Development Kit', 'Delegation2', 'NERMV2025Delegation2']
---

# Delegation2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatorId** | Pointer to **string** | The id of the delegator | [optional] 
**DelegateId** | Pointer to **string** | The id of the delegate | [optional] 
**Expiration** | Pointer to **SailPointTime** | The expiration date of the delegation | [optional] 

## Methods

### NewDelegation2

`func NewDelegation2() *Delegation2`

NewDelegation2 instantiates a new Delegation2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegation2WithDefaults

`func NewDelegation2WithDefaults() *Delegation2`

NewDelegation2WithDefaults instantiates a new Delegation2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatorId

`func (o *Delegation2) GetDelegatorId() string`

GetDelegatorId returns the DelegatorId field if non-nil, zero value otherwise.

### GetDelegatorIdOk

`func (o *Delegation2) GetDelegatorIdOk() (*string, bool)`

GetDelegatorIdOk returns a tuple with the DelegatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorId

`func (o *Delegation2) SetDelegatorId(v string)`

SetDelegatorId sets DelegatorId field to given value.

### HasDelegatorId

`func (o *Delegation2) HasDelegatorId() bool`

HasDelegatorId returns a boolean if a field has been set.

### GetDelegateId

`func (o *Delegation2) GetDelegateId() string`

GetDelegateId returns the DelegateId field if non-nil, zero value otherwise.

### GetDelegateIdOk

`func (o *Delegation2) GetDelegateIdOk() (*string, bool)`

GetDelegateIdOk returns a tuple with the DelegateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateId

`func (o *Delegation2) SetDelegateId(v string)`

SetDelegateId sets DelegateId field to given value.

### HasDelegateId

`func (o *Delegation2) HasDelegateId() bool`

HasDelegateId returns a boolean if a field has been set.

### GetExpiration

`func (o *Delegation2) GetExpiration() SailPointTime`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Delegation2) GetExpirationOk() (*SailPointTime, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Delegation2) SetExpiration(v SailPointTime)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Delegation2) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


