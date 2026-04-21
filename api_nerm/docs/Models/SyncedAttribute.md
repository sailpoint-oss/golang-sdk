---
id: nerm-synced-attribute
title: SyncedAttribute
pagination_label: SyncedAttribute
sidebar_label: SyncedAttribute
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SyncedAttribute', 'NERMSyncedAttribute'] 
slug: /tools/sdk/go/nerm/models/synced-attribute
tags: ['SDK', 'Software Development Kit', 'SyncedAttribute', 'NERMSyncedAttribute']
---

# SyncedAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | synced attribute's ID | [optional] [readonly] 
**ProfileTypeId** | Pointer to **string** | Profile type ID of synced attribute | [optional] [readonly] 
**NeAttributeId** | Pointer to **string** | Ne attribute ID of synced attribute | [optional] [readonly] 

## Methods

### NewSyncedAttribute

`func NewSyncedAttribute() *SyncedAttribute`

NewSyncedAttribute instantiates a new SyncedAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncedAttributeWithDefaults

`func NewSyncedAttributeWithDefaults() *SyncedAttribute`

NewSyncedAttributeWithDefaults instantiates a new SyncedAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SyncedAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncedAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncedAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyncedAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfileTypeId

`func (o *SyncedAttribute) GetProfileTypeId() string`

GetProfileTypeId returns the ProfileTypeId field if non-nil, zero value otherwise.

### GetProfileTypeIdOk

`func (o *SyncedAttribute) GetProfileTypeIdOk() (*string, bool)`

GetProfileTypeIdOk returns a tuple with the ProfileTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileTypeId

`func (o *SyncedAttribute) SetProfileTypeId(v string)`

SetProfileTypeId sets ProfileTypeId field to given value.

### HasProfileTypeId

`func (o *SyncedAttribute) HasProfileTypeId() bool`

HasProfileTypeId returns a boolean if a field has been set.

### GetNeAttributeId

`func (o *SyncedAttribute) GetNeAttributeId() string`

GetNeAttributeId returns the NeAttributeId field if non-nil, zero value otherwise.

### GetNeAttributeIdOk

`func (o *SyncedAttribute) GetNeAttributeIdOk() (*string, bool)`

GetNeAttributeIdOk returns a tuple with the NeAttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeId

`func (o *SyncedAttribute) SetNeAttributeId(v string)`

SetNeAttributeId sets NeAttributeId field to given value.

### HasNeAttributeId

`func (o *SyncedAttribute) HasNeAttributeId() bool`

HasNeAttributeId returns a boolean if a field has been set.


