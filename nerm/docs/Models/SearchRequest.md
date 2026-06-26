---
id: nerm-search-request
title: SearchRequest
pagination_label: SearchRequest
sidebar_label: SearchRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SearchRequest', 'NERMSearchRequest'] 
slug: /tools/sdk/go/nerm/models/search-request
tags: ['SDK', 'Software Development Kit', 'SearchRequest', 'NERMSearchRequest']
---

# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEvents** | Pointer to [**SearchRequestAuditEvents**](search-request-audit-events) |  | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest() *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEvents

`func (o *SearchRequest) GetAuditEvents() SearchRequestAuditEvents`

GetAuditEvents returns the AuditEvents field if non-nil, zero value otherwise.

### GetAuditEventsOk

`func (o *SearchRequest) GetAuditEventsOk() (*SearchRequestAuditEvents, bool)`

GetAuditEventsOk returns a tuple with the AuditEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEvents

`func (o *SearchRequest) SetAuditEvents(v SearchRequestAuditEvents)`

SetAuditEvents sets AuditEvents field to given value.

### HasAuditEvents

`func (o *SearchRequest) HasAuditEvents() bool`

HasAuditEvents returns a boolean if a field has been set.


