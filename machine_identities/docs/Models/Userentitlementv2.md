---
id: v1-userentitlementv2
title: Userentitlementv2
pagination_label: Userentitlementv2
sidebar_label: Userentitlementv2
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Userentitlementv2', 'V1Userentitlementv2'] 
slug: /tools/sdk/go/machineidentities/models/userentitlementv2
tags: ['SDK', 'Software Development Kit', 'Userentitlementv2', 'V1Userentitlementv2']
---

# Userentitlementv2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | Pointer to **string** | The source ID of the entitlement. | [optional] 
**EntitlementId** | Pointer to **string** | The ID of the entitlement. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the entitlement. | [optional] 
**Source** | Pointer to [**Userentitlementv2Source**](userentitlementv2-source) |  | [optional] 

## Methods

### NewUserentitlementv2

`func NewUserentitlementv2() *Userentitlementv2`

NewUserentitlementv2 instantiates a new Userentitlementv2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserentitlementv2WithDefaults

`func NewUserentitlementv2WithDefaults() *Userentitlementv2`

NewUserentitlementv2WithDefaults instantiates a new Userentitlementv2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *Userentitlementv2) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Userentitlementv2) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Userentitlementv2) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Userentitlementv2) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetEntitlementId

`func (o *Userentitlementv2) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *Userentitlementv2) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *Userentitlementv2) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.

### HasEntitlementId

`func (o *Userentitlementv2) HasEntitlementId() bool`

HasEntitlementId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Userentitlementv2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Userentitlementv2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Userentitlementv2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Userentitlementv2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *Userentitlementv2) GetSource() Userentitlementv2Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Userentitlementv2) GetSourceOk() (*Userentitlementv2Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Userentitlementv2) SetSource(v Userentitlementv2Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *Userentitlementv2) HasSource() bool`

HasSource returns a boolean if a field has been set.


