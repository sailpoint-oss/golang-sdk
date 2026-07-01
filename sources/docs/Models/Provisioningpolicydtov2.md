---
id: v1-provisioningpolicydtov2
title: Provisioningpolicydtov2
pagination_label: Provisioningpolicydtov2
sidebar_label: Provisioningpolicydtov2
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Provisioningpolicydtov2', 'V1Provisioningpolicydtov2'] 
slug: /tools/sdk/go/sources/models/provisioningpolicydtov2
tags: ['SDK', 'Software Development Kit', 'Provisioningpolicydtov2', 'V1Provisioningpolicydtov2']
---

# Provisioningpolicydtov2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | System-generated unique ID of the provisioning policy. | [optional] 
**Name** | **NullableString** | the provisioning policy name | 
**SubtypeId** | Pointer to **NullableString** | Subtype ID for which provisioning policy will be created when usageType is CREATE_MACHINE_ACCOUNT. | [optional] 
**Description** | Pointer to **string** | the description of the provisioning policy | [optional] 
**UsageType** | Pointer to **Usagetype** |  | [optional] 
**Fields** | Pointer to [**[]Fielddetailsdtov2**](fielddetailsdtov2) |  | [optional] 

## Methods

### NewProvisioningpolicydtov2

`func NewProvisioningpolicydtov2(name NullableString, ) *Provisioningpolicydtov2`

NewProvisioningpolicydtov2 instantiates a new Provisioningpolicydtov2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningpolicydtov2WithDefaults

`func NewProvisioningpolicydtov2WithDefaults() *Provisioningpolicydtov2`

NewProvisioningpolicydtov2WithDefaults instantiates a new Provisioningpolicydtov2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Provisioningpolicydtov2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Provisioningpolicydtov2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Provisioningpolicydtov2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Provisioningpolicydtov2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Provisioningpolicydtov2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provisioningpolicydtov2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provisioningpolicydtov2) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Provisioningpolicydtov2) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Provisioningpolicydtov2) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubtypeId

`func (o *Provisioningpolicydtov2) GetSubtypeId() string`

GetSubtypeId returns the SubtypeId field if non-nil, zero value otherwise.

### GetSubtypeIdOk

`func (o *Provisioningpolicydtov2) GetSubtypeIdOk() (*string, bool)`

GetSubtypeIdOk returns a tuple with the SubtypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtypeId

`func (o *Provisioningpolicydtov2) SetSubtypeId(v string)`

SetSubtypeId sets SubtypeId field to given value.

### HasSubtypeId

`func (o *Provisioningpolicydtov2) HasSubtypeId() bool`

HasSubtypeId returns a boolean if a field has been set.

### SetSubtypeIdNil

`func (o *Provisioningpolicydtov2) SetSubtypeIdNil(b bool)`

 SetSubtypeIdNil sets the value for SubtypeId to be an explicit nil

### UnsetSubtypeId
`func (o *Provisioningpolicydtov2) UnsetSubtypeId()`

UnsetSubtypeId ensures that no value is present for SubtypeId, not even an explicit nil
### GetDescription

`func (o *Provisioningpolicydtov2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Provisioningpolicydtov2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Provisioningpolicydtov2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Provisioningpolicydtov2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUsageType

`func (o *Provisioningpolicydtov2) GetUsageType() Usagetype`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Provisioningpolicydtov2) GetUsageTypeOk() (*Usagetype, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Provisioningpolicydtov2) SetUsageType(v Usagetype)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Provisioningpolicydtov2) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetFields

`func (o *Provisioningpolicydtov2) GetFields() []Fielddetailsdtov2`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Provisioningpolicydtov2) GetFieldsOk() (*[]Fielddetailsdtov2, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Provisioningpolicydtov2) SetFields(v []Fielddetailsdtov2)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Provisioningpolicydtov2) HasFields() bool`

HasFields returns a boolean if a field has been set.


