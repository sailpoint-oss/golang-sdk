---
id: v2025-machine-account
title: MachineAccount
pagination_label: MachineAccount
sidebar_label: MachineAccount
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccount', 'V2025MachineAccount'] 
slug: /tools/sdk/go/v2025/models/machine-account
tags: ['SDK', 'Software Development Kit', 'MachineAccount', 'V2025MachineAccount']
---

# MachineAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | System-generated unique ID of the Object | [optional] [readonly] 
**Name** | **NullableString** | Name of the Object | 
**Created** | Pointer to **SailPointTime** | Creation date of the Object | [optional] [readonly] 
**Modified** | Pointer to **SailPointTime** | Last modification date of the Object | [optional] [readonly] 
**Description** | Pointer to **NullableString** | A description of the machine account | [optional] 
**NativeIdentity** | **string** | The unique ID of the machine account generated by the source system | 
**Uuid** | Pointer to **NullableString** | The unique ID of the account as determined by the account schema | [optional] 
**ClassificationMethod** | **string** | Classification Method | 
**MachineIdentity** | Pointer to **map[string]interface{}** | The machine identity this account is associated with | [optional] 
**OwnerIdentity** | Pointer to **map[string]interface{}** | The identity who owns this account. | [optional] 
**AccessType** | Pointer to **string** | The connection type of the source this account is from | [optional] 
**Subtype** | Pointer to **NullableString** | The sub-type | [optional] 
**Environment** | Pointer to **NullableString** | Environment | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Custom attributes specific to the machine account | [optional] 
**ConnectorAttributes** | **map[string]interface{}** | The connector attributes for the account | 
**ManuallyCorrelated** | Pointer to **bool** | Indicates if the account has been manually correlated to an identity | [optional] [default to false]
**ManuallyEdited** | **bool** | Indicates if the account has been manually edited | [default to false]
**Locked** | **bool** | Indicates if the account is currently locked | 
**Enabled** | **bool** | Indicates if the account is enabled | [default to false]
**HasEntitlements** | **bool** | Indicates if the account has entitlements | [default to true]
**Source** | **map[string]interface{}** | The source this machine account belongs to. | 

## Methods

### NewMachineAccount

`func NewMachineAccount(name NullableString, nativeIdentity string, classificationMethod string, connectorAttributes map[string]interface{}, manuallyEdited bool, locked bool, enabled bool, hasEntitlements bool, source map[string]interface{}, ) *MachineAccount`

NewMachineAccount instantiates a new MachineAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAccountWithDefaults

`func NewMachineAccountWithDefaults() *MachineAccount`

NewMachineAccountWithDefaults instantiates a new MachineAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MachineAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MachineAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MachineAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineAccount) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *MachineAccount) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MachineAccount) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreated

`func (o *MachineAccount) GetCreated() SailPointTime`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MachineAccount) GetCreatedOk() (*SailPointTime, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MachineAccount) SetCreated(v SailPointTime)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MachineAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *MachineAccount) GetModified() SailPointTime`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *MachineAccount) GetModifiedOk() (*SailPointTime, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *MachineAccount) SetModified(v SailPointTime)`

SetModified sets Modified field to given value.

### HasModified

`func (o *MachineAccount) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDescription

`func (o *MachineAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MachineAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MachineAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MachineAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MachineAccount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MachineAccount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNativeIdentity

`func (o *MachineAccount) GetNativeIdentity() string`

GetNativeIdentity returns the NativeIdentity field if non-nil, zero value otherwise.

### GetNativeIdentityOk

`func (o *MachineAccount) GetNativeIdentityOk() (*string, bool)`

GetNativeIdentityOk returns a tuple with the NativeIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeIdentity

`func (o *MachineAccount) SetNativeIdentity(v string)`

SetNativeIdentity sets NativeIdentity field to given value.


### GetUuid

`func (o *MachineAccount) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MachineAccount) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MachineAccount) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MachineAccount) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *MachineAccount) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *MachineAccount) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetClassificationMethod

`func (o *MachineAccount) GetClassificationMethod() string`

GetClassificationMethod returns the ClassificationMethod field if non-nil, zero value otherwise.

### GetClassificationMethodOk

`func (o *MachineAccount) GetClassificationMethodOk() (*string, bool)`

GetClassificationMethodOk returns a tuple with the ClassificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationMethod

`func (o *MachineAccount) SetClassificationMethod(v string)`

SetClassificationMethod sets ClassificationMethod field to given value.


### GetMachineIdentity

`func (o *MachineAccount) GetMachineIdentity() map[string]interface{}`

GetMachineIdentity returns the MachineIdentity field if non-nil, zero value otherwise.

### GetMachineIdentityOk

`func (o *MachineAccount) GetMachineIdentityOk() (*map[string]interface{}, bool)`

GetMachineIdentityOk returns a tuple with the MachineIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineIdentity

`func (o *MachineAccount) SetMachineIdentity(v map[string]interface{})`

SetMachineIdentity sets MachineIdentity field to given value.

### HasMachineIdentity

`func (o *MachineAccount) HasMachineIdentity() bool`

HasMachineIdentity returns a boolean if a field has been set.

### GetOwnerIdentity

`func (o *MachineAccount) GetOwnerIdentity() map[string]interface{}`

GetOwnerIdentity returns the OwnerIdentity field if non-nil, zero value otherwise.

### GetOwnerIdentityOk

`func (o *MachineAccount) GetOwnerIdentityOk() (*map[string]interface{}, bool)`

GetOwnerIdentityOk returns a tuple with the OwnerIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdentity

`func (o *MachineAccount) SetOwnerIdentity(v map[string]interface{})`

SetOwnerIdentity sets OwnerIdentity field to given value.

### HasOwnerIdentity

`func (o *MachineAccount) HasOwnerIdentity() bool`

HasOwnerIdentity returns a boolean if a field has been set.

### SetOwnerIdentityNil

`func (o *MachineAccount) SetOwnerIdentityNil(b bool)`

 SetOwnerIdentityNil sets the value for OwnerIdentity to be an explicit nil

### UnsetOwnerIdentity
`func (o *MachineAccount) UnsetOwnerIdentity()`

UnsetOwnerIdentity ensures that no value is present for OwnerIdentity, not even an explicit nil
### GetAccessType

`func (o *MachineAccount) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *MachineAccount) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *MachineAccount) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *MachineAccount) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetSubtype

