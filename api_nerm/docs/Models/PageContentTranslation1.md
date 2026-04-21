---
id: nerm-page-content-translation1
title: PageContentTranslation1
pagination_label: PageContentTranslation1
sidebar_label: PageContentTranslation1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageContentTranslation1', 'NERMPageContentTranslation1'] 
slug: /tools/sdk/go/nerm/models/page-content-translation1
tags: ['SDK', 'Software Development Kit', 'PageContentTranslation1', 'NERMPageContentTranslation1']
---

# PageContentTranslation1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The user-specified identifier for the record | [optional] 
**PageContentUid** | Pointer to **string** | The user-specified identifier of the page content record this translation applies to; one of page_content_id or page_content_uid must be present. | [optional] 
**PageContentId** | Pointer to **string** | The id of the page content record this translation applies to; one of page_content_id or page_content_uid must be present. | [optional] 
**Locale** | Pointer to **string** | The language locale this translation contains. | [optional] 
**Value** | Pointer to **string** | The translated string to present in the user interface. | [optional] 

## Methods

### NewPageContentTranslation1

`func NewPageContentTranslation1() *PageContentTranslation1`

NewPageContentTranslation1 instantiates a new PageContentTranslation1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageContentTranslation1WithDefaults

`func NewPageContentTranslation1WithDefaults() *PageContentTranslation1`

NewPageContentTranslation1WithDefaults instantiates a new PageContentTranslation1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *PageContentTranslation1) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PageContentTranslation1) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PageContentTranslation1) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *PageContentTranslation1) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetPageContentUid

`func (o *PageContentTranslation1) GetPageContentUid() string`

GetPageContentUid returns the PageContentUid field if non-nil, zero value otherwise.

### GetPageContentUidOk

`func (o *PageContentTranslation1) GetPageContentUidOk() (*string, bool)`

GetPageContentUidOk returns a tuple with the PageContentUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageContentUid

`func (o *PageContentTranslation1) SetPageContentUid(v string)`

SetPageContentUid sets PageContentUid field to given value.

### HasPageContentUid

`func (o *PageContentTranslation1) HasPageContentUid() bool`

HasPageContentUid returns a boolean if a field has been set.

### GetPageContentId

`func (o *PageContentTranslation1) GetPageContentId() string`

GetPageContentId returns the PageContentId field if non-nil, zero value otherwise.

### GetPageContentIdOk

`func (o *PageContentTranslation1) GetPageContentIdOk() (*string, bool)`

GetPageContentIdOk returns a tuple with the PageContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageContentId

`func (o *PageContentTranslation1) SetPageContentId(v string)`

SetPageContentId sets PageContentId field to given value.

### HasPageContentId

`func (o *PageContentTranslation1) HasPageContentId() bool`

HasPageContentId returns a boolean if a field has been set.

### GetLocale

`func (o *PageContentTranslation1) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PageContentTranslation1) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PageContentTranslation1) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PageContentTranslation1) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetValue

`func (o *PageContentTranslation1) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PageContentTranslation1) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PageContentTranslation1) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PageContentTranslation1) HasValue() bool`

HasValue returns a boolean if a field has been set.


