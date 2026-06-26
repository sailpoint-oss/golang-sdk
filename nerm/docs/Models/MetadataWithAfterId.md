---
id: nerm-metadata-with-after-id
title: MetadataWithAfterId
pagination_label: MetadataWithAfterId
sidebar_label: MetadataWithAfterId
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MetadataWithAfterId', 'NERMMetadataWithAfterId'] 
slug: /tools/sdk/go/nerm/models/metadata-with-after-id
tags: ['SDK', 'Software Development Kit', 'MetadataWithAfterId', 'NERMMetadataWithAfterId']
---

# MetadataWithAfterId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The maximum number of records to return in the search | [optional] 
**Offset** | Pointer to **int32** | The number of records to skip before starting to return results. | [optional] 
**Total** | Pointer to **int32** | The total number of records available matching the search criteria. | [optional] 
**Next** | Pointer to **string** | The ID of the first record in the next set of results | [optional] 
**Previous** | Pointer to **string** | The ID of the last record in the previous set of results | [optional] 
**AfterId** | Pointer to **string** | The ID from which the search will start, ignoring all records before it. | [optional] 

## Methods

### NewMetadataWithAfterId

`func NewMetadataWithAfterId() *MetadataWithAfterId`

NewMetadataWithAfterId instantiates a new MetadataWithAfterId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithAfterIdWithDefaults

`func NewMetadataWithAfterIdWithDefaults() *MetadataWithAfterId`

NewMetadataWithAfterIdWithDefaults instantiates a new MetadataWithAfterId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *MetadataWithAfterId) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetadataWithAfterId) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetadataWithAfterId) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MetadataWithAfterId) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *MetadataWithAfterId) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MetadataWithAfterId) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MetadataWithAfterId) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *MetadataWithAfterId) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *MetadataWithAfterId) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MetadataWithAfterId) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MetadataWithAfterId) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MetadataWithAfterId) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetNext

`func (o *MetadataWithAfterId) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *MetadataWithAfterId) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *MetadataWithAfterId) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *MetadataWithAfterId) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *MetadataWithAfterId) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *MetadataWithAfterId) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *MetadataWithAfterId) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *MetadataWithAfterId) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetAfterId

`func (o *MetadataWithAfterId) GetAfterId() string`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *MetadataWithAfterId) GetAfterIdOk() (*string, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *MetadataWithAfterId) SetAfterId(v string)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *MetadataWithAfterId) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.


