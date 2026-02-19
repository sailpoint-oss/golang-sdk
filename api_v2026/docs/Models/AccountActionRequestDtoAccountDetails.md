---
id: v2026-account-action-request-dto-account-details
title: AccountActionRequestDtoAccountDetails
pagination_label: AccountActionRequestDtoAccountDetails
sidebar_label: AccountActionRequestDtoAccountDetails
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AccountActionRequestDtoAccountDetails', 'V2026AccountActionRequestDtoAccountDetails'] 
slug: /tools/sdk/go/v2026/models/account-action-request-dto-account-details
tags: ['SDK', 'Software Development Kit', 'AccountActionRequestDtoAccountDetails', 'V2026AccountActionRequestDtoAccountDetails']
---

# AccountActionRequestDtoAccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | ID of account | [optional] 
**AccountName** | Pointer to **string** | Account name | [optional] 
**AccountNativeIdentity** | Pointer to **string** | Native identity of account | [optional] 
**AccountUuid** | Pointer to **string** | UUID associated with account | [optional] 
**AccountType** | Pointer to **string** | Type of account | [optional] 
**AccountSubtypeId** | Pointer to **NullableString** | Sub Type ID of account | [optional] 
**AccountSubtype** | Pointer to **NullableString** | Subtype of account | [optional] 
**Description** | Pointer to **NullableString** | Account Description | [optional] 
**SourceId** | Pointer to **string** | ID of source | [optional] 
**SourceName** | Pointer to **string** | Name of source | [optional] 
**HasEntitlements** | Pointer to **bool** | Indicates entitlements assigned to identity or not | [optional] [default to false]
**Disabled** | Pointer to **bool** | Indicates account is enabled/disabled | [optional] [default to false]
**Locked** | Pointer to **bool** | Indicates account locked/unlocked | [optional] [default to false]
**OwnerIdentity** | Pointer to [**NullableAccountDetailsOwnerIdentity**](account-details-owner-identity) |  | [optional] 

## Methods

### NewAccountActionRequestDtoAccountDetails

`func NewAccountActionRequestDtoAccountDetails() *AccountActionRequestDtoAccountDetails`

NewAccountActionRequestDtoAccountDetails instantiates a new AccountActionRequestDtoAccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountActionRequestDtoAccountDetailsWithDefaults

`func NewAccountActionRequestDtoAccountDetailsWithDefaults() *AccountActionRequestDtoAccountDetails`

NewAccountActionRequestDtoAccountDetailsWithDefaults instantiates a new AccountActionRequestDtoAccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountActionRequestDtoAccountDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountActionRequestDtoAccountDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountActionRequestDtoAccountDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountActionRequestDtoAccountDetails) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *AccountActionRequestDtoAccountDetails) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AccountActionRequestDtoAccountDetails) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AccountActionRequestDtoAccountDetails) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AccountActionRequestDtoAccountDetails) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNativeIdentity

`func (o *AccountActionRequestDtoAccountDetails) GetAccountNativeIdentity() string`

GetAccountNativeIdentity returns the AccountNativeIdentity field if non-nil, zero value otherwise.

### GetAccountNativeIdentityOk

`func (o *AccountActionRequestDtoAccountDetails) GetAccountNativeIdentityOk() (*string, bool)`

GetAccountNativeIdentityOk returns a tuple with the AccountNativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNativeIdentity

`func (o *AccountActionRequestDtoAccountDetails) SetAccountNativeIdentity(v string)`

SetAccountNativeIdentity sets AccountNativeIdentity field to given value.

### HasAccountNativeIdentity

`func (o *AccountActionRequestDtoAccountDetails) HasAccountNativeIdentity() bool`

HasAccountNativeIdentity returns a boolean if a field has been set.

### GetAccountUuid

`func (o *AccountActionRequestDtoAccountDetails) GetAccountUuid() string`

GetAccountUuid returns the AccountUuid field if non-nil, zero value otherwise.

### GetAccountUuidOk

`func (o *AccountActionRequestDtoAccountDetails) GetAccountUuidOk() (*string, bool)`

GetAccountUuidOk returns a tuple with the AccountUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUuid

`func (o *AccountActionRequestDtoAccountDetails) SetAccountUuid(v string)`

SetAccountUuid sets AccountUuid field to given value.

### HasAccountUuid

`func (o *AccountActionRequestDtoAccountDetails) HasAccountUuid() bool`

HasAccountUuid returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountActionRequestDtoAccountDetails) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountActionRequestDtoAccountDetails) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountActionRequestDtoAccountDetails) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountActionRequestDtoAccountDetails) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccountSubtypeId

`func (o *AccountActionRequestDtoAccountDetails) GetAccountSubtypeId() string`

GetAccountSubtypeId returns the AccountSubtypeId field if non-nil, zero value otherwise.

### GetAccountSubtypeIdOk

`func (o *AccountActionRequestDtoAccountDetails) GetAccountSubtypeIdOk() (*string, bool)`

GetAccountSubtypeIdOk returns a tuple with the AccountSubtypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtypeId

`func (o *AccountActionRequestDtoAccountDetails) SetAccountSubtypeId(v string)`

