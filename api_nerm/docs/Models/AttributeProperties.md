---
id: nerm-attribute-properties
title: AttributeProperties
pagination_label: AttributeProperties
sidebar_label: AttributeProperties
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AttributeProperties', 'NERMAttributeProperties'] 
slug: /tools/sdk/go/nerm/models/attribute-properties
tags: ['SDK', 'Software Development Kit', 'AttributeProperties', 'NERMAttributeProperties']
---

# AttributeProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the attribute | [optional] [readonly] 
**Uid** | Pointer to **string** | The user-specified identifier of the attribute | [optional] [readonly] 
**Label** | Pointer to **string** | The label for the attribute | [optional] 
**Description** | Pointer to **string** | A description of the attribute | [optional] 
**ToolTip** | Pointer to **string** | The helper text that accompanies the attribute | [optional] 
**Crypt** | Pointer to **bool** | Whether or not the attribute is encrypted | [optional] 
**Archived** | Pointer to **bool** | Whether the attribute is archived | [optional] 
**ArchivedOn** | Pointer to **SailPointTime** | When the attribute was archived | [optional] [readonly] 
**CreatedAt** | Pointer to **SailPointTime** | When the attribute was created | [optional] [readonly] 
**UpdatedAt** | Pointer to **SailPointTime** | When the attribute was last updated | [optional] [readonly] 
**DateFormat** | Pointer to **string** | The format of the date input if it is a date input | [optional] 
**SelectableStatus** | Pointer to **string** | The status of the profiles that can be selected | [optional] 
**RiskScoreSetting** | Pointer to **string** | What setting is used for the risk score | [optional] 
**RiskType** | Pointer to **string** | Type of risk that applies to the attribute | [optional] 
**OwnershipDriven** | Pointer to **bool** | Only shows profiles that the user currently has access to, to be selected | [optional] 
**AllowMultipleSelections** | Pointer to **bool** | Whether or not multiple selections can be made on something like a contributor search. | [optional] 
**FilteredByNeAttribute** | Pointer to **bool** | Whether or not the attribute is filtered by another attribute | [optional] 
**FilteringNeAttributeId** | Pointer to **string** | The ID of the filtering attribute | [optional] 
**NeAttributeFilterId** | Pointer to **string** | The ID of the attribute filter | [optional] 
**ReverseAssociationAttributeId** | Pointer to **string** | The ID of the attribute used with reverse association | [optional] 
**ProfileTypeId** | Pointer to **string** | The ID of the profile type the attribute applies to | [optional] 
**LegacyId** | Pointer to **string** | The legacy ID | [optional] 
**TmpCreatedAt** | Pointer to **SailPointTime** | the temp of when attribute was created | [optional] [readonly] 
**TmpUpdatedAt** | Pointer to **SailPointTime** | the temp of when attribute was last updated | [optional] [readonly] 

## Methods

### NewAttributeProperties

`func NewAttributeProperties() *AttributeProperties`

NewAttributeProperties instantiates a new AttributeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributePropertiesWithDefaults

`func NewAttributePropertiesWithDefaults() *AttributeProperties`

NewAttributePropertiesWithDefaults instantiates a new AttributeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttributeProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeProperties) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttributeProperties) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *AttributeProperties) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AttributeProperties) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AttributeProperties) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AttributeProperties) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLabel

`func (o *AttributeProperties) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AttributeProperties) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AttributeProperties) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AttributeProperties) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *AttributeProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttributeProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttributeProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttributeProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetToolTip

`func (o *AttributeProperties) GetToolTip() string`

GetToolTip returns the ToolTip field if non-nil, zero value otherwise.

### GetToolTipOk

`func (o *AttributeProperties) GetToolTipOk() (*string, bool)`

GetToolTipOk returns a tuple with the ToolTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolTip

`func (o *AttributeProperties) SetToolTip(v string)`

SetToolTip sets ToolTip field to given value.

### HasToolTip

`func (o *AttributeProperties) HasToolTip() bool`

HasToolTip returns a boolean if a field has been set.

### GetCrypt

`func (o *AttributeProperties) GetCrypt() bool`

GetCrypt returns the Crypt field if non-nil, zero value otherwise.

### GetCryptOk

`func (o *AttributeProperties) GetCryptOk() (*bool, bool)`

GetCryptOk returns a tuple with the Crypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrypt

`func (o *AttributeProperties) SetCrypt(v bool)`

SetCrypt sets Crypt field to given value.

### HasCrypt

`func (o *AttributeProperties) HasCrypt() bool`

HasCrypt returns a boolean if a field has been set.

### GetArchived

`func (o *AttributeProperties) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *AttributeProperties) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *AttributeProperties) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *AttributeProperties) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedOn

`func (o *AttributeProperties) GetArchivedOn() SailPointTime`

