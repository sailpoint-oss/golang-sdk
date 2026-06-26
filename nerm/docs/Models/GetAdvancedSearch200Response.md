---
id: nerm-get-advanced-search200-response
title: GetAdvancedSearch200Response
pagination_label: GetAdvancedSearch200Response
sidebar_label: GetAdvancedSearch200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetAdvancedSearch200Response', 'NERMGetAdvancedSearch200Response'] 
slug: /tools/sdk/go/nerm/models/get-advanced-search200-response
tags: ['SDK', 'Software Development Kit', 'GetAdvancedSearch200Response', 'NERMGetAdvancedSearch200Response']
---

# GetAdvancedSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSearch** | Pointer to [**[]AdvancedSearch**](advanced-search) |  | [optional] 

## Methods

### NewGetAdvancedSearch200Response

`func NewGetAdvancedSearch200Response() *GetAdvancedSearch200Response`

NewGetAdvancedSearch200Response instantiates a new GetAdvancedSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdvancedSearch200ResponseWithDefaults

`func NewGetAdvancedSearch200ResponseWithDefaults() *GetAdvancedSearch200Response`

NewGetAdvancedSearch200ResponseWithDefaults instantiates a new GetAdvancedSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSearch

`func (o *GetAdvancedSearch200Response) GetAdvancedSearch() []AdvancedSearch`

GetAdvancedSearch returns the AdvancedSearch field if non-nil, zero value otherwise.

### GetAdvancedSearchOk

`func (o *GetAdvancedSearch200Response) GetAdvancedSearchOk() (*[]AdvancedSearch, bool)`

GetAdvancedSearchOk returns a tuple with the AdvancedSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSearch

`func (o *GetAdvancedSearch200Response) SetAdvancedSearch(v []AdvancedSearch)`

SetAdvancedSearch sets AdvancedSearch field to given value.

### HasAdvancedSearch

`func (o *GetAdvancedSearch200Response) HasAdvancedSearch() bool`

HasAdvancedSearch returns a boolean if a field has been set.


