---
id: nermv2025-delegation
title: Delegation
pagination_label: Delegation
sidebar_label: Delegation
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Delegation', 'NERMV2025Delegation'] 
slug: /tools/sdk/go/nermv2025/models/delegation
tags: ['SDK', 'Software Development Kit', 'Delegation', 'NERMV2025Delegation']
---

# Delegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the delegation | [optional] 
**DelegatorId** | Pointer to **map[string]interface{}** | The id of the delegator user | [optional] 
**DelegateId** | Pointer to **map[string]interface{}** | The id of the delegate user | [optional] 
**Delegator** | Pointer to [**DelegatorUser**](delegator-user) |  | [optional] 
**Delegate** | Pointer to [**DelegateUser**](delegate-user) |  | [optional] 
**Expiration** | Pointer to **SailPointTime** | The expiration date of the delegation | [optional] 
**Expired** | Pointer to **bool** | Indicates if the delegation is expired | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The date-time the record created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **SailPointTime** | The date-time the record was last updated. | [optional] [readonly] 

## Methods

### NewDelegation

`func NewDelegation() *Delegation`

NewDelegation instantiates a new Delegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationWithDefaults

`func NewDelegationWithDefaults() *Delegation`

NewDelegationWithDefaults instantiates a new Delegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Delegation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Delegation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Delegation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Delegation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDelegatorId

`func (o *Delegation) GetDelegatorId() map[string]interface{}`

GetDelegatorId returns the DelegatorId field if non-nil, zero value otherwise.

### GetDelegatorIdOk

`func (o *Delegation) GetDelegatorIdOk() (*map[string]interface{}, bool)`

GetDelegatorIdOk returns a tuple with the DelegatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatorId

`func (o *Delegation) SetDelegatorId(v map[string]interface{})`

SetDelegatorId sets DelegatorId field to given value.

### HasDelegatorId

`func (o *Delegation) HasDelegatorId() bool`

HasDelegatorId returns a boolean if a field has been set.

### GetDelegateId

`func (o *Delegation) GetDelegateId() map[string]interface{}`

GetDelegateId returns the DelegateId field if non-nil, zero value otherwise.

### GetDelegateIdOk

`func (o *Delegation) GetDelegateIdOk() (*map[string]interface{}, bool)`

GetDelegateIdOk returns a tuple with the DelegateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateId

`func (o *Delegation) SetDelegateId(v map[string]interface{})`

SetDelegateId sets DelegateId field to given value.

### HasDelegateId

`func (o *Delegation) HasDelegateId() bool`

HasDelegateId returns a boolean if a field has been set.

### GetDelegator

`func (o *Delegation) GetDelegator() DelegatorUser`

GetDelegator returns the Delegator field if non-nil, zero value otherwise.

### GetDelegatorOk

`func (o *Delegation) GetDelegatorOk() (*DelegatorUser, bool)`

GetDelegatorOk returns a tuple with the Delegator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegator

`func (o *Delegation) SetDelegator(v DelegatorUser)`

SetDelegator sets Delegator field to given value.

### HasDelegator

`func (o *Delegation) HasDelegator() bool`

HasDelegator returns a boolean if a field has been set.

### GetDelegate

`func (o *Delegation) GetDelegate() DelegateUser`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *Delegation) GetDelegateOk() (*DelegateUser, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *Delegation) SetDelegate(v DelegateUser)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *Delegation) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetExpiration

`func (o *Delegation) GetExpiration() SailPointTime`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Delegation) GetExpirationOk() (*SailPointTime, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Delegation) SetExpiration(v SailPointTime)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Delegation) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpired

`func (o *Delegation) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *Delegation) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *Delegation) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *Delegation) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Delegation) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Delegation) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Delegation) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Delegation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Delegation) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Delegation) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Delegation) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Delegation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


