---
id: nerm-attribute1
title: Attribute1
pagination_label: Attribute1
sidebar_label: Attribute1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Attribute1', 'NERMAttribute1'] 
slug: /tools/sdk/go/nerm/models/attribute1
tags: ['SDK', 'Software Development Kit', 'Attribute1', 'NERMAttribute1']
---

# Attribute1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the attribute | [optional] [readonly] 
**Uid** | Pointer to **string** | The user-specified identifier of the attribute | [optional] [readonly] 
**Label** | Pointer to **string** | The label for the attribute | [optional] 
**Description** | Pointer to **string** | A description of the attribute | [optional] 
**ToolTip** | Pointer to **string** | The helper text that accompanies the attribute | [optional] 
**Archived** | Pointer to **bool** | Whether the attribute is archived | [optional] 
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
**ValidationsAttributes** | Pointer to [**Attribute1ValidationsAttributes**](attribute1-validations-attributes) |  | [optional] 

## Methods

### NewAttribute1

`func NewAttribute1() *Attribute1`

NewAttribute1 instantiates a new Attribute1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttribute1WithDefaults

`func NewAttribute1WithDefaults() *Attribute1`

NewAttribute1WithDefaults instantiates a new Attribute1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attribute1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attribute1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attribute1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Attribute1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *Attribute1) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Attribute1) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Attribute1) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Attribute1) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLabel

`func (o *Attribute1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Attribute1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Attribute1) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Attribute1) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *Attribute1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Attribute1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Attribute1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Attribute1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetToolTip

`func (o *Attribute1) GetToolTip() string`

GetToolTip returns the ToolTip field if non-nil, zero value otherwise.

### GetToolTipOk

`func (o *Attribute1) GetToolTipOk() (*string, bool)`

GetToolTipOk returns a tuple with the ToolTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolTip

`func (o *Attribute1) SetToolTip(v string)`

SetToolTip sets ToolTip field to given value.

### HasToolTip

`func (o *Attribute1) HasToolTip() bool`

HasToolTip returns a boolean if a field has been set.

### GetArchived

`func (o *Attribute1) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Attribute1) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Attribute1) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Attribute1) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDateFormat

