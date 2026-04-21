---
id: nerm-search-request-audit-events
title: SearchRequestAuditEvents
pagination_label: SearchRequestAuditEvents
sidebar_label: SearchRequestAuditEvents
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SearchRequestAuditEvents', 'NERMSearchRequestAuditEvents'] 
slug: /tools/sdk/go/nerm/models/search-request-audit-events
tags: ['SDK', 'Software Development Kit', 'SearchRequestAuditEvents', 'NERMSearchRequestAuditEvents']
---

# SearchRequestAuditEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | How many records to skip before pulling records to return. | [optional] 
**SortBy** | Pointer to **string** | A column that we are sorting these records from. | [optional] 
**Limit** | Pointer to **int32** | The limiting count for the amount of records returned. | [optional] 
**Order** | Pointer to **string** | Which direction the list should be sorted by | [optional] 
**Filters** | Pointer to [**AuditEvent**](audit-event) |  | [optional] 

## Methods

### NewSearchRequestAuditEvents

`func NewSearchRequestAuditEvents() *SearchRequestAuditEvents`

NewSearchRequestAuditEvents instantiates a new SearchRequestAuditEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestAuditEventsWithDefaults

`func NewSearchRequestAuditEventsWithDefaults() *SearchRequestAuditEvents`

NewSearchRequestAuditEventsWithDefaults instantiates a new SearchRequestAuditEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *SearchRequestAuditEvents) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchRequestAuditEvents) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchRequestAuditEvents) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchRequestAuditEvents) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortBy

`func (o *SearchRequestAuditEvents) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *SearchRequestAuditEvents) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *SearchRequestAuditEvents) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *SearchRequestAuditEvents) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetLimit

`func (o *SearchRequestAuditEvents) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchRequestAuditEvents) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchRequestAuditEvents) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchRequestAuditEvents) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrder

`func (o *SearchRequestAuditEvents) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SearchRequestAuditEvents) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SearchRequestAuditEvents) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SearchRequestAuditEvents) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetFilters

`func (o *SearchRequestAuditEvents) GetFilters() AuditEvent`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SearchRequestAuditEvents) GetFiltersOk() (*AuditEvent, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SearchRequestAuditEvents) SetFilters(v AuditEvent)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SearchRequestAuditEvents) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


