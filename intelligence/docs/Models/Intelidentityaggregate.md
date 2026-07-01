---
id: v1-intelidentityaggregate
title: Intelidentityaggregate
pagination_label: Intelidentityaggregate
sidebar_label: Intelidentityaggregate
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentityaggregate', 'V1Intelidentityaggregate'] 
slug: /tools/sdk/go/intelligence/models/intelidentityaggregate
tags: ['SDK', 'Software Development Kit', 'Intelidentityaggregate', 'V1Intelidentityaggregate']
---

# Intelidentityaggregate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identity Security Cloud identifier for this identity. | 
**Type** | **string** | Identity type for the matched record. | 
**DisplayName** | Pointer to **string** | Preferred display name for the identity across administrative experiences. | [optional] 
**Description** | Pointer to **NullableString** | Optional free-text description assigned to the identity profile when present. | [optional] 
**Subtype** | Pointer to **NullableString** | NERM classification for the identity. | [optional] 
**Owners** | Pointer to **NullableString** | Serialized owner reference information when populated by upstream identity services. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Arbitrary SCIM-style attribute bag returned for the identity context view. | [optional] 
**Created** | Pointer to **SailPointTime** | Timestamp when the identity record was created in Identity Security Cloud. | [optional] 
**Modified** | Pointer to **SailPointTime** | Timestamp when the identity record was last modified in Identity Security Cloud. | [optional] 
**Alias** | Pointer to **string** | Primary login or account alias for the identity. | [optional] 
**Email** | Pointer to **string** | Primary business email address for the identity. | [optional] 
**IdentityStatus** | Pointer to **string** | Current identity lifecycle status label from Identity Security Cloud. | [optional] 
**IsManager** | **bool** | True when the identity is flagged as a people manager in the organization. | 
**Accounts** | [**Intelaccountsslice**](intelaccountsslice) | First page of accounts for the identity. | 
**PrivilegedAccess** | [**Intelprivilegedaccessslice**](intelprivilegedaccessslice) | Full privileged access result for the identity. | 
**Outliers** | Pointer to [**Inteloutliersslice**](inteloutliersslice) | Rare access slice; omitted when the tenant lacks the IDA-outliers license. | [optional] 
**AccessHistory** | [**Intelaccesshistory**](intelaccesshistory) | Access-history split into access items and certifications sub-slices. | 

## Methods

### NewIntelidentityaggregate

`func NewIntelidentityaggregate(id string, type_ string, isManager bool, accounts Intelaccountsslice, privilegedAccess Intelprivilegedaccessslice, accessHistory Intelaccesshistory, ) *Intelidentityaggregate`

NewIntelidentityaggregate instantiates a new Intelidentityaggregate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentityaggregateWithDefaults

`func NewIntelidentityaggregateWithDefaults() *Intelidentityaggregate`

NewIntelidentityaggregateWithDefaults instantiates a new Intelidentityaggregate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Intelidentityaggregate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Intelidentityaggregate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Intelidentityaggregate) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Intelidentityaggregate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Intelidentityaggregate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Intelidentityaggregate) SetType(v string)`

SetType sets Type field to given value.


### GetDisplayName

`func (o *Intelidentityaggregate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Intelidentityaggregate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Intelidentityaggregate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Intelidentityaggregate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *Intelidentityaggregate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Intelidentityaggregate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Intelidentityaggregate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Intelidentityaggregate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Intelidentityaggregate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Intelidentityaggregate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubtype

`func (o *Intelidentityaggregate) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Intelidentityaggregate) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Intelidentityaggregate) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Intelidentityaggregate) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *Intelidentityaggregate) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *Intelidentityaggregate) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetOwners

`func (o *Intelidentityaggregate) GetOwners() string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *Intelidentityaggregate) GetOwnersOk() (*string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *Intelidentityaggregate) SetOwners(v string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *Intelidentityaggregate) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### SetOwnersNil

`func (o *Intelidentityaggregate) SetOwnersNil(b bool)`

 SetOwnersNil sets the value for Owners to be an explicit nil

### UnsetOwners
`func (o *Intelidentityaggregate) UnsetOwners()`

