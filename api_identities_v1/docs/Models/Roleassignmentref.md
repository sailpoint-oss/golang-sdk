---
id: v1-roleassignmentref
title: Roleassignmentref
pagination_label: Roleassignmentref
sidebar_label: Roleassignmentref
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Roleassignmentref', 'V1Roleassignmentref'] 
slug: /tools/sdk/go/identitiesv1/models/roleassignmentref
tags: ['SDK', 'Software Development Kit', 'Roleassignmentref', 'V1Roleassignmentref']
---

# Roleassignmentref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Assignment Id | [optional] 
**Role** | Pointer to [**Basereferencedto**](basereferencedto) |  | [optional] 
**AddedDate** | Pointer to **time.Time** | Date that the assignment was added | [optional] 
**RemoveDate** | Pointer to **NullableTime** | Date that the assignment will be removed | [optional] 

## Methods

### NewRoleassignmentref

`func NewRoleassignmentref() *Roleassignmentref`

NewRoleassignmentref instantiates a new Roleassignmentref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleassignmentrefWithDefaults

`func NewRoleassignmentrefWithDefaults() *Roleassignmentref`

NewRoleassignmentrefWithDefaults instantiates a new Roleassignmentref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Roleassignmentref) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Roleassignmentref) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Roleassignmentref) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Roleassignmentref) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *Roleassignmentref) GetRole() Basereferencedto`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Roleassignmentref) GetRoleOk() (*Basereferencedto, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Roleassignmentref) SetRole(v Basereferencedto)`

SetRole sets Role field to given value.

### HasRole

`func (o *Roleassignmentref) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAddedDate

`func (o *Roleassignmentref) GetAddedDate() time.Time`

GetAddedDate returns the AddedDate field if non-nil, zero value otherwise.

### GetAddedDateOk

`func (o *Roleassignmentref) GetAddedDateOk() (*time.Time, bool)`

GetAddedDateOk returns a tuple with the AddedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedDate

`func (o *Roleassignmentref) SetAddedDate(v time.Time)`

SetAddedDate sets AddedDate field to given value.

### HasAddedDate

`func (o *Roleassignmentref) HasAddedDate() bool`

HasAddedDate returns a boolean if a field has been set.

### GetRemoveDate

`func (o *Roleassignmentref) GetRemoveDate() time.Time`

GetRemoveDate returns the RemoveDate field if non-nil, zero value otherwise.

### GetRemoveDateOk

`func (o *Roleassignmentref) GetRemoveDateOk() (*time.Time, bool)`

GetRemoveDateOk returns a tuple with the RemoveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDate

`func (o *Roleassignmentref) SetRemoveDate(v time.Time)`

SetRemoveDate sets RemoveDate field to given value.

### HasRemoveDate

`func (o *Roleassignmentref) HasRemoveDate() bool`

HasRemoveDate returns a boolean if a field has been set.

### SetRemoveDateNil

`func (o *Roleassignmentref) SetRemoveDateNil(b bool)`

 SetRemoveDateNil sets the value for RemoveDate to be an explicit nil

### UnsetRemoveDate
`func (o *Roleassignmentref) UnsetRemoveDate()`

UnsetRemoveDate ensures that no value is present for RemoveDate, not even an explicit nil

