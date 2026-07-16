---
id: v1-resourcev2
title: Resourcev2
pagination_label: Resourcev2
sidebar_label: Resourcev2
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Resourcev2', 'V1Resourcev2'] 
slug: /tools/sdk/go/machineidentities/models/resourcev2
tags: ['SDK', 'Software Development Kit', 'Resourcev2', 'V1Resourcev2']
---

# Resourcev2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The source resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of the source resource. | [optional] 
**Name** | Pointer to **string** | The display name of the source resource. | [optional] 
**Features** | Pointer to **[]string** | The set of features supported by the source resource. | [optional] 

## Methods

### NewResourcev2

`func NewResourcev2() *Resourcev2`

NewResourcev2 instantiates a new Resourcev2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcev2WithDefaults

`func NewResourcev2WithDefaults() *Resourcev2`

NewResourcev2WithDefaults instantiates a new Resourcev2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Resourcev2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resourcev2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resourcev2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Resourcev2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Resourcev2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Resourcev2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Resourcev2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Resourcev2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Resourcev2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resourcev2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resourcev2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Resourcev2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFeatures

`func (o *Resourcev2) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Resourcev2) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Resourcev2) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Resourcev2) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


