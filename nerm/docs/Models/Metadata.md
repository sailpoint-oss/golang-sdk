---
id: nerm-metadata
title: Metadata
pagination_label: Metadata
sidebar_label: Metadata
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Metadata', 'NERMMetadata'] 
slug: /tools/sdk/go/nerm/models/metadata
tags: ['SDK', 'Software Development Kit', 'Metadata', 'NERMMetadata']
---

# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *Metadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Metadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Metadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Metadata) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *Metadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Metadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Metadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Metadata) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *Metadata) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Metadata) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Metadata) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Metadata) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetNext

`func (o *Metadata) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Metadata) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Metadata) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Metadata) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *Metadata) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Metadata) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Metadata) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *Metadata) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


