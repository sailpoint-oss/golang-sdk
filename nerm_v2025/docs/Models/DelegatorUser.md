---
id: nermv2025-delegator-user
title: DelegatorUser
pagination_label: DelegatorUser
sidebar_label: DelegatorUser
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DelegatorUser', 'NERMV2025DelegatorUser'] 
slug: /tools/sdk/go/nermv2025/models/delegator-user
tags: ['SDK', 'Software Development Kit', 'DelegatorUser', 'NERMV2025DelegatorUser']
---

# DelegatorUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the delegator user | [optional] 
**Uid** | Pointer to **string** | The uid of the delegator user | [optional] 
**Type** | Pointer to **string** | The type of the delegator user | [optional] 
**Name** | Pointer to **string** | The name of the delegator user | [optional] 
**Email** | Pointer to **string** | The email of the delegator user | [optional] 
**Status** | Pointer to **string** | The status of the delegator user | [optional] 
**Login** | Pointer to **string** | The login of the delegator user | [optional] 
**LastLogin** | Pointer to **SailPointTime** | The last login timestamp of the delegator user | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The date-time the record created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **SailPointTime** | The date-time the record was last updated. | [optional] [readonly] 

## Methods

### NewDelegatorUser

`func NewDelegatorUser() *DelegatorUser`

NewDelegatorUser instantiates a new DelegatorUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatorUserWithDefaults

`func NewDelegatorUserWithDefaults() *DelegatorUser`

NewDelegatorUserWithDefaults instantiates a new DelegatorUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DelegatorUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegatorUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegatorUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DelegatorUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *DelegatorUser) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DelegatorUser) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DelegatorUser) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DelegatorUser) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *DelegatorUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DelegatorUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DelegatorUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DelegatorUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DelegatorUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DelegatorUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DelegatorUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DelegatorUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *DelegatorUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DelegatorUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DelegatorUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DelegatorUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *DelegatorUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DelegatorUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DelegatorUser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DelegatorUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogin

`func (o *DelegatorUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DelegatorUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DelegatorUser) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *DelegatorUser) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetLastLogin

`func (o *DelegatorUser) GetLastLogin() SailPointTime`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *DelegatorUser) GetLastLoginOk() (*SailPointTime, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *DelegatorUser) SetLastLogin(v SailPointTime)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *DelegatorUser) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DelegatorUser) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DelegatorUser) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DelegatorUser) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DelegatorUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DelegatorUser) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DelegatorUser) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DelegatorUser) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DelegatorUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


