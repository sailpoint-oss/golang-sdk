---
id: v1-intelprivilegedaccessitemwire
title: Intelprivilegedaccessitemwire
pagination_label: Intelprivilegedaccessitemwire
sidebar_label: Intelprivilegedaccessitemwire
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelprivilegedaccessitemwire', 'V1Intelprivilegedaccessitemwire'] 
slug: /tools/sdk/go/intelligencepackagev1/models/intelprivilegedaccessitemwire
tags: ['SDK', 'Software Development Kit', 'Intelprivilegedaccessitemwire', 'V1Intelprivilegedaccessitemwire']
---

# Intelprivilegedaccessitemwire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privileged** | **bool** | True when SDS Search classifies this item as privileged access for the identity. | 
**DisplayName** | Pointer to **string** | Display label for the privileged access item in administrative experiences. | [optional] 
**Name** | Pointer to **string** | Technical name of the privileged access item from SDS Search. | [optional] 
**Standalone** | Pointer to **bool** | True when the privileged item is modeled as a standalone entitlement or access object. | [optional] [default to false]
**Id** | **string** | Identifier of the privileged access item returned by SDS Search. | 
**Source** | Pointer to [**IntelprivilegedaccessitemwireSource**](intelprivilegedaccessitemwire-source) |  | [optional] 
**Attribute** | Pointer to **string** | Source attribute name that carries the privileged value when applicable. | [optional] 
**Type** | **string** | Object type classification from SDS Search (for example ENTITLEMENT). | 
**Value** | Pointer to **string** | Privileged value on the source attribute when applicable. | [optional] 

## Methods

### NewIntelprivilegedaccessitemwire

`func NewIntelprivilegedaccessitemwire(privileged bool, id string, type_ string, ) *Intelprivilegedaccessitemwire`

NewIntelprivilegedaccessitemwire instantiates a new Intelprivilegedaccessitemwire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelprivilegedaccessitemwireWithDefaults

`func NewIntelprivilegedaccessitemwireWithDefaults() *Intelprivilegedaccessitemwire`

NewIntelprivilegedaccessitemwireWithDefaults instantiates a new Intelprivilegedaccessitemwire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivileged

`func (o *Intelprivilegedaccessitemwire) GetPrivileged() bool`

GetPrivileged returns the Privileged field if non-nil, zero value otherwise.

### GetPrivilegedOk

`func (o *Intelprivilegedaccessitemwire) GetPrivilegedOk() (*bool, bool)`

GetPrivilegedOk returns a tuple with the Privileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileged

`func (o *Intelprivilegedaccessitemwire) SetPrivileged(v bool)`

SetPrivileged sets Privileged field to given value.


### GetDisplayName

`func (o *Intelprivilegedaccessitemwire) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Intelprivilegedaccessitemwire) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Intelprivilegedaccessitemwire) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Intelprivilegedaccessitemwire) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *Intelprivilegedaccessitemwire) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Intelprivilegedaccessitemwire) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Intelprivilegedaccessitemwire) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Intelprivilegedaccessitemwire) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStandalone

`func (o *Intelprivilegedaccessitemwire) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *Intelprivilegedaccessitemwire) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *Intelprivilegedaccessitemwire) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *Intelprivilegedaccessitemwire) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetId

`func (o *Intelprivilegedaccessitemwire) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Intelprivilegedaccessitemwire) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Intelprivilegedaccessitemwire) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *Intelprivilegedaccessitemwire) GetSource() IntelprivilegedaccessitemwireSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Intelprivilegedaccessitemwire) GetSourceOk() (*IntelprivilegedaccessitemwireSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Intelprivilegedaccessitemwire) SetSource(v IntelprivilegedaccessitemwireSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Intelprivilegedaccessitemwire) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAttribute

`func (o *Intelprivilegedaccessitemwire) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Intelprivilegedaccessitemwire) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Intelprivilegedaccessitemwire) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Intelprivilegedaccessitemwire) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetType

`func (o *Intelprivilegedaccessitemwire) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Intelprivilegedaccessitemwire) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Intelprivilegedaccessitemwire) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *Intelprivilegedaccessitemwire) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Intelprivilegedaccessitemwire) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Intelprivilegedaccessitemwire) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Intelprivilegedaccessitemwire) HasValue() bool`

HasValue returns a boolean if a field has been set.


