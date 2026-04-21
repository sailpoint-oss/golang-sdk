---
id: nerm-form
title: Form
pagination_label: Form
sidebar_label: Form
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Form', 'NERMForm'] 
slug: /tools/sdk/go/nerm/models/form
tags: ['SDK', 'Software Development Kit', 'Form', 'NERMForm']
---

# Form

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The user-specified identifier of the form | [optional] 
**Description** | Pointer to **string** | The description of the form | [optional] 
**Name** | Pointer to **string** | The name of the form | [optional] 
**Archived** | Pointer to **bool** | Determines whether the form is archived | [optional] 
**Id** | Pointer to **string** | The id of the form | [optional] [readonly] 

## Methods

### NewForm

`func NewForm() *Form`

NewForm instantiates a new Form object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormWithDefaults

`func NewFormWithDefaults() *Form`

NewFormWithDefaults instantiates a new Form object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *Form) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Form) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Form) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Form) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDescription

`func (o *Form) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Form) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Form) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Form) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Form) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Form) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Form) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Form) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArchived

`func (o *Form) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Form) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Form) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Form) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetId

`func (o *Form) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Form) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Form) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Form) HasId() bool`

HasId returns a boolean if a field has been set.


