---
id: nerm-profile
title: Profile
pagination_label: Profile
sidebar_label: Profile
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Profile', 'NERMProfile'] 
slug: /tools/sdk/go/nerm/models/profile
tags: ['SDK', 'Software Development Kit', 'Profile', 'NERMProfile']
---

# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The objects ID | [optional] [readonly] 
**Uid** | Pointer to **string** | The objects UID | [optional] [readonly] 
**Name** | Pointer to **string** | This is the name of the profile. | [optional] 
**ProfileTypeId** | Pointer to **string** | This is the ID of the profile type the profile belongs to | [optional] 
**Status** | Pointer to **string** | This is the status of the profile | [optional] 
**IdProofingStatus** | Pointer to **string** | This is the ID proofing staus of the profile | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The date and time the profile was created | [optional] 
**UpdatedAt** | Pointer to **SailPointTime** | The date and time the profile was updated | [optional] 
**Attributes** | Pointer to **map[string]string** | Attributes that belong to this profile. | [optional] 

## Methods

### NewProfile

`func NewProfile() *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Profile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Profile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Profile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Profile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *Profile) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Profile) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Profile) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Profile) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *Profile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Profile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Profile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Profile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *Profile) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *Profile) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *Profile) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *Profile) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.

### GetStatus

`func (o *Profile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Profile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Profile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Profile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIdProofingStatus

`func (o *Profile) GetIdProofingStatus() string`

GetIdProofingStatus returns the IdProofingStatus field if non-nil, zero value otherwise.

### GetIdProofingStatusOk

`func (o *Profile) GetIdProofingStatusOk() (*string, bool)`

GetIdProofingStatusOk returns a tuple with the IdProofingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProofingStatus

`func (o *Profile) SetIdProofingStatus(v string)`

SetIdProofingStatus sets IdProofingStatus field to given value.

### HasIdProofingStatus

`func (o *Profile) HasIdProofingStatus() bool`

HasIdProofingStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Profile) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Profile) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Profile) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Profile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Profile) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Profile) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Profile) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Profile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttributes

`func (o *Profile) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Profile) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Profile) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Profile) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


