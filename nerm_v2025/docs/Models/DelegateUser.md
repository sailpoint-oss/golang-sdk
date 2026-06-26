---
id: nermv2025-delegate-user
title: DelegateUser
pagination_label: DelegateUser
sidebar_label: DelegateUser
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DelegateUser', 'NERMV2025DelegateUser'] 
slug: /tools/sdk/go/nermv2025/models/delegate-user
tags: ['SDK', 'Software Development Kit', 'DelegateUser', 'NERMV2025DelegateUser']
---

# DelegateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the delegate user | [optional] 
**Uid** | Pointer to **string** | The uid of the delegate user | [optional] 
**Type** | Pointer to **string** | The type of the delegate user | [optional] 
**Name** | Pointer to **string** | The name of the delegate user | [optional] 
**Email** | Pointer to **string** | The email of the delegate user | [optional] 
**Status** | Pointer to **string** | The status of the delegate user | [optional] 
**Login** | Pointer to **string** | The login of the delegate user | [optional] 
**LastLogin** | Pointer to **SailPointTime** | The last login timestamp of the delegate user | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The date-time the record created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **SailPointTime** | The date-time the record was last updated. | [optional] [readonly] 

## Methods

### NewDelegateUser

`func NewDelegateUser() *DelegateUser`

NewDelegateUser instantiates a new DelegateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegateUserWithDefaults

`func NewDelegateUserWithDefaults() *DelegateUser`

NewDelegateUserWithDefaults instantiates a new DelegateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DelegateUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegateUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegateUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DelegateUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *DelegateUser) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DelegateUser) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DelegateUser) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DelegateUser) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *DelegateUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DelegateUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DelegateUser) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DelegateUser) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DelegateUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DelegateUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DelegateUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DelegateUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *DelegateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DelegateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DelegateUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DelegateUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *DelegateUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DelegateUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DelegateUser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DelegateUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogin

`func (o *DelegateUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DelegateUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DelegateUser) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *DelegateUser) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetLastLogin

`func (o *DelegateUser) GetLastLogin() SailPointTime`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *DelegateUser) GetLastLoginOk() (*SailPointTime, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *DelegateUser) SetLastLogin(v SailPointTime)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *DelegateUser) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DelegateUser) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DelegateUser) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DelegateUser) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DelegateUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DelegateUser) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DelegateUser) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DelegateUser) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DelegateUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


