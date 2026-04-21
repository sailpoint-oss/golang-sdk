---
id: nerm-form1
title: Form1
pagination_label: Form1
sidebar_label: Form1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Form1', 'NERMForm1'] 
slug: /tools/sdk/go/nerm/models/form1
tags: ['SDK', 'Software Development Kit', 'Form1', 'NERMForm1']
---

# Form1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The user-specified identifier of the form | [optional] 
**Description** | Pointer to **string** | The description of the form | [optional] 
**Name** | Pointer to **string** | The name of the form | [optional] 
**Archived** | Pointer to **bool** | Determines whether the form is archived | [optional] 

## Methods

### NewForm1

`func NewForm1() *Form1`

NewForm1 instantiates a new Form1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForm1WithDefaults

`func NewForm1WithDefaults() *Form1`

NewForm1WithDefaults instantiates a new Form1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Form1) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Form1) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Form1) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Form1) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *Form1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Form1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Form1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Form1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Form1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Form1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Form1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Form1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *Form1) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Form1) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Form1) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Form1) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


