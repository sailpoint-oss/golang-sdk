---
id: nerm-profile-page
title: ProfilePage
pagination_label: ProfilePage
sidebar_label: ProfilePage
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfilePage', 'NERMProfilePage'] 
slug: /tools/sdk/go/nerm/models/profile-page
tags: ['SDK', 'Software Development Kit', 'ProfilePage', 'NERMProfilePage']
---

# ProfilePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The user-specified identifier of the page | [optional] 
**Description** | Pointer to **string** | The description of the page | [optional] 
**Name** | Pointer to **string** | The name of the page | [optional] 
**Archived** | Pointer to **bool** | Determines whether the page is archived | [optional] 

## Methods

### NewProfilePage

`func NewProfilePage() *ProfilePage`

NewProfilePage instantiates a new ProfilePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilePageWithDefaults

`func NewProfilePageWithDefaults() *ProfilePage`

NewProfilePageWithDefaults instantiates a new ProfilePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ProfilePage) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ProfilePage) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ProfilePage) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ProfilePage) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *ProfilePage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProfilePage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProfilePage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProfilePage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ProfilePage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfilePage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfilePage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProfilePage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *ProfilePage) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ProfilePage) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ProfilePage) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ProfilePage) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


