---
id: v1-intelidentityresponse
title: Intelidentityresponse
pagination_label: Intelidentityresponse
sidebar_label: Intelidentityresponse
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentityresponse', 'V1Intelidentityresponse'] 
slug: /tools/sdk/go/intelligencepackagev1/models/intelidentityresponse
tags: ['SDK', 'Software Development Kit', 'Intelidentityresponse', 'V1Intelidentityresponse']
---

# Intelidentityresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identity Security Cloud identifier for this identity. | 
**Type** | **string** | Discriminator indicating whether this identity is human or machine backed. | 
**DisplayName** | Pointer to **string** | Preferred display name for the identity across administrative experiences. | [optional] 
**Description** | Pointer to **NullableString** | Optional free-text description assigned to the identity profile when present. | [optional] 
**Subtype** | Pointer to **NullableString** | For HUMAN identities, NERM classification (Employee, Non Employee, or Cannot Determine). For MACHINE identities, connector subtype string from the authoritative source.  | [optional] 
**Owners** | Pointer to **NullableString** | Serialized owner reference information when populated by upstream identity services. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Arbitrary SCIM-style attribute bag returned for the identity context view. | [optional] 
**Timestamps** | [**Intelidentitytimestamps**](intelidentitytimestamps) | Created and modified timestamps for the identity record in Identity Security Cloud. | 
**Human** | Pointer to [**NullableIntelhuman**](intelhuman) | Human identity extension payload when type is HUMAN. | [optional] 
**Machine** | Pointer to [**NullableIntelmachine**](intelmachine) | Machine identity extension payload when type is MACHINE. | [optional] 
**Links** | Pointer to [**NullableIntelidentitylinks**](intelidentitylinks) | Hyperlinks to related Intelligence Package sub-resources; present for HUMAN only. | [optional] 

## Methods

### NewIntelidentityresponse

`func NewIntelidentityresponse(id string, type_ string, timestamps Intelidentitytimestamps, ) *Intelidentityresponse`

NewIntelidentityresponse instantiates a new Intelidentityresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentityresponseWithDefaults

`func NewIntelidentityresponseWithDefaults() *Intelidentityresponse`

NewIntelidentityresponseWithDefaults instantiates a new Intelidentityresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Intelidentityresponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Intelidentityresponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Intelidentityresponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Intelidentityresponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Intelidentityresponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Intelidentityresponse) SetType(v string)`

SetType sets Type field to given value.


### GetDisplayName

`func (o *Intelidentityresponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Intelidentityresponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Intelidentityresponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Intelidentityresponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *Intelidentityresponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Intelidentityresponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Intelidentityresponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Intelidentityresponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Intelidentityresponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Intelidentityresponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubtype

`func (o *Intelidentityresponse) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Intelidentityresponse) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Intelidentityresponse) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Intelidentityresponse) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *Intelidentityresponse) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *Intelidentityresponse) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetOwners

`func (o *Intelidentityresponse) GetOwners() string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *Intelidentityresponse) GetOwnersOk() (*string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *Intelidentityresponse) SetOwners(v string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *Intelidentityresponse) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### SetOwnersNil

`func (o *Intelidentityresponse) SetOwnersNil(b bool)`

 SetOwnersNil sets the value for Owners to be an explicit nil

### UnsetOwners
`func (o *Intelidentityresponse) UnsetOwners()`

UnsetOwners ensures that no value is present for Owners, not even an explicit nil
### GetAttributes

`func (o *Intelidentityresponse) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Intelidentityresponse) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Intelidentityresponse) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Intelidentityresponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetTimestamps

`func (o *Intelidentityresponse) GetTimestamps() Intelidentitytimestamps`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *Intelidentityresponse) GetTimestampsOk() (*Intelidentitytimestamps, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *Intelidentityresponse) SetTimestamps(v Intelidentitytimestamps)`

SetTimestamps sets Timestamps field to given value.


### GetHuman

`func (o *Intelidentityresponse) GetHuman() Intelhuman`

GetHuman returns the Human field if non-nil, zero value otherwise.

### GetHumanOk

`func (o *Intelidentityresponse) GetHumanOk() (*Intelhuman, bool)`

GetHumanOk returns a tuple with the Human field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuman

`func (o *Intelidentityresponse) SetHuman(v Intelhuman)`

SetHuman sets Human field to given value.

### HasHuman

`func (o *Intelidentityresponse) HasHuman() bool`

HasHuman returns a boolean if a field has been set.

### SetHumanNil

`func (o *Intelidentityresponse) SetHumanNil(b bool)`

 SetHumanNil sets the value for Human to be an explicit nil

### UnsetHuman
`func (o *Intelidentityresponse) UnsetHuman()`

UnsetHuman ensures that no value is present for Human, not even an explicit nil
### GetMachine

`func (o *Intelidentityresponse) GetMachine() Intelmachine`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *Intelidentityresponse) GetMachineOk() (*Intelmachine, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *Intelidentityresponse) SetMachine(v Intelmachine)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *Intelidentityresponse) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### SetMachineNil

`func (o *Intelidentityresponse) SetMachineNil(b bool)`

 SetMachineNil sets the value for Machine to be an explicit nil

### UnsetMachine
`func (o *Intelidentityresponse) UnsetMachine()`

UnsetMachine ensures that no value is present for Machine, not even an explicit nil
### GetLinks

`func (o *Intelidentityresponse) GetLinks() Intelidentitylinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Intelidentityresponse) GetLinksOk() (*Intelidentitylinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Intelidentityresponse) SetLinks(v Intelidentitylinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Intelidentityresponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *Intelidentityresponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *Intelidentityresponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

