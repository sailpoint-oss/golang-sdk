---
id: nerm-profile1
title: Profile1
pagination_label: Profile1
sidebar_label: Profile1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Profile1', 'NERMProfile1'] 
slug: /tools/sdk/go/nerm/models/profile1
tags: ['SDK', 'Software Development Kit', 'Profile1', 'NERMProfile1']
---

# Profile1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The profile name. | [optional] 
**ProfileTypeId** | **string** | The profile type id. | 
**Status** | **string** | The profile status. | 
**IdProofingStatus** | Pointer to **string** | The id proofing status of the profile. | [optional] 
**Archived** | Pointer to **bool** | Describes whether the profile is archived or not. | [optional] [default to false]
**Attributes** | Pointer to **map[string]string** | The attributes associated with the profile. | [optional] 

## Methods

### NewProfile1

`func NewProfile1(profileTypeId string, status string, ) *Profile1`

NewProfile1 instantiates a new Profile1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfile1WithDefaults

`func NewProfile1WithDefaults() *Profile1`

NewProfile1WithDefaults instantiates a new Profile1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Profile1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Profile1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Profile1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Profile1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *Profile1) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *Profile1) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *Profile1) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.


### GetStatus

`func (o *Profile1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Profile1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Profile1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIdProofingStatus

`func (o *Profile1) GetIdProofingStatus() string`

GetIdProofingStatus returns the IdProofingStatus field if non-nil, zero value otherwise.

### GetIdProofingStatusOk

`func (o *Profile1) GetIdProofingStatusOk() (*string, bool)`

GetIdProofingStatusOk returns a tuple with the IdProofingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProofingStatus

`func (o *Profile1) SetIdProofingStatus(v string)`

SetIdProofingStatus sets IdProofingStatus field to given value.

### HasIdProofingStatus

`func (o *Profile1) HasIdProofingStatus() bool`

HasIdProofingStatus returns a boolean if a field has been set.

### GetArchived

`func (o *Profile1) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Profile1) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Profile1) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Profile1) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetAttributes

`func (o *Profile1) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Profile1) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Profile1) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Profile1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


