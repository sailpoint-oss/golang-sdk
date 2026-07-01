---
id: v1-intelaccessaccountwire
title: Intelaccessaccountwire
pagination_label: Intelaccessaccountwire
sidebar_label: Intelaccessaccountwire
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelaccessaccountwire', 'V1Intelaccessaccountwire'] 
slug: /tools/sdk/go/intelligence/models/intelaccessaccountwire
tags: ['SDK', 'Software Development Kit', 'Intelaccessaccountwire', 'V1Intelaccessaccountwire']
---

# Intelaccessaccountwire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique account identifier in Identity Security Cloud. | 
**Name** | **string** | Account name or login value on the correlated source. | 
**Source** | Pointer to [**Intelaccesssourcewire**](intelaccesssourcewire) | Source metadata for the account as returned by List Accounts wire format. | [optional] 
**Disabled** | **bool** | True when the account is administratively disabled on the source. | 
**Locked** | **bool** | True when the account is locked from interactive sign-in on the source. | 
**Authoritative** | **bool** | True when the account is treated as authoritative for attribute synchronization. | 
**SystemAccount** | **bool** | True when the account represents a non-interactive or system principal. | 
**IsMachine** | **bool** | True when the account belongs to a machine or service identity. | 
**ManuallyCorrelated** | **bool** | True when an administrator manually correlated the account to an identity. | 
**NativeIdentity** | Pointer to **NullableString** | Native identifier string on the source directory or application. | [optional] 
**Created** | **SailPointTime** | Timestamp when the account record was created in Identity Security Cloud. | 
**Modified** | **SailPointTime** | Timestamp when the account record was last modified in Identity Security Cloud. | 

## Methods

### NewIntelaccessaccountwire

`func NewIntelaccessaccountwire(id string, name string, disabled bool, locked bool, authoritative bool, systemAccount bool, isMachine bool, manuallyCorrelated bool, created SailPointTime, modified SailPointTime, ) *Intelaccessaccountwire`

NewIntelaccessaccountwire instantiates a new Intelaccessaccountwire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelaccessaccountwireWithDefaults

`func NewIntelaccessaccountwireWithDefaults() *Intelaccessaccountwire`

NewIntelaccessaccountwireWithDefaults instantiates a new Intelaccessaccountwire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Intelaccessaccountwire) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Intelaccessaccountwire) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Intelaccessaccountwire) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Intelaccessaccountwire) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Intelaccessaccountwire) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Intelaccessaccountwire) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *Intelaccessaccountwire) GetSource() Intelaccesssourcewire`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Intelaccessaccountwire) GetSourceOk() (*Intelaccesssourcewire, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Intelaccessaccountwire) SetSource(v Intelaccesssourcewire)`

SetSource sets Source field to given value.

### HasSource

`func (o *Intelaccessaccountwire) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDisabled

`func (o *Intelaccessaccountwire) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Intelaccessaccountwire) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Intelaccessaccountwire) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetLocked

`func (o *Intelaccessaccountwire) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Intelaccessaccountwire) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Intelaccessaccountwire) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetAuthoritative

`func (o *Intelaccessaccountwire) GetAuthoritative() bool`

GetAuthoritative returns the Authoritative field if non-nil, zero value otherwise.

### GetAuthoritativeOk

`func (o *Intelaccessaccountwire) GetAuthoritativeOk() (*bool, bool)`

GetAuthoritativeOk returns a tuple with the Authoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritative

`func (o *Intelaccessaccountwire) SetAuthoritative(v bool)`

SetAuthoritative sets Authoritative field to given value.


### GetSystemAccount

`func (o *Intelaccessaccountwire) GetSystemAccount() bool`

GetSystemAccount returns the SystemAccount field if non-nil, zero value otherwise.

### GetSystemAccountOk

`func (o *Intelaccessaccountwire) GetSystemAccountOk() (*bool, bool)`

GetSystemAccountOk returns a tuple with the SystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccount

`func (o *Intelaccessaccountwire) SetSystemAccount(v bool)`

SetSystemAccount sets SystemAccount field to given value.


### GetIsMachine

`func (o *Intelaccessaccountwire) GetIsMachine() bool`

GetIsMachine returns the IsMachine field if non-nil, zero value otherwise.

### GetIsMachineOk

`func (o *Intelaccessaccountwire) GetIsMachineOk() (*bool, bool)`

GetIsMachineOk returns a tuple with the IsMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMachine

`func (o *Intelaccessaccountwire) SetIsMachine(v bool)`

SetIsMachine sets IsMachine field to given value.


### GetManuallyCorrelated

`func (o *Intelaccessaccountwire) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *Intelaccessaccountwire) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *Intelaccessaccountwire) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.


### GetNativeIdentity

`func (o *Intelaccessaccountwire) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *Intelaccessaccountwire) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *Intelaccessaccountwire) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.

### HasNativeIdentity

`func (o *Intelaccessaccountwire) HasNativeIdentity() bool`

HasNativeIdentity returns a boolean if a field has been set.

### SetNativeIdentityNil

`func (o *Intelaccessaccountwire) SetNativeIdentityNil(b bool)`

 SetNativeIdentityNil sets the value for NativeIdentity to be an explicit nil

### UnsetNativeIdentity
`func (o *Intelaccessaccountwire) UnsetNativeIdentity()`

UnsetNativeIdentity ensures that no value is present for NativeIdentity, not even an explicit nil
### GetCreated

`func (o *Intelaccessaccountwire) GetCreated() SailPointTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Intelaccessaccountwire) GetCreatedOk() (*SailPointTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Intelaccessaccountwire) SetCreated(v SailPointTime)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Intelaccessaccountwire) GetModified() SailPointTime`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Intelaccessaccountwire) GetModifiedOk() (*SailPointTime, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Intelaccessaccountwire) SetModified(v SailPointTime)`

SetModified sets Modified field to given value.



