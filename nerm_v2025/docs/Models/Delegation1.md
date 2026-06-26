---
id: nermv2025-delegation1
title: Delegation1
pagination_label: Delegation1
sidebar_label: Delegation1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Delegation1', 'NERMV2025Delegation1'] 
slug: /tools/sdk/go/nermv2025/models/delegation1
tags: ['SDK', 'Software Development Kit', 'Delegation1', 'NERMV2025Delegation1']
---

# Delegation1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the delegation | [optional] 
**DelegatorId** | Pointer to **map[string]interface{}** | The id of the delegator user | [optional] 
**DelegateId** | Pointer to **map[string]interface{}** | The id of the delegate user | [optional] 
**Expiration** | Pointer to **SailPointTime** | The expiration date of the delegation | [optional] 
**Expired** | Pointer to **bool** | Indicates if the delegation is expired | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The date-time the record created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **SailPointTime** | The date-time the record was last updated. | [optional] [readonly] 

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

### GetId

`func (o *Delegation1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Delegation1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Delegation1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Delegation1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDelegatorId

`func (o *Delegation1) GetDelegatorId() map[string]interface{}`

GetDelegatorId returns the DelegatorId field if non-nil, zero value otherwise.

### GetDelegatorIdOk

`func (o *Delegation1) GetDelegatorIdOk() (*map[string]interface{}, bool)`

GetDelegatorIdOk returns a tuple with the DelegatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorId

`func (o *Delegation1) SetDelegatorId(v map[string]interface{})`

SetDelegatorId sets DelegatorId field to given value.

### HasDelegatorId

`func (o *Delegation1) HasDelegatorId() bool`

HasDelegatorId returns a boolean if a field has been set.

### GetDelegateId

`func (o *Delegation1) GetDelegateId() map[string]interface{}`

GetDelegateId returns the DelegateId field if non-nil, zero value otherwise.

### GetDelegateIdOk

`func (o *Delegation1) GetDelegateIdOk() (*map[string]interface{}, bool)`

GetDelegateIdOk returns a tuple with the DelegateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateId

`func (o *Delegation1) SetDelegateId(v map[string]interface{})`

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

### GetExpired

`func (o *Delegation1) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *Delegation1) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *Delegation1) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *Delegation1) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Delegation1) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Delegation1) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Delegation1) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Delegation1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Delegation1) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Delegation1) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Delegation1) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Delegation1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


