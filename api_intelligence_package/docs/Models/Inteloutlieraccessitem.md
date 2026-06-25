---
id: v1-inteloutlieraccessitem
title: Inteloutlieraccessitem
pagination_label: Inteloutlieraccessitem
sidebar_label: Inteloutlieraccessitem
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Inteloutlieraccessitem', 'V1Inteloutlieraccessitem'] 
slug: /tools/sdk/go/intelligencepackage/models/inteloutlieraccessitem
tags: ['SDK', 'Software Development Kit', 'Inteloutlieraccessitem', 'V1Inteloutlieraccessitem']
---

# Inteloutlieraccessitem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable identifier of the outlier access-item row. | 
**DisplayName** | **string** | Display label of the risky access item. | 
**Description** | Pointer to **NullableString** | Optional descriptive text for the risky access item. | [optional] 
**AccessType** | **string** | Access item type (for example ENTITLEMENT, ROLE, ACCESS_PROFILE, ACCOUNT, or APP). | 
**SourceName** | **string** | Source name where the risky access item exists. | 
**ExtremelyRare** | **bool** | Indicates whether analytics marked this item as extremely rare. | 

## Methods

### NewInteloutlieraccessitem

`func NewInteloutlieraccessitem(id string, displayName string, accessType string, sourceName string, extremelyRare bool, ) *Inteloutlieraccessitem`

NewInteloutlieraccessitem instantiates a new Inteloutlieraccessitem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteloutlieraccessitemWithDefaults

`func NewInteloutlieraccessitemWithDefaults() *Inteloutlieraccessitem`

NewInteloutlieraccessitemWithDefaults instantiates a new Inteloutlieraccessitem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Inteloutlieraccessitem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Inteloutlieraccessitem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Inteloutlieraccessitem) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *Inteloutlieraccessitem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Inteloutlieraccessitem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Inteloutlieraccessitem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *Inteloutlieraccessitem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Inteloutlieraccessitem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Inteloutlieraccessitem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Inteloutlieraccessitem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Inteloutlieraccessitem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Inteloutlieraccessitem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAccessType

`func (o *Inteloutlieraccessitem) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *Inteloutlieraccessitem) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *Inteloutlieraccessitem) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.


### GetSourceName

`func (o *Inteloutlieraccessitem) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Inteloutlieraccessitem) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Inteloutlieraccessitem) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetExtremelyRare

`func (o *Inteloutlieraccessitem) GetExtremelyRare() bool`

GetExtremelyRare returns the ExtremelyRare field if non-nil, zero value otherwise.

### GetExtremelyRareOk

`func (o *Inteloutlieraccessitem) GetExtremelyRareOk() (*bool, bool)`

GetExtremelyRareOk returns a tuple with the ExtremelyRare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtremelyRare

`func (o *Inteloutlieraccessitem) SetExtremelyRare(v bool)`

SetExtremelyRare sets ExtremelyRare field to given value.



