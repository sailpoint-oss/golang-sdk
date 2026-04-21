---
id: nerm-update-page-content-translation-by-id-request
title: UpdatePageContentTranslationByIdRequest
pagination_label: UpdatePageContentTranslationByIdRequest
sidebar_label: UpdatePageContentTranslationByIdRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UpdatePageContentTranslationByIdRequest', 'NERMUpdatePageContentTranslationByIdRequest'] 
slug: /tools/sdk/go/nerm/models/update-page-content-translation-by-id-request
tags: ['SDK', 'Software Development Kit', 'UpdatePageContentTranslationByIdRequest', 'NERMUpdatePageContentTranslationByIdRequest']
---

# UpdatePageContentTranslationByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageContentTranslation** | Pointer to [**PageContent1**](page-content1) |  | [optional] 

## Methods

### NewUpdatePageContentTranslationByIdRequest

`func NewUpdatePageContentTranslationByIdRequest() *UpdatePageContentTranslationByIdRequest`

NewUpdatePageContentTranslationByIdRequest instantiates a new UpdatePageContentTranslationByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePageContentTranslationByIdRequestWithDefaults

`func NewUpdatePageContentTranslationByIdRequestWithDefaults() *UpdatePageContentTranslationByIdRequest`

NewUpdatePageContentTranslationByIdRequestWithDefaults instantiates a new UpdatePageContentTranslationByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageContentTranslation

`func (o *UpdatePageContentTranslationByIdRequest) GetPageContentTranslation() PageContent1`

GetPageContentTranslation returns the PageContentTranslation field if non-nil, zero value otherwise.

### GetPageContentTranslationOk

`func (o *UpdatePageContentTranslationByIdRequest) GetPageContentTranslationOk() (*PageContent1, bool)`

GetPageContentTranslationOk returns a tuple with the PageContentTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageContentTranslation

`func (o *UpdatePageContentTranslationByIdRequest) SetPageContentTranslation(v PageContent1)`

SetPageContentTranslation sets PageContentTranslation field to given value.

### HasPageContentTranslation

`func (o *UpdatePageContentTranslationByIdRequest) HasPageContentTranslation() bool`

HasPageContentTranslation returns a boolean if a field has been set.


