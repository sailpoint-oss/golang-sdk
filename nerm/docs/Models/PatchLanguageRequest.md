---
id: nerm-patch-language-request
title: PatchLanguageRequest
pagination_label: PatchLanguageRequest
sidebar_label: PatchLanguageRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PatchLanguageRequest', 'NERMPatchLanguageRequest'] 
slug: /tools/sdk/go/nerm/models/patch-language-request
tags: ['SDK', 'Software Development Kit', 'PatchLanguageRequest', 'NERMPatchLanguageRequest']
---

# PatchLanguageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**Language**](language) |  | [optional] 

## Methods

### NewPatchLanguageRequest

`func NewPatchLanguageRequest() *PatchLanguageRequest`

NewPatchLanguageRequest instantiates a new PatchLanguageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchLanguageRequestWithDefaults

`func NewPatchLanguageRequestWithDefaults() *PatchLanguageRequest`

NewPatchLanguageRequestWithDefaults instantiates a new PatchLanguageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *PatchLanguageRequest) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PatchLanguageRequest) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PatchLanguageRequest) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PatchLanguageRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