`func (o *Attribute1) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *Attribute1) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *Attribute1) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *Attribute1) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetSelectableStatus

`func (o *Attribute1) GetSelectableStatus() string`

GetSelectableStatus returns the SelectableStatus field if non-nil, zero value otherwise.

### GetSelectableStatusOk

`func (o *Attribute1) GetSelectableStatusOk() (*string, bool)`

GetSelectableStatusOk returns a tuple with the SelectableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectableStatus

`func (o *Attribute1) SetSelectableStatus(v string)`

SetSelectableStatus sets SelectableStatus field to given value.

### HasSelectableStatus

`func (o *Attribute1) HasSelectableStatus() bool`

HasSelectableStatus returns a boolean if a field has been set.

### GetRiskType

`func (o *Attribute1) GetRiskType() string`

GetRiskType returns the RiskType field if non-nil, zero value otherwise.

### GetRiskTypeOk

`func (o *Attribute1) GetRiskTypeOk() (*string, bool)`

GetRiskTypeOk returns a tuple with the RiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskType

`func (o *Attribute1) SetRiskType(v string)`

SetRiskType sets RiskType field to given value.

### HasRiskType

`func (o *Attribute1) HasRiskType() bool`

HasRiskType returns a boolean if a field has been set.

### GetOwnershipDriven

`func (o *Attribute1) GetOwnershipDriven() bool`

GetOwnershipDriven returns the OwnershipDriven field if non-nil, zero value otherwise.

### GetOwnershipDrivenOk

`func (o *Attribute1) GetOwnershipDrivenOk() (*bool, bool)`

GetOwnershipDrivenOk returns a tuple with the OwnershipDriven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipDriven

`func (o *Attribute1) SetOwnershipDriven(v bool)`

SetOwnershipDriven sets OwnershipDriven field to given value.

### HasOwnershipDriven

`func (o *Attribute1) HasOwnershipDriven() bool`

HasOwnershipDriven returns a boolean if a field has been set.

### GetAllowMultipleSelections

`func (o *Attribute1) GetAllowMultipleSelections() bool`

GetAllowMultipleSelections returns the AllowMultipleSelections field if non-nil, zero value otherwise.

### GetAllowMultipleSelectionsOk

`func (o *Attribute1) GetAllowMultipleSelectionsOk() (*bool, bool)`

GetAllowMultipleSelectionsOk returns a tuple with the AllowMultipleSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSelections

`func (o *Attribute1) SetAllowMultipleSelections(v bool)`

SetAllowMultipleSelections sets AllowMultipleSelections field to given value.

### HasAllowMultipleSelections

`func (o *Attribute1) HasAllowMultipleSelections() bool`

HasAllowMultipleSelections returns a boolean if a field has been set.

### GetFilteredByNeAttribute

`func (o *Attribute1) GetFilteredByNeAttribute() bool`

GetFilteredByNeAttribute returns the FilteredByNeAttribute field if non-nil, zero value otherwise.

### GetFilteredByNeAttributeOk

`func (o *Attribute1) GetFilteredByNeAttributeOk() (*bool, bool)`

GetFilteredByNeAttributeOk returns a tuple with the FilteredByNeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredByNeAttribute

`func (o *Attribute1) SetFilteredByNeAttribute(v bool)`

SetFilteredByNeAttribute sets FilteredByNeAttribute field to given value.

### HasFilteredByNeAttribute

`func (o *Attribute1) HasFilteredByNeAttribute() bool`

HasFilteredByNeAttribute returns a boolean if a field has been set.

### GetFilteringNeAttributeId

`func (o *Attribute1) GetFilteringNeAttributeId() string`

GetFilteringNeAttributeId returns the FilteringNeAttributeId field if non-nil, zero value otherwise.

### GetFilteringNeAttributeIdOk

`func (o *Attribute1) GetFilteringNeAttributeIdOk() (*string, bool)`

GetFilteringNeAttributeIdOk returns a tuple with the FilteringNeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringNeAttributeId

`func (o *Attribute1) SetFilteringNeAttributeId(v string)`

SetFilteringNeAttributeId sets FilteringNeAttributeId field to given value.

### HasFilteringNeAttributeId

`func (o *Attribute1) HasFilteringNeAttributeId() bool`

HasFilteringNeAttributeId returns a boolean if a field has been set.

### GetNeAttributeFilterId

`func (o *Attribute1) GetNeAttributeFilterId() string`

GetNeAttributeFilterId returns the NeAttributeFilterId field if non-nil, zero value otherwise.

### GetNeAttributeFilterIdOk

`func (o *Attribute1) GetNeAttributeFilterIdOk() (*string, bool)`

GetNeAttributeFilterIdOk returns a tuple with the NeAttributeFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeFilterId

`func (o *Attribute1) SetNeAttributeFilterId(v string)`

SetNeAttributeFilterId sets NeAttributeFilterId field to given value.

### HasNeAttributeFilterId

`func (o *Attribute1) HasNeAttributeFilterId() bool`

HasNeAttributeFilterId returns a boolean if a field has been set.

### GetReverseAssociationAttribute

`func (o *Attribute1) GetReverseAssociationAttribute() AttributeProperties`

GetReverseAssociationAttribute returns the ReverseAssociationAttribute field if non-nil, zero value otherwise.

### GetReverseAssociationAttributeOk

`func (o *Attribute1) GetReverseAssociationAttributeOk() (*AttributeProperties, bool)`

GetReverseAssociationAttributeOk returns a tuple with the ReverseAssociationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseAssociationAttribute

`func (o *Attribute1) SetReverseAssociationAttribute(v AttributeProperties)`

SetReverseAssociationAttribute sets ReverseAssociationAttribute field to given value.

### HasReverseAssociationAttribute

`func (o *Attribute1) HasReverseAssociationAttribute() bool`

HasReverseAssociationAttribute returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *Attribute1) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *Attribute1) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *Attribute1) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *Attribute1) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.

### GetDataType

`func (o *Attribute1) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *Attribute1) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *Attribute1) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *Attribute1) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetType

`func (o *Attribute1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Attribute1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Attribute1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Attribute1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidationsAttributes

`func (o *Attribute1) GetValidationsAttributes() Attribute1ValidationsAttributes`

GetValidationsAttributes returns the ValidationsAttributes field if non-nil, zero value otherwise.

### GetValidationsAttributesOk

`func (o *Attribute1) GetValidationsAttributesOk() (*Attribute1ValidationsAttributes, bool)`

GetValidationsAttributesOk returns a tuple with the ValidationsAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationsAttributes

`func (o *Attribute1) SetValidationsAttributes(v Attribute1ValidationsAttributes)`

SetValidationsAttributes sets ValidationsAttributes field to given value.

### HasValidationsAttributes

`func (o *Attribute1) HasValidationsAttributes() bool`

HasValidationsAttributes returns a boolean if a field has been set.


