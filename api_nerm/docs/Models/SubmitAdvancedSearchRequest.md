---
id: nerm-submit-advanced-search-request
title: SubmitAdvancedSearchRequest
pagination_label: SubmitAdvancedSearchRequest
sidebar_label: SubmitAdvancedSearchRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitAdvancedSearchRequest', 'NERMSubmitAdvancedSearchRequest'] 
slug: /tools/sdk/go/nerm/models/submit-advanced-search-request
tags: ['SDK', 'Software Development Kit', 'SubmitAdvancedSearchRequest', 'NERMSubmitAdvancedSearchRequest']
---

# SubmitAdvancedSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSearch** | Pointer to [**AdvancedSearch1**](advanced-search1) |  | [optional] 

## Methods

### NewSubmitAdvancedSearchRequest

`func NewSubmitAdvancedSearchRequest() *SubmitAdvancedSearchRequest`

NewSubmitAdvancedSearchRequest instantiates a new SubmitAdvancedSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitAdvancedSearchRequestWithDefaults

`func NewSubmitAdvancedSearchRequestWithDefaults() *SubmitAdvancedSearchRequest`

NewSubmitAdvancedSearchRequestWithDefaults instantiates a new SubmitAdvancedSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSearch

`func (o *SubmitAdvancedSearchRequest) GetAdvancedSearch() AdvancedSearch1`

GetAdvancedSearch returns the AdvancedSearch field if non-nil, zero value otherwise.

### GetAdvancedSearchOk

`func (o *SubmitAdvancedSearchRequest) GetAdvancedSearchOk() (*AdvancedSearch1, bool)`

GetAdvancedSearchOk returns a tuple with the AdvancedSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSearch

`func (o *SubmitAdvancedSearchRequest) SetAdvancedSearch(v AdvancedSearch1)`

SetAdvancedSearch sets AdvancedSearch field to given value.

### HasAdvancedSearch

`func (o *SubmitAdvancedSearchRequest) HasAdvancedSearch() bool`

HasAdvancedSearch returns a boolean if a field has been set.


