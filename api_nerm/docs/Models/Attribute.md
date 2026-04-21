---
id: nerm-attribute
title: Attribute
pagination_label: Attribute
sidebar_label: Attribute
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Attribute', 'NERMAttribute'] 
slug: /tools/sdk/go/nerm/models/attribute
tags: ['SDK', 'Software Development Kit', 'Attribute', 'NERMAttribute']
---

# Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the attribute | [optional] [readonly] 
**Uid** | Pointer to **string** | The user-specified identifier of the attribute | [optional] [readonly] 
**Label** | Pointer to **string** | The label for the attribute | [optional] 
**Description** | Pointer to **string** | A description of the attribute | [optional] 
**ToolTip** | Pointer to **string** | The helper text that accompanies the attribute | [optional] 
**Crypt** | Pointer to **bool** | Whether the attribute is encrypted | [optional] 
**Archived** | Pointer to **bool** | Whether the attribute is archived | [optional] 
**ArchivedOn** | Pointer to **SailPointTime** | When the attribute was archived | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | When the attribute was created | [optional] 
**UpdatedAt** | Pointer to **SailPointTime** | When the attribute was updated | [optional] 
**DateFormat** | Pointer to **string** | The format of the date input if it is a date input | [optional] 
**SelectableStatus** | Pointer to **string** | The status of the profiles that can be selected | [optional] 
**RiskType** | Pointer to **string** | Type of risk that applies to the attribute | [optional] 
**OwnershipDriven** | Pointer to **bool** | Only shows profiles that the user currently has access to, to be selected | [optional] 
**AllowMultipleSelections** | Pointer to **bool** | Whether or not multiple selections can be made on something like a contributor search. | [optional] 
**FilteredByNeAttribute** | Pointer to **bool** | Whether or not the attribute is filtered by another attribute | [optional] 
**FilteringNeAttributeId** | Pointer to **string** | The ID of the filtering attribute | [optional] 
**NeAttributeFilterId** | Pointer to **string** | The ID of the attribute filter | [optional] 
**ReverseAssociationAttribute** | Pointer to [**AttributeProperties**](attribute-properties) |  | [optional] 
**ProfileTypeId** | Pointer to **string** | The ID of the profile type the attribute applies to | [optional] 
**DataType** | Pointer to **string** | The type of data that applies to the attribute | [optional] 
**Type** | Pointer to **string** | The attribute's type | [optional] 

## Methods

### NewAttribute

`func NewAttribute() *Attribute`

NewAttribute instantiates a new Attribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeWithDefaults

`func NewAttributeWithDefaults() *Attribute`

NewAttributeWithDefaults instantiates a new Attribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Attribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *Attribute) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Attribute) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Attribute) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Attribute) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLabel

`func (o *Attribute) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Attribute) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Attribute) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Attribute) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *Attribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Attribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Attribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Attribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetToolTip

`func (o *Attribute) GetToolTip() string`

GetToolTip returns the ToolTip field if non-nil, zero value otherwise.

### GetToolTipOk

`func (o *Attribute) GetToolTipOk() (*string, bool)`

GetToolTipOk returns a tuple with the ToolTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolTip

`func (o *Attribute) SetToolTip(v string)`

SetToolTip sets ToolTip field to given value.

### HasToolTip

`func (o *Attribute) HasToolTip() bool`

HasToolTip returns a boolean if a field has been set.

### GetCrypt

`func (o *Attribute) GetCrypt() bool`

GetCrypt returns the Crypt field if non-nil, zero value otherwise.

### GetCryptOk

`func (o *Attribute) GetCryptOk() (*bool, bool)`

GetCryptOk returns a tuple with the Crypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrypt

`func (o *Attribute) SetCrypt(v bool)`

SetCrypt sets Crypt field to given value.

### HasCrypt

`func (o *Attribute) HasCrypt() bool`

HasCrypt returns a boolean if a field has been set.

### GetArchived

`func (o *Attribute) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Attribute) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Attribute) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Attribute) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArchivedOn

`func (o *Attribute) GetArchivedOn() SailPointTime`

GetArchivedOn returns the ArchivedOn field if non-nil, zero value otherwise.

### GetArchivedOnOk

`func (o *Attribute) GetArchivedOnOk() (*SailPointTime, bool)`

GetArchivedOnOk returns a tuple with the ArchivedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedOn

`func (o *Attribute) SetArchivedOn(v SailPointTime)`

SetArchivedOn sets ArchivedOn field to given value.

### HasArchivedOn

`func (o *Attribute) HasArchivedOn() bool`

HasArchivedOn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Attribute) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Attribute) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Attribute) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Attribute) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Attribute) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Attribute) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Attribute) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Attribute) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDateFormat

