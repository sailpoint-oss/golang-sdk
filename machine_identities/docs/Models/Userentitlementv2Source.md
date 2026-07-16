---
id: v1-userentitlementv2-source
title: Userentitlementv2Source
pagination_label: Userentitlementv2Source
sidebar_label: Userentitlementv2Source
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Userentitlementv2Source', 'V1Userentitlementv2Source'] 
slug: /tools/sdk/go/machineidentities/models/userentitlementv2-source
tags: ['SDK', 'Software Development Kit', 'Userentitlementv2Source', 'V1Userentitlementv2Source']
---

# Userentitlementv2Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **Dtotype** |  | [optional] 
**Id** | Pointer to **string** | ID of the object to which this reference applies | [optional] 
**Name** | Pointer to **string** | Human-readable display name of the object to which this reference applies | [optional] 

## Methods

### NewUserentitlementv2Source

`func NewUserentitlementv2Source() *Userentitlementv2Source`

NewUserentitlementv2Source instantiates a new Userentitlementv2Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserentitlementv2SourceWithDefaults

`func NewUserentitlementv2SourceWithDefaults() *Userentitlementv2Source`

NewUserentitlementv2SourceWithDefaults instantiates a new Userentitlementv2Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Userentitlementv2Source) GetType() Dtotype`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Userentitlementv2Source) GetTypeOk() (*Dtotype, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Userentitlementv2Source) SetType(v Dtotype)`

SetType sets Type field to given value.

### HasType

`func (o *Userentitlementv2Source) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Userentitlementv2Source) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Userentitlementv2Source) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Userentitlementv2Source) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Userentitlementv2Source) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Userentitlementv2Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Userentitlementv2Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Userentitlementv2Source) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Userentitlementv2Source) HasName() bool`

HasName returns a boolean if a field has been set.


