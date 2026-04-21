---
id: nerm-create-page-content-translation-request
title: CreatePageContentTranslationRequest
pagination_label: CreatePageContentTranslationRequest
sidebar_label: CreatePageContentTranslationRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreatePageContentTranslationRequest', 'NERMCreatePageContentTranslationRequest'] 
slug: /tools/sdk/go/nerm/models/create-page-content-translation-request
tags: ['SDK', 'Software Development Kit', 'CreatePageContentTranslationRequest', 'NERMCreatePageContentTranslationRequest']
---

# CreatePageContentTranslationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageContentTranslation** | Pointer to [**PageContentTranslation1**](page-content-translation1) |  | [optional] 

## Methods

### NewCreatePageContentTranslationRequest

`func NewCreatePageContentTranslationRequest() *CreatePageContentTranslationRequest`

NewCreatePageContentTranslationRequest instantiates a new CreatePageContentTranslationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePageContentTranslationRequestWithDefaults

`func NewCreatePageContentTranslationRequestWithDefaults() *CreatePageContentTranslationRequest`

NewCreatePageContentTranslationRequestWithDefaults instantiates a new CreatePageContentTranslationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageContentTranslation

`func (o *CreatePageContentTranslationRequest) GetPageContentTranslation() PageContentTranslation1`

GetPageContentTranslation returns the PageContentTranslation field if non-nil, zero value otherwise.

### GetPageContentTranslationOk

`func (o *CreatePageContentTranslationRequest) GetPageContentTranslationOk() (*PageContentTranslation1, bool)`

GetPageContentTranslationOk returns a tuple with the PageContentTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageContentTranslation

`func (o *CreatePageContentTranslationRequest) SetPageContentTranslation(v PageContentTranslation1)`

SetPageContentTranslation sets PageContentTranslation field to given value.

### HasPageContentTranslation

`func (o *CreatePageContentTranslationRequest) HasPageContentTranslation() bool`

HasPageContentTranslation returns a boolean if a field has been set.


