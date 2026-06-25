---
id: v1-intelhuman
title: Intelhuman
pagination_label: Intelhuman
sidebar_label: Intelhuman
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelhuman', 'V1Intelhuman'] 
slug: /tools/sdk/go/intelligencepackage/models/intelhuman
tags: ['SDK', 'Software Development Kit', 'Intelhuman', 'V1Intelhuman']
---

# Intelhuman

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Primary login or account alias for the human identity. | 
**Email** | **string** | Primary business email address for the human identity. | 
**IdentityStatus** | **string** | Current identity lifecycle status label from Identity Security Cloud. | 
**LifecycleState** | Pointer to **NullableString** | Lifecycle state name assigned through provisioning policy when present. | [optional] 
**ProcessingState** | Pointer to **NullableString** | Processing state for outstanding identity change operations when present. | [optional] 
**IsProtected** | **bool** | True when the identity is marked protected from automated changes. | 
**Manager** | Pointer to **NullableString** | Legacy manager identity identifier or display reference when assigned. | [optional] 
**ManagerId** | Pointer to **NullableString** | Manager identity identifier when correlated in Identity Security Cloud. | [optional] 
**ManagerName** | Pointer to **NullableString** | Manager display name when available from identity services. | [optional] 
**IsManager** | **bool** | True when the identity is flagged as a people manager in the organization. | 
**LastRefreshAt** | Pointer to **NullableTime** | Timestamp of the last successful identity refresh from sources when known. | [optional] 

## Methods

### NewIntelhuman

`func NewIntelhuman(alias string, email string, identityStatus string, isProtected bool, isManager bool, ) *Intelhuman`

NewIntelhuman instantiates a new Intelhuman object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelhumanWithDefaults

`func NewIntelhumanWithDefaults() *Intelhuman`

NewIntelhumanWithDefaults instantiates a new Intelhuman object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *Intelhuman) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Intelhuman) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Intelhuman) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetEmail

`func (o *Intelhuman) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Intelhuman) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Intelhuman) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIdentityStatus

`func (o *Intelhuman) GetIdentityStatus() string`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *Intelhuman) GetIdentityStatusOk() (*string, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *Intelhuman) SetIdentityStatus(v string)`

SetIdentityStatus sets IdentityStatus field to given value.


### GetLifecycleState

`func (o *Intelhuman) GetLifecycleState() string`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *Intelhuman) GetLifecycleStateOk() (*string, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *Intelhuman) SetLifecycleState(v string)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *Intelhuman) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### SetLifecycleStateNil

`func (o *Intelhuman) SetLifecycleStateNil(b bool)`

 SetLifecycleStateNil sets the value for LifecycleState to be an explicit nil

### UnsetLifecycleState
`func (o *Intelhuman) UnsetLifecycleState()`

UnsetLifecycleState ensures that no value is present for LifecycleState, not even an explicit nil
### GetProcessingState

`func (o *Intelhuman) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *Intelhuman) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *Intelhuman) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *Intelhuman) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### SetProcessingStateNil

`func (o *Intelhuman) SetProcessingStateNil(b bool)`

 SetProcessingStateNil sets the value for ProcessingState to be an explicit nil

### UnsetProcessingState
`func (o *Intelhuman) UnsetProcessingState()`

UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
### GetIsProtected

`func (o *Intelhuman) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *Intelhuman) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *Intelhuman) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.


### GetManager

`func (o *Intelhuman) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Intelhuman) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Intelhuman) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Intelhuman) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *Intelhuman) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *Intelhuman) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetManagerId

`func (o *Intelhuman) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *Intelhuman) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *Intelhuman) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *Intelhuman) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### SetManagerIdNil

`func (o *Intelhuman) SetManagerIdNil(b bool)`

 SetManagerIdNil sets the value for ManagerId to be an explicit nil

### UnsetManagerId
`func (o *Intelhuman) UnsetManagerId()`

UnsetManagerId ensures that no value is present for ManagerId, not even an explicit nil
### GetManagerName

`func (o *Intelhuman) GetManagerName() string`

GetManagerName returns the ManagerName field if non-nil, zero value otherwise.

### GetManagerNameOk

`func (o *Intelhuman) GetManagerNameOk() (*string, bool)`

GetManagerNameOk returns a tuple with the ManagerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerName

`func (o *Intelhuman) SetManagerName(v string)`

SetManagerName sets ManagerName field to given value.

### HasManagerName

`func (o *Intelhuman) HasManagerName() bool`

HasManagerName returns a boolean if a field has been set.

### SetManagerNameNil

`func (o *Intelhuman) SetManagerNameNil(b bool)`

 SetManagerNameNil sets the value for ManagerName to be an explicit nil

### UnsetManagerName
`func (o *Intelhuman) UnsetManagerName()`

UnsetManagerName ensures that no value is present for ManagerName, not even an explicit nil
### GetIsManager

`func (o *Intelhuman) GetIsManager() bool`

GetIsManager returns the IsManager field if non-nil, zero value otherwise.

### GetIsManagerOk

`func (o *Intelhuman) GetIsManagerOk() (*bool, bool)`

GetIsManagerOk returns a tuple with the IsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManager

`func (o *Intelhuman) SetIsManager(v bool)`

SetIsManager sets IsManager field to given value.


### GetLastRefreshAt

`func (o *Intelhuman) GetLastRefreshAt() SailPointTime`

GetLastRefreshAt returns the LastRefreshAt field if non-nil, zero value otherwise.

### GetLastRefreshAtOk

`func (o *Intelhuman) GetLastRefreshAtOk() (*SailPointTime, bool)`

GetLastRefreshAtOk returns a tuple with the LastRefreshAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshAt

`func (o *Intelhuman) SetLastRefreshAt(v SailPointTime)`

SetLastRefreshAt sets LastRefreshAt field to given value.

### HasLastRefreshAt

`func (o *Intelhuman) HasLastRefreshAt() bool`

HasLastRefreshAt returns a boolean if a field has been set.

### SetLastRefreshAtNil

`func (o *Intelhuman) SetLastRefreshAtNil(b bool)`

 SetLastRefreshAtNil sets the value for LastRefreshAt to be an explicit nil

### UnsetLastRefreshAt
`func (o *Intelhuman) UnsetLastRefreshAt()`

UnsetLastRefreshAt ensures that no value is present for LastRefreshAt, not even an explicit nil