UnsetOwners ensures that no value is present for Owners, not even an explicit nil
### GetAttributes

`func (o *Intelidentityaggregate) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Intelidentityaggregate) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Intelidentityaggregate) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Intelidentityaggregate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCreated

`func (o *Intelidentityaggregate) GetCreated() SailPointTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Intelidentityaggregate) GetCreatedOk() (*SailPointTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Intelidentityaggregate) SetCreated(v SailPointTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Intelidentityaggregate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Intelidentityaggregate) GetModified() SailPointTime`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Intelidentityaggregate) GetModifiedOk() (*SailPointTime, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Intelidentityaggregate) SetModified(v SailPointTime)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Intelidentityaggregate) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetAlias

`func (o *Intelidentityaggregate) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Intelidentityaggregate) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Intelidentityaggregate) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Intelidentityaggregate) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetEmail

`func (o *Intelidentityaggregate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Intelidentityaggregate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Intelidentityaggregate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Intelidentityaggregate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIdentityStatus

`func (o *Intelidentityaggregate) GetIdentityStatus() string`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *Intelidentityaggregate) GetIdentityStatusOk() (*string, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *Intelidentityaggregate) SetIdentityStatus(v string)`

SetIdentityStatus sets IdentityStatus field to given value.

### HasIdentityStatus

`func (o *Intelidentityaggregate) HasIdentityStatus() bool`

HasIdentityStatus returns a boolean if a field has been set.

### GetIsManager

`func (o *Intelidentityaggregate) GetIsManager() bool`

GetIsManager returns the IsManager field if non-nil, zero value otherwise.

### GetIsManagerOk

`func (o *Intelidentityaggregate) GetIsManagerOk() (*bool, bool)`

GetIsManagerOk returns a tuple with the IsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManager

`func (o *Intelidentityaggregate) SetIsManager(v bool)`

SetIsManager sets IsManager field to given value.


### GetAccounts

`func (o *Intelidentityaggregate) GetAccounts() Intelaccountsslice`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Intelidentityaggregate) GetAccountsOk() (*Intelaccountsslice, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Intelidentityaggregate) SetAccounts(v Intelaccountsslice)`

SetAccounts sets Accounts field to given value.


### GetPrivilegedAccess

`func (o *Intelidentityaggregate) GetPrivilegedAccess() Intelprivilegedaccessslice`

GetPrivilegedAccess returns the PrivilegedAccess field if non-nil, zero value otherwise.

### GetPrivilegedAccessOk

`func (o *Intelidentityaggregate) GetPrivilegedAccessOk() (*Intelprivilegedaccessslice, bool)`

GetPrivilegedAccessOk returns a tuple with the PrivilegedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedAccess

`func (o *Intelidentityaggregate) SetPrivilegedAccess(v Intelprivilegedaccessslice)`

SetPrivilegedAccess sets PrivilegedAccess field to given value.


### GetOutliers

`func (o *Intelidentityaggregate) GetOutliers() Inteloutliersslice`

GetOutliers returns the Outliers field if non-nil, zero value otherwise.

### GetOutliersOk

`func (o *Intelidentityaggregate) GetOutliersOk() (*Inteloutliersslice, bool)`

GetOutliersOk returns a tuple with the Outliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutliers

`func (o *Intelidentityaggregate) SetOutliers(v Inteloutliersslice)`

SetOutliers sets Outliers field to given value.

### HasOutliers

`func (o *Intelidentityaggregate) HasOutliers() bool`

HasOutliers returns a boolean if a field has been set.

### GetAccessHistory

`func (o *Intelidentityaggregate) GetAccessHistory() Intelaccesshistory`

GetAccessHistory returns the AccessHistory field if non-nil, zero value otherwise.

### GetAccessHistoryOk

`func (o *Intelidentityaggregate) GetAccessHistoryOk() (*Intelaccesshistory, bool)`

GetAccessHistoryOk returns a tuple with the AccessHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessHistory

`func (o *Intelidentityaggregate) SetAccessHistory(v Intelaccesshistory)`

SetAccessHistory sets AccessHistory field to given value.



