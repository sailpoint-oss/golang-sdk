---
id: v1-intelmachinesource
title: Intelmachinesource
pagination_label: Intelmachinesource
sidebar_label: Intelmachinesource
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelmachinesource', 'V1Intelmachinesource'] 
slug: /tools/sdk/go/intelligencepackage/models/intelmachinesource
tags: ['SDK', 'Software Development Kit', 'Intelmachinesource', 'V1Intelmachinesource']
---

# Intelmachinesource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the correlated source in Identity Security Cloud. | [optional] 
**Name** | Pointer to **string** | Display name of the source as configured in Identity Security Cloud. | [optional] 
**Type** | Pointer to **string** | Connector or source type classification for the backing system. | [optional] 

## Methods

### NewIntelmachinesource

`func NewIntelmachinesource() *Intelmachinesource`

NewIntelmachinesource instantiates a new Intelmachinesource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelmachinesourceWithDefaults

`func NewIntelmachinesourceWithDefaults() *Intelmachinesource`

NewIntelmachinesourceWithDefaults instantiates a new Intelmachinesource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Intelmachinesource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Intelmachinesource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Intelmachinesource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Intelmachinesource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Intelmachinesource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Intelmachinesource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Intelmachinesource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Intelmachinesource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Intelmachinesource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Intelmachinesource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Intelmachinesource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Intelmachinesource) HasType() bool`

HasType returns a boolean if a field has been set.


