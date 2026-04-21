---
id: nerm-search-advanced-search200-response
title: SearchAdvancedSearch200Response
pagination_label: SearchAdvancedSearch200Response
sidebar_label: SearchAdvancedSearch200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SearchAdvancedSearch200Response', 'NERMSearchAdvancedSearch200Response'] 
slug: /tools/sdk/go/nerm/models/search-advanced-search200-response
tags: ['SDK', 'Software Development Kit', 'SearchAdvancedSearch200Response', 'NERMSearchAdvancedSearch200Response']
---

# SearchAdvancedSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]Profile**](profile) |  | [optional] 

## Methods

### NewSearchAdvancedSearch200Response

`func NewSearchAdvancedSearch200Response() *SearchAdvancedSearch200Response`

NewSearchAdvancedSearch200Response instantiates a new SearchAdvancedSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAdvancedSearch200ResponseWithDefaults

`func NewSearchAdvancedSearch200ResponseWithDefaults() *SearchAdvancedSearch200Response`

NewSearchAdvancedSearch200ResponseWithDefaults instantiates a new SearchAdvancedSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *SearchAdvancedSearch200Response) GetProfiles() []Profile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *SearchAdvancedSearch200Response) GetProfilesOk() (*[]Profile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *SearchAdvancedSearch200Response) SetProfiles(v []Profile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *SearchAdvancedSearch200Response) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


