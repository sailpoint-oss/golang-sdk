---
id: nerm-forms
title: Forms
pagination_label: Forms
sidebar_label: Forms
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Forms', 'NERMForms'] 
slug: /tools/sdk/go/nerm/models/forms
tags: ['SDK', 'Software Development Kit', 'Forms', 'NERMForms']
---

# Forms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The user-specified identifier of the form | [optional] 
**Description** | Pointer to **string** | The description of the form | [optional] 
**Name** | Pointer to **string** | The name of the form | [optional] 
**Archived** | Pointer to **bool** | Determines whether the form is archived | [optional] 

## Methods

### NewForms

`func NewForms() *Forms`

NewForms instantiates a new Forms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormsWithDefaults

`func NewFormsWithDefaults() *Forms`

NewFormsWithDefaults instantiates a new Forms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Forms) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Forms) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Forms) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Forms) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *Forms) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Forms) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Forms) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Forms) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Forms) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Forms) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Forms) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Forms) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *Forms) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Forms) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Forms) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Forms) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