`func (o *MachineAccount) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MachineAccount) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MachineAccount) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *MachineAccount) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *MachineAccount) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *MachineAccount) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetEnvironment

`func (o *MachineAccount) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *MachineAccount) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *MachineAccount) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *MachineAccount) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *MachineAccount) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *MachineAccount) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetAttributes

`func (o *MachineAccount) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MachineAccount) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MachineAccount) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MachineAccount) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *MachineAccount) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *MachineAccount) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetConnectorAttributes

`func (o *MachineAccount) GetConnectorAttributes() map[string]interface{}`

GetConnectorAttributes returns the ConnectorAttributes field if non-nil, zero value otherwise.

### GetConnectorAttributesOk

`func (o *MachineAccount) GetConnectorAttributesOk() (*map[string]interface{}, bool)`

GetConnectorAttributesOk returns a tuple with the ConnectorAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorAttributes

`func (o *MachineAccount) SetConnectorAttributes(v map[string]interface{})`

SetConnectorAttributes sets ConnectorAttributes field to given value.


### SetConnectorAttributesNil

`func (o *MachineAccount) SetConnectorAttributesNil(b bool)`

 SetConnectorAttributesNil sets the value for ConnectorAttributes to be an explicit nil

### UnsetConnectorAttributes
`func (o *MachineAccount) UnsetConnectorAttributes()`

UnsetConnectorAttributes ensures that no value is present for ConnectorAttributes, not even an explicit nil
### GetManuallyCorrelated

`func (o *MachineAccount) GetManuallyCorrelated() bool`

GetManuallyCorrelated returns the ManuallyCorrelated field if non-nil, zero value otherwise.

### GetManuallyCorrelatedOk

`func (o *MachineAccount) GetManuallyCorrelatedOk() (*bool, bool)`

GetManuallyCorrelatedOk returns a tuple with the ManuallyCorrelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyCorrelated

`func (o *MachineAccount) SetManuallyCorrelated(v bool)`

SetManuallyCorrelated sets ManuallyCorrelated field to given value.

### HasManuallyCorrelated

`func (o *MachineAccount) HasManuallyCorrelated() bool`

HasManuallyCorrelated returns a boolean if a field has been set.

### GetManuallyEdited

`func (o *MachineAccount) GetManuallyEdited() bool`

GetManuallyEdited returns the ManuallyEdited field if non-nil, zero value otherwise.

### GetManuallyEditedOk

`func (o *MachineAccount) GetManuallyEditedOk() (*bool, bool)`

GetManuallyEditedOk returns a tuple with the ManuallyEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManuallyEdited

`func (o *MachineAccount) SetManuallyEdited(v bool)`

SetManuallyEdited sets ManuallyEdited field to given value.


### GetLocked

`func (o *MachineAccount) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MachineAccount) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MachineAccount) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetEnabled

`func (o *MachineAccount) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MachineAccount) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MachineAccount) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasEntitlements

`func (o *MachineAccount) GetHasEntitlements() bool`

GetHasEntitlements returns the HasEntitlements field if non-nil, zero value otherwise.

### GetHasEntitlementsOk

`func (o *MachineAccount) GetHasEntitlementsOk() (*bool, bool)`

GetHasEntitlementsOk returns a tuple with the HasEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntitlements

`func (o *MachineAccount) SetHasEntitlements(v bool)`

SetHasEntitlements sets HasEntitlements field to given value.


### GetSource

`func (o *MachineAccount) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MachineAccount) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MachineAccount) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.