`func (o *Attribute) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *Attribute) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *Attribute) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *Attribute) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetSelectableStatus

`func (o *Attribute) GetSelectableStatus() string`

GetSelectableStatus returns the SelectableStatus field if non-nil, zero value otherwise.

### GetSelectableStatusOk

`func (o *Attribute) GetSelectableStatusOk() (*string, bool)`

GetSelectableStatusOk returns a tuple with the SelectableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectableStatus

`func (o *Attribute) SetSelectableStatus(v string)`

SetSelectableStatus sets SelectableStatus field to given value.

### HasSelectableStatus

`func (o *Attribute) HasSelectableStatus() bool`

HasSelectableStatus returns a boolean if a field has been set.

### GetRiskType

`func (o *Attribute) GetRiskType() string`

GetRiskType returns the RiskType field if non-nil, zero value otherwise.

### GetRiskTypeOk

`func (o *Attribute) GetRiskTypeOk() (*string, bool)`

GetRiskTypeOk returns a tuple with the RiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskType

`func (o *Attribute) SetRiskType(v string)`

SetRiskType sets RiskType field to given value.

### HasRiskType

`func (o *Attribute) HasRiskType() bool`

HasRiskType returns a boolean if a field has been set.

### GetOwnershipDriven

`func (o *Attribute) GetOwnershipDriven() bool`

GetOwnershipDriven returns the OwnershipDriven field if non-nil, zero value otherwise.

### GetOwnershipDrivenOk

`func (o *Attribute) GetOwnershipDrivenOk() (*bool, bool)`

GetOwnershipDrivenOk returns a tuple with the OwnershipDriven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipDriven

`func (o *Attribute) SetOwnershipDriven(v bool)`

SetOwnershipDriven sets OwnershipDriven field to given value.

### HasOwnershipDriven

`func (o *Attribute) HasOwnershipDriven() bool`

HasOwnershipDriven returns a boolean if a field has been set.

### GetAllowMultipleSelections

`func (o *Attribute) GetAllowMultipleSelections() bool`

GetAllowMultipleSelections returns the AllowMultipleSelections field if non-nil, zero value otherwise.

### GetAllowMultipleSelectionsOk

`func (o *Attribute) GetAllowMultipleSelectionsOk() (*bool, bool)`

GetAllowMultipleSelectionsOk returns a tuple with the AllowMultipleSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSelections

`func (o *Attribute) SetAllowMultipleSelections(v bool)`

SetAllowMultipleSelections sets AllowMultipleSelections field to given value.

### HasAllowMultipleSelections

`func (o *Attribute) HasAllowMultipleSelections() bool`

HasAllowMultipleSelections returns a boolean if a field has been set.

### GetFilteredByNeAttribute

`func (o *Attribute) GetFilteredByNeAttribute() bool`

GetFilteredByNeAttribute returns the FilteredByNeAttribute field if non-nil, zero value otherwise.

### GetFilteredByNeAttributeOk

`func (o *Attribute) GetFilteredByNeAttributeOk() (*bool, bool)`

GetFilteredByNeAttributeOk returns a tuple with the FilteredByNeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredByNeAttribute

`func (o *Attribute) SetFilteredByNeAttribute(v bool)`

SetFilteredByNeAttribute sets FilteredByNeAttribute field to given value.

### HasFilteredByNeAttribute

`func (o *Attribute) HasFilteredByNeAttribute() bool`

HasFilteredByNeAttribute returns a boolean if a field has been set.

### GetFilteringNeAttributeId

`func (o *Attribute) GetFilteringNeAttributeId() string`

GetFilteringNeAttributeId returns the FilteringNeAttributeId field if non-nil, zero value otherwise.

### GetFilteringNeAttributeIdOk

`func (o *Attribute) GetFilteringNeAttributeIdOk() (*string, bool)`

GetFilteringNeAttributeIdOk returns a tuple with the FilteringNeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringNeAttributeId

`func (o *Attribute) SetFilteringNeAttributeId(v string)`

SetFilteringNeAttributeId sets FilteringNeAttributeId field to given value.

### HasFilteringNeAttributeId

`func (o *Attribute) HasFilteringNeAttributeId() bool`

HasFilteringNeAttributeId returns a boolean if a field has been set.

### GetNeAttributeFilterId

`func (o *Attribute) GetNeAttributeFilterId() string`

GetNeAttributeFilterId returns the NeAttributeFilterId field if non-nil, zero value otherwise.

### GetNeAttributeFilterIdOk

`func (o *Attribute) GetNeAttributeFilterIdOk() (*string, bool)`

GetNeAttributeFilterIdOk returns a tuple with the NeAttributeFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeFilterId

`func (o *Attribute) SetNeAttributeFilterId(v string)`

SetNeAttributeFilterId sets NeAttributeFilterId field to given value.

### HasNeAttributeFilterId

`func (o *Attribute) HasNeAttributeFilterId() bool`

HasNeAttributeFilterId returns a boolean if a field has been set.

### GetReverseAssociationAttribute

`func (o *Attribute) GetReverseAssociationAttribute() AttributeProperties`

GetReverseAssociationAttribute returns the ReverseAssociationAttribute field if non-nil, zero value otherwise.

### GetReverseAssociationAttributeOk

`func (o *Attribute) GetReverseAssociationAttributeOk() (*AttributeProperties, bool)`

GetReverseAssociationAttributeOk returns a tuple with the ReverseAssociationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseAssociationAttribute

`func (o *Attribute) SetReverseAssociationAttribute(v AttributeProperties)`

SetReverseAssociationAttribute sets ReverseAssociationAttribute field to given value.

### HasReverseAssociationAttribute

`func (o *Attribute) HasReverseAssociationAttribute() bool`

HasReverseAssociationAttribute returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *Attribute) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *Attribute) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *Attribute) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *Attribute) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.

### GetDataType

`func (o *Attribute) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *Attribute) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *Attribute) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *Attribute) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetType

`func (o *Attribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Attribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Attribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Attribute) HasType() bool`

HasType returns a boolean if a field has been set.


