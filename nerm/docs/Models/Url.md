---
id: nerm-url
title: Url
pagination_label: Url
sidebar_label: Url
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Url', 'NERMUrl'] 
slug: /tools/sdk/go/nerm/models/url
tags: ['SDK', 'Software Development Kit', 'Url', 'NERMUrl']
---

# Url

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUrl

`func NewUrl() *Url`

NewUrl instantiates a new Url object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlWithDefaults

`func NewUrlWithDefaults() *Url`

NewUrlWithDefaults instantiates a new Url object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Url) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Url) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Url) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Url) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


