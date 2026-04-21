---
id: nerm-profile-type1
title: ProfileType1
pagination_label: ProfileType1
sidebar_label: ProfileType1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileType1', 'NERMProfileType1'] 
slug: /tools/sdk/go/nerm/models/profile-type1
tags: ['SDK', 'Software Development Kit', 'ProfileType1', 'NERMProfileType1']
---

# ProfileType1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | This is the name of the profile type. | [optional] 
**Category** | Pointer to **string** | This is the category the profile type falls into. | [optional] 
**BypassDupProtection** | Pointer to **bool** | Whether or not duplication protection is bypassed. | [optional] [default to false]
**Archived** | Pointer to **bool** | Whether or not the profile type is archived. | [optional] [default to false]
**PermittedRoleIds** | Pointer to **[]string** | The role ids that are permitted for this profile type. | [optional] 
**IscSynced** | Pointer to **bool** | Is this profile type synced with ics | [optional] [default to false]
**ProfileTypeDupAttributesAttributes** | Pointer to [**[]ProfileType1ProfileTypeDupAttributesAttributesInner**](profile-type1-profile-type-dup-attributes-attributes-inner) |  | [optional] 
**ProfileTypeNamingsAttributes** | Pointer to [**[]ProfileType1ProfileTypeNamingsAttributesInner**](profile-type1-profile-type-namings-attributes-inner) |  | [optional] 

## Methods

### NewProfileType1

`func NewProfileType1() *ProfileType1`

NewProfileType1 instantiates a new ProfileType1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileType1WithDefaults

`func NewProfileType1WithDefaults() *ProfileType1`

NewProfileType1WithDefaults instantiates a new ProfileType1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProfileType1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfileType1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfileType1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProfileType1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *ProfileType1) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ProfileType1) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ProfileType1) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ProfileType1) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetBypassDupProtection

`func (o *ProfileType1) GetBypassDupProtection() bool`

GetBypassDupProtection returns the BypassDupProtection field if non-nil, zero value otherwise.

### GetBypassDupProtectionOk

`func (o *ProfileType1) GetBypassDupProtectionOk() (*bool, bool)`

GetBypassDupProtectionOk returns a tuple with the BypassDupProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassDupProtection

`func (o *ProfileType1) SetBypassDupProtection(v bool)`

SetBypassDupProtection sets BypassDupProtection field to given value.

### HasBypassDupProtection

`func (o *ProfileType1) HasBypassDupProtection() bool`

HasBypassDupProtection returns a boolean if a field has been set.

### GetArchived

`func (o *ProfileType1) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ProfileType1) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ProfileType1) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ProfileType1) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetPermittedRoleIds

`func (o *ProfileType1) GetPermittedRoleIds() []string`

GetPermittedRoleIds returns the PermittedRoleIds field if non-nil, zero value otherwise.

### GetPermittedRoleIdsOk

`func (o *ProfileType1) GetPermittedRoleIdsOk() (*[]string, bool)`

GetPermittedRoleIdsOk returns a tuple with the PermittedRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermittedRoleIds

`func (o *ProfileType1) SetPermittedRoleIds(v []string)`

SetPermittedRoleIds sets PermittedRoleIds field to given value.

### HasPermittedRoleIds

`func (o *ProfileType1) HasPermittedRoleIds() bool`

HasPermittedRoleIds returns a boolean if a field has been set.

### GetIscSynced

`func (o *ProfileType1) GetIscSynced() bool`

GetIscSynced returns the IscSynced field if non-nil, zero value otherwise.

### GetIscSyncedOk

`func (o *ProfileType1) GetIscSyncedOk() (*bool, bool)`

GetIscSyncedOk returns a tuple with the IscSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscSynced

`func (o *ProfileType1) SetIscSynced(v bool)`

SetIscSynced sets IscSynced field to given value.

### HasIscSynced

`func (o *ProfileType1) HasIscSynced() bool`

HasIscSynced returns a boolean if a field has been set.

### GetProfileTypeDupAttributesAttributes

`func (o *ProfileType1) GetProfileTypeDupAttributesAttributes() []ProfileType1ProfileTypeDupAttributesAttributesInner`

GetProfileTypeDupAttributesAttributes returns the ProfileTypeDupAttributesAttributes field if non-nil, zero value otherwise.

### GetProfileTypeDupAttributesAttributesOk

`func (o *ProfileType1) GetProfileTypeDupAttributesAttributesOk() (*[]ProfileType1ProfileTypeDupAttributesAttributesInner, bool)`

GetProfileTypeDupAttributesAttributesOk returns a tuple with the ProfileTypeDupAttributesAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeDupAttributesAttributes

`func (o *ProfileType1) SetProfileTypeDupAttributesAttributes(v []ProfileType1ProfileTypeDupAttributesAttributesInner)`

SetProfileTypeDupAttributesAttributes sets ProfileTypeDupAttributesAttributes field to given value.

### HasProfileTypeDupAttributesAttributes

`func (o *ProfileType1) HasProfileTypeDupAttributesAttributes() bool`

HasProfileTypeDupAttributesAttributes returns a boolean if a field has been set.

### GetProfileTypeNamingsAttributes

`func (o *ProfileType1) GetProfileTypeNamingsAttributes() []ProfileType1ProfileTypeNamingsAttributesInner`

GetProfileTypeNamingsAttributes returns the ProfileTypeNamingsAttributes field if non-nil, zero value otherwise.

### GetProfileTypeNamingsAttributesOk

`func (o *ProfileType1) GetProfileTypeNamingsAttributesOk() (*[]ProfileType1ProfileTypeNamingsAttributesInner, bool)`

GetProfileTypeNamingsAttributesOk returns a tuple with the ProfileTypeNamingsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeNamingsAttributes

`func (o *ProfileType1) SetProfileTypeNamingsAttributes(v []ProfileType1ProfileTypeNamingsAttributesInner)`

SetProfileTypeNamingsAttributes sets ProfileTypeNamingsAttributes field to given value.

### HasProfileTypeNamingsAttributes

`func (o *ProfileType1) HasProfileTypeNamingsAttributes() bool`

HasProfileTypeNamingsAttributes returns a boolean if a field has been set.


