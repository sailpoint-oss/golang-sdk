---
id: nerm-page-content-translation
title: PageContentTranslation
pagination_label: PageContentTranslation
sidebar_label: PageContentTranslation
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageContentTranslation', 'NERMPageContentTranslation'] 
slug: /tools/sdk/go/nerm/models/page-content-translation
tags: ['SDK', 'Software Development Kit', 'PageContentTranslation', 'NERMPageContentTranslation']
---

# PageContentTranslation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the page content | [optional] [readonly] 
**Uid** | Pointer to **string** | The user-specified identifier for the record | [optional] 
**Locale** | Pointer to **string** | The language locale this translation contains. | [optional] 
**Value** | Pointer to **string** | The translated string to present in the user interface. | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The date-time the record created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **SailPointTime** | The date-time the record was last updated. | [optional] [readonly] 

## Methods

### NewPageContentTranslation

`func NewPageContentTranslation() *PageContentTranslation`

NewPageContentTranslation instantiates a new PageContentTranslation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageContentTranslationWithDefaults

`func NewPageContentTranslationWithDefaults() *PageContentTranslation`

NewPageContentTranslationWithDefaults instantiates a new PageContentTranslation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageContentTranslation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageContentTranslation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageContentTranslation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageContentTranslation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *PageContentTranslation) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PageContentTranslation) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PageContentTranslation) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *PageContentTranslation) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLocale

`func (o *PageContentTranslation) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PageContentTranslation) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PageContentTranslation) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PageContentTranslation) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetValue

`func (o *PageContentTranslation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PageContentTranslation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PageContentTranslation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PageContentTranslation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PageContentTranslation) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageContentTranslation) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageContentTranslation) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageContentTranslation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PageContentTranslation) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PageContentTranslation) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PageContentTranslation) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PageContentTranslation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