GetArchivedOn returns the ArchivedOn field if non-nil, zero value otherwise.

### GetArchivedOnOk

`func (o *AttributeProperties) GetArchivedOnOk() (*SailPointTime, bool)`

GetArchivedOnOk returns a tuple with the ArchivedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedOn

`func (o *AttributeProperties) SetArchivedOn(v SailPointTime)`

SetArchivedOn sets ArchivedOn field to given value.

### HasArchivedOn

`func (o *AttributeProperties) HasArchivedOn() bool`

HasArchivedOn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AttributeProperties) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttributeProperties) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttributeProperties) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AttributeProperties) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AttributeProperties) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AttributeProperties) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AttributeProperties) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AttributeProperties) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDateFormat

`func (o *AttributeProperties) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *AttributeProperties) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *AttributeProperties) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *AttributeProperties) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetSelectableStatus

`func (o *AttributeProperties) GetSelectableStatus() string`

GetSelectableStatus returns the SelectableStatus field if non-nil, zero value otherwise.

### GetSelectableStatusOk

`func (o *AttributeProperties) GetSelectableStatusOk() (*string, bool)`

GetSelectableStatusOk returns a tuple with the SelectableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectableStatus

`func (o *AttributeProperties) SetSelectableStatus(v string)`

SetSelectableStatus sets SelectableStatus field to given value.

### HasSelectableStatus

`func (o *AttributeProperties) HasSelectableStatus() bool`

HasSelectableStatus returns a boolean if a field has been set.

### GetRiskScoreSetting

`func (o *AttributeProperties) GetRiskScoreSetting() string`

GetRiskScoreSetting returns the RiskScoreSetting field if non-nil, zero value otherwise.

### GetRiskScoreSettingOk

`func (o *AttributeProperties) GetRiskScoreSettingOk() (*string, bool)`

GetRiskScoreSettingOk returns a tuple with the RiskScoreSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreSetting

`func (o *AttributeProperties) SetRiskScoreSetting(v string)`

SetRiskScoreSetting sets RiskScoreSetting field to given value.

### HasRiskScoreSetting

`func (o *AttributeProperties) HasRiskScoreSetting() bool`

HasRiskScoreSetting returns a boolean if a field has been set.

### GetRiskType

`func (o *AttributeProperties) GetRiskType() string`

GetRiskType returns the RiskType field if non-nil, zero value otherwise.

### GetRiskTypeOk

`func (o *AttributeProperties) GetRiskTypeOk() (*string, bool)`

GetRiskTypeOk returns a tuple with the RiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskType

`func (o *AttributeProperties) SetRiskType(v string)`

SetRiskType sets RiskType field to given value.

### HasRiskType

`func (o *AttributeProperties) HasRiskType() bool`

HasRiskType returns a boolean if a field has been set.

### GetOwnershipDriven

`func (o *AttributeProperties) GetOwnershipDriven() bool`

GetOwnershipDriven returns the OwnershipDriven field if non-nil, zero value otherwise.

### GetOwnershipDrivenOk

`func (o *AttributeProperties) GetOwnershipDrivenOk() (*bool, bool)`

GetOwnershipDrivenOk returns a tuple with the OwnershipDriven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipDriven

`func (o *AttributeProperties) SetOwnershipDriven(v bool)`

SetOwnershipDriven sets OwnershipDriven field to given value.

### HasOwnershipDriven

`func (o *AttributeProperties) HasOwnershipDriven() bool`

HasOwnershipDriven returns a boolean if a field has been set.

### GetAllowMultipleSelections

`func (o *AttributeProperties) GetAllowMultipleSelections() bool`

GetAllowMultipleSelections returns the AllowMultipleSelections field if non-nil, zero value otherwise.

### GetAllowMultipleSelectionsOk

`func (o *AttributeProperties) GetAllowMultipleSelectionsOk() (*bool, bool)`

GetAllowMultipleSelectionsOk returns a tuple with the AllowMultipleSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSelections

`func (o *AttributeProperties) SetAllowMultipleSelections(v bool)`

SetAllowMultipleSelections sets AllowMultipleSelections field to given value.

### HasAllowMultipleSelections

`func (o *AttributeProperties) HasAllowMultipleSelections() bool`

HasAllowMultipleSelections returns a boolean if a field has been set.

### GetFilteredByNeAttribute

`func (o *AttributeProperties) GetFilteredByNeAttribute() bool`

GetFilteredByNeAttribute returns the FilteredByNeAttribute field if non-nil, zero value otherwise.

### GetFilteredByNeAttributeOk

`func (o *AttributeProperties) GetFilteredByNeAttributeOk() (*bool, bool)`

GetFilteredByNeAttributeOk returns a tuple with the FilteredByNeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredByNeAttribute

`func (o *AttributeProperties) SetFilteredByNeAttribute(v bool)`

SetFilteredByNeAttribute sets FilteredByNeAttribute field to given value.

### HasFilteredByNeAttribute

`func (o *AttributeProperties) HasFilteredByNeAttribute() bool`

HasFilteredByNeAttribute returns a boolean if a field has been set.

### GetFilteringNeAttributeId

`func (o *AttributeProperties) GetFilteringNeAttributeId() string`

GetFilteringNeAttributeId returns the FilteringNeAttributeId field if non-nil, zero value otherwise.

### GetFilteringNeAttributeIdOk

`func (o *AttributeProperties) GetFilteringNeAttributeIdOk() (*string, bool)`

GetFilteringNeAttributeIdOk returns a tuple with the FilteringNeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringNeAttributeId

`func (o *AttributeProperties) SetFilteringNeAttributeId(v string)`

SetFilteringNeAttributeId sets FilteringNeAttributeId field to given value.

### HasFilteringNeAttributeId

`func (o *AttributeProperties) HasFilteringNeAttributeId() bool`

HasFilteringNeAttributeId returns a boolean if a field has been set.

### GetNeAttributeFilterId

`func (o *AttributeProperties) GetNeAttributeFilterId() string`

GetNeAttributeFilterId returns the NeAttributeFilterId field if non-nil, zero value otherwise.

### GetNeAttributeFilterIdOk

`func (o *AttributeProperties) GetNeAttributeFilterIdOk() (*string, bool)`

GetNeAttributeFilterIdOk returns a tuple with the NeAttributeFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeFilterId

`func (o *AttributeProperties) SetNeAttributeFilterId(v string)`

SetNeAttributeFilterId sets NeAttributeFilterId field to given value.

### HasNeAttributeFilterId

`func (o *AttributeProperties) HasNeAttributeFilterId() bool`

HasNeAttributeFilterId returns a boolean if a field has been set.

### GetReverseAssociationAttributeId

`func (o *AttributeProperties) GetReverseAssociationAttributeId() string`

GetReverseAssociationAttributeId returns the ReverseAssociationAttributeId field if non-nil, zero value otherwise.

### GetReverseAssociationAttributeIdOk

`func (o *AttributeProperties) GetReverseAssociationAttributeIdOk() (*string, bool)`

GetReverseAssociationAttributeIdOk returns a tuple with the ReverseAssociationAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseAssociationAttributeId

`func (o *AttributeProperties) SetReverseAssociationAttributeId(v string)`

SetReverseAssociationAttributeId sets ReverseAssociationAttributeId field to given value.

### HasReverseAssociationAttributeId

`func (o *AttributeProperties) HasReverseAssociationAttributeId() bool`

HasReverseAssociationAttributeId returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *AttributeProperties) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *AttributeProperties) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *AttributeProperties) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *AttributeProperties) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.

### GetLegacyId

`func (o *AttributeProperties) GetLegacyId() string`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *AttributeProperties) GetLegacyIdOk() (*string, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *AttributeProperties) SetLegacyId(v string)`

