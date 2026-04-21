---
id: nerm-profile-type-attribute
title: ProfileTypeAttribute
pagination_label: ProfileTypeAttribute
sidebar_label: ProfileTypeAttribute
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileTypeAttribute', 'NERMProfileTypeAttribute'] 
slug: /tools/sdk/go/nerm/models/profile-type-attribute
tags: ['SDK', 'Software Development Kit', 'ProfileTypeAttribute', 'NERMProfileTypeAttribute']
---

# ProfileTypeAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of ne attribute | [optional] [readonly] 
**Uid** | Pointer to **string** | Ne attribute's uid | [optional] [readonly] 
**Label** | **string** | Ne attribute's label | [readonly] 
**Synced** | Pointer to **string** | synced_attribute ID if there is one | [optional] 

## Methods

### NewProfileTypeAttribute

`func NewProfileTypeAttribute(label string, ) *ProfileTypeAttribute`

NewProfileTypeAttribute instantiates a new ProfileTypeAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileTypeAttributeWithDefaults

`func NewProfileTypeAttributeWithDefaults() *ProfileTypeAttribute`

NewProfileTypeAttributeWithDefaults instantiates a new ProfileTypeAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileTypeAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileTypeAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileTypeAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileTypeAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *ProfileTypeAttribute) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfileTypeAttribute) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfileTypeAttribute) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfileTypeAttribute) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLabel

`func (o *ProfileTypeAttribute) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProfileTypeAttribute) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProfileTypeAttribute) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSynced

`func (o *ProfileTypeAttribute) GetSynced() string`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *ProfileTypeAttribute) GetSyncedOk() (*string, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *ProfileTypeAttribute) SetSynced(v string)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *ProfileTypeAttribute) HasSynced() bool`

HasSynced returns a boolean if a field has been set.