SetAccountSubtypeId sets AccountSubtypeId field to given value.

### HasAccountSubtypeId

`func (o *AccountActionRequestDtoAccountDetails) HasAccountSubtypeId() bool`

HasAccountSubtypeId returns a boolean if a field has been set.

### SetAccountSubtypeIdNil

`func (o *AccountActionRequestDtoAccountDetails) SetAccountSubtypeIdNil(b bool)`

 SetAccountSubtypeIdNil sets the value for AccountSubtypeId to be an explicit nil

### UnsetAccountSubtypeId
`func (o *AccountActionRequestDtoAccountDetails) UnsetAccountSubtypeId()`

UnsetAccountSubtypeId ensures that no value is present for AccountSubtypeId, not even an explicit nil
### GetAccountSubtype

`func (o *AccountActionRequestDtoAccountDetails) GetAccountSubtype() string`

GetAccountSubtype returns the AccountSubtype field if non-nil, zero value otherwise.

### GetAccountSubtypeOk

`func (o *AccountActionRequestDtoAccountDetails) GetAccountSubtypeOk() (*string, bool)`

GetAccountSubtypeOk returns a tuple with the AccountSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtype

`func (o *AccountActionRequestDtoAccountDetails) SetAccountSubtype(v string)`

SetAccountSubtype sets AccountSubtype field to given value.

### HasAccountSubtype

`func (o *AccountActionRequestDtoAccountDetails) HasAccountSubtype() bool`

HasAccountSubtype returns a boolean if a field has been set.

### SetAccountSubtypeNil

`func (o *AccountActionRequestDtoAccountDetails) SetAccountSubtypeNil(b bool)`

 SetAccountSubtypeNil sets the value for AccountSubtype to be an explicit nil

### UnsetAccountSubtype
`func (o *AccountActionRequestDtoAccountDetails) UnsetAccountSubtype()`

UnsetAccountSubtype ensures that no value is present for AccountSubtype, not even an explicit nil
### GetDescription

`func (o *AccountActionRequestDtoAccountDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountActionRequestDtoAccountDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountActionRequestDtoAccountDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountActionRequestDtoAccountDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccountActionRequestDtoAccountDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccountActionRequestDtoAccountDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSourceId

`func (o *AccountActionRequestDtoAccountDetails) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AccountActionRequestDtoAccountDetails) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AccountActionRequestDtoAccountDetails) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AccountActionRequestDtoAccountDetails) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *AccountActionRequestDtoAccountDetails) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AccountActionRequestDtoAccountDetails) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AccountActionRequestDtoAccountDetails) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *AccountActionRequestDtoAccountDetails) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetHasEntitlements

`func (o *AccountActionRequestDtoAccountDetails) GetHasEntitlements() bool`

GetHasEntitlements returns the HasEntitlements field if non-nil, zero value otherwise.

### GetHasEntitlementsOk

`func (o *AccountActionRequestDtoAccountDetails) GetHasEntitlementsOk() (*bool, bool)`

GetHasEntitlementsOk returns a tuple with the HasEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitlements

`func (o *AccountActionRequestDtoAccountDetails) SetHasEntitlements(v bool)`

SetHasEntitlements sets HasEntitlements field to given value.

### HasHasEntitlements

`func (o *AccountActionRequestDtoAccountDetails) HasHasEntitlements() bool`

HasHasEntitlements returns a boolean if a field has been set.

### GetDisabled

`func (o *AccountActionRequestDtoAccountDetails) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AccountActionRequestDtoAccountDetails) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AccountActionRequestDtoAccountDetails) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AccountActionRequestDtoAccountDetails) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetLocked

`func (o *AccountActionRequestDtoAccountDetails) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *AccountActionRequestDtoAccountDetails) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *AccountActionRequestDtoAccountDetails) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *AccountActionRequestDtoAccountDetails) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetOwnerIdentity

`func (o *AccountActionRequestDtoAccountDetails) GetOwnerIdentity() AccountDetailsOwnerIdentity`

GetOwnerIdentity returns the OwnerIdentity field if non-nil, zero value otherwise.

### GetOwnerIdentityOk

`func (o *AccountActionRequestDtoAccountDetails) GetOwnerIdentityOk() (*AccountDetailsOwnerIdentity, bool)`

GetOwnerIdentityOk returns a tuple with the OwnerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdentity

`func (o *AccountActionRequestDtoAccountDetails) SetOwnerIdentity(v AccountDetailsOwnerIdentity)`

SetOwnerIdentity sets OwnerIdentity field to given value.

### HasOwnerIdentity

`func (o *AccountActionRequestDtoAccountDetails) HasOwnerIdentity() bool`

HasOwnerIdentity returns a boolean if a field has been set.

### SetOwnerIdentityNil

`func (o *AccountActionRequestDtoAccountDetails) SetOwnerIdentityNil(b bool)`

 SetOwnerIdentityNil sets the value for OwnerIdentity to be an explicit nil

### UnsetOwnerIdentity
`func (o *AccountActionRequestDtoAccountDetails) UnsetOwnerIdentity()`

UnsetOwnerIdentity ensures that no value is present for OwnerIdentity, not even an explicit nil