SetLegacyId sets LegacyId field to given value.

### HasLegacyId

`func (o *AttributeProperties) HasLegacyId() bool`

HasLegacyId returns a boolean if a field has been set.

### GetTmpCreatedAt

`func (o *AttributeProperties) GetTmpCreatedAt() SailPointTime`

GetTmpCreatedAt returns the TmpCreatedAt field if non-nil, zero value otherwise.

### GetTmpCreatedAtOk

`func (o *AttributeProperties) GetTmpCreatedAtOk() (*SailPointTime, bool)`

GetTmpCreatedAtOk returns a tuple with the TmpCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpCreatedAt

`func (o *AttributeProperties) SetTmpCreatedAt(v SailPointTime)`

SetTmpCreatedAt sets TmpCreatedAt field to given value.

### HasTmpCreatedAt

`func (o *AttributeProperties) HasTmpCreatedAt() bool`

HasTmpCreatedAt returns a boolean if a field has been set.

### GetTmpUpdatedAt

`func (o *AttributeProperties) GetTmpUpdatedAt() SailPointTime`

GetTmpUpdatedAt returns the TmpUpdatedAt field if non-nil, zero value otherwise.

### GetTmpUpdatedAtOk

`func (o *AttributeProperties) GetTmpUpdatedAtOk() (*SailPointTime, bool)`

GetTmpUpdatedAtOk returns a tuple with the TmpUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpUpdatedAt

`func (o *AttributeProperties) SetTmpUpdatedAt(v SailPointTime)`

SetTmpUpdatedAt sets TmpUpdatedAt field to given value.

### HasTmpUpdatedAt

`func (o *AttributeProperties) HasTmpUpdatedAt() bool`

HasTmpUpdatedAt returns a boolean if a field has been set.


