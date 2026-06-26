---
id: nerm-profile-type
title: ProfileType
pagination_label: ProfileType
sidebar_label: ProfileType
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileType', 'NERMProfileType'] 
slug: /tools/sdk/go/nerm/models/profile-type
tags: ['SDK', 'Software Development Kit', 'ProfileType', 'NERMProfileType']
---

# ProfileType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The objects ID. | [optional] [readonly] 
**Uid** | Pointer to **string** | The objects UID. | [optional] [readonly] 
**Name** | Pointer to **string** | This is the name of the profile type. | [optional] 
**Category** | Pointer to **string** | This is the category the profile type falls into. | [optional] 
**BypassDupProtection** | Pointer to **bool** | Whether or not duplication protection is bypassed. | [optional] 
**Archived** | Pointer to **bool** | Whether or not the profile type is archived. | [optional] 
**PermittedRoleIds** | Pointer to **[]string** | The role ids that are permitted for this profile type. | [optional] 
**IscSynced** | Pointer to **bool** | Is this profile type synced with ics | [optional] 
**ProfileTypeDupAttributes** | Pointer to [**[]ProfileTypeProfileTypeDupAttributesInner**](profile-type-profile-type-dup-attributes-inner) |  | [optional] 
**ProfileTypeNamings** | Pointer to [**[]ProfileTypeProfileTypeNamingsInner**](profile-type-profile-type-namings-inner) |  | [optional] 

## Methods

### NewProfileType

`func NewProfileType() *ProfileType`

NewProfileType instantiates a new ProfileType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileTypeWithDefaults

`func NewProfileTypeWithDefaults() *ProfileType`

NewProfileTypeWithDefaults instantiates a new ProfileType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *ProfileType) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfileType) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfileType) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfileType) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetName

`func (o *ProfileType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfileType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfileType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProfileType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *ProfileType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ProfileType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ProfileType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ProfileType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetBypassDupProtection

`func (o *ProfileType) GetBypassDupProtection() bool`

GetBypassDupProtection returns the BypassDupProtection field if non-nil, zero value otherwise.

### GetBypassDupProtectionOk

`func (o *ProfileType) GetBypassDupProtectionOk() (*bool, bool)`

GetBypassDupProtectionOk returns a tuple with the BypassDupProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassDupProtection

`func (o *ProfileType) SetBypassDupProtection(v bool)`

SetBypassDupProtection sets BypassDupProtection field to given value.

### HasBypassDupProtection

`func (o *ProfileType) HasBypassDupProtection() bool`

HasBypassDupProtection returns a boolean if a field has been set.

### GetArchived

`func (o *ProfileType) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ProfileType) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ProfileType) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ProfileType) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetPermittedRoleIds

`func (o *ProfileType) GetPermittedRoleIds() []string`

GetPermittedRoleIds returns the PermittedRoleIds field if non-nil, zero value otherwise.

### GetPermittedRoleIdsOk

`func (o *ProfileType) GetPermittedRoleIdsOk() (*[]string, bool)`

GetPermittedRoleIdsOk returns a tuple with the PermittedRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedRoleIds

`func (o *ProfileType) SetPermittedRoleIds(v []string)`

SetPermittedRoleIds sets PermittedRoleIds field to given value.

### HasPermittedRoleIds

`func (o *ProfileType) HasPermittedRoleIds() bool`

HasPermittedRoleIds returns a boolean if a field has been set.

### GetIscSynced

`func (o *ProfileType) GetIscSynced() bool`

GetIscSynced returns the IscSynced field if non-nil, zero value otherwise.

### GetIscSyncedOk

`func (o *ProfileType) GetIscSyncedOk() (*bool, bool)`

GetIscSyncedOk returns a tuple with the IscSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscSynced

`func (o *ProfileType) SetIscSynced(v bool)`

SetIscSynced sets IscSynced field to given value.

### HasIscSynced

`func (o *ProfileType) HasIscSynced() bool`

HasIscSynced returns a boolean if a field has been set.

### GetProfileTypeDupAttributes

`func (o *ProfileType) GetProfileTypeDupAttributes() []ProfileTypeProfileTypeDupAttributesInner`

GetProfileTypeDupAttributes returns the ProfileTypeDupAttributes field if non-nil, zero value otherwise.

### GetProfileTypeDupAttributesOk

`func (o *ProfileType) GetProfileTypeDupAttributesOk() (*[]ProfileTypeProfileTypeDupAttributesInner, bool)`

GetProfileTypeDupAttributesOk returns a tuple with the ProfileTypeDupAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeDupAttributes

`func (o *ProfileType) SetProfileTypeDupAttributes(v []ProfileTypeProfileTypeDupAttributesInner)`

SetProfileTypeDupAttributes sets ProfileTypeDupAttributes field to given value.

### HasProfileTypeDupAttributes

`func (o *ProfileType) HasProfileTypeDupAttributes() bool`

HasProfileTypeDupAttributes returns a boolean if a field has been set.

### GetProfileTypeNamings

`func (o *ProfileType) GetProfileTypeNamings() []ProfileTypeProfileTypeNamingsInner`

GetProfileTypeNamings returns the ProfileTypeNamings field if non-nil, zero value otherwise.

### GetProfileTypeNamingsOk

`func (o *ProfileType) GetProfileTypeNamingsOk() (*[]ProfileTypeProfileTypeNamingsInner, bool)`

GetProfileTypeNamingsOk returns a tuple with the ProfileTypeNamings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeNamings

`func (o *ProfileType) SetProfileTypeNamings(v []ProfileTypeProfileTypeNamingsInner)`

SetProfileTypeNamings sets ProfileTypeNamings field to given value.

### HasProfileTypeNamings

`func (o *ProfileType) HasProfileTypeNamings() bool`

HasProfileTypeNamings returns a boolean if a field has been set.


