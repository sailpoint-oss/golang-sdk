---
id: nerm-language
title: Language
pagination_label: Language
sidebar_label: Language
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Language', 'NERMLanguage'] 
slug: /tools/sdk/go/nerm/models/language
tags: ['SDK', 'Software Development Kit', 'Language', 'NERMLanguage']
---

# Language

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | Set default to True to set a default language, to set another language to default, set default to True on the target language and the current default will become disabled | [optional] 
**Enabled** | Pointer to **bool** | True when the language is enabled, False when it is not | [optional] 
**Locale** | Pointer to **string** | The locale string for the language, current options are- en, fr, es, de, fr-CA | [optional] 

## Methods

### NewLanguage

`func NewLanguage() *Language`

NewLanguage instantiates a new Language object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageWithDefaults

`func NewLanguageWithDefaults() *Language`

NewLanguageWithDefaults instantiates a new Language object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *Language) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Language) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Language) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Language) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnabled

`func (o *Language) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Language) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Language) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Language) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocale

`func (o *Language) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Language) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Language) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Language) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


