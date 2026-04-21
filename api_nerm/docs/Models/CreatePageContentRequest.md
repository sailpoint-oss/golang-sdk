---
id: nerm-create-page-content-request
title: CreatePageContentRequest
pagination_label: CreatePageContentRequest
sidebar_label: CreatePageContentRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreatePageContentRequest', 'NERMCreatePageContentRequest'] 
slug: /tools/sdk/go/nerm/models/create-page-content-request
tags: ['SDK', 'Software Development Kit', 'CreatePageContentRequest', 'NERMCreatePageContentRequest']
---

# CreatePageContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageContent** | Pointer to [**PageContent1**](page-content1) |  | [optional] 

## Methods

### NewCreatePageContentRequest

`func NewCreatePageContentRequest() *CreatePageContentRequest`

NewCreatePageContentRequest instantiates a new CreatePageContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePageContentRequestWithDefaults

`func NewCreatePageContentRequestWithDefaults() *CreatePageContentRequest`

NewCreatePageContentRequestWithDefaults instantiates a new CreatePageContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageContent

`func (o *CreatePageContentRequest) GetPageContent() PageContent1`

GetPageContent returns the PageContent field if non-nil, zero value otherwise.

### GetPageContentOk

`func (o *CreatePageContentRequest) GetPageContentOk() (*PageContent1, bool)`

GetPageContentOk returns a tuple with the PageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageContent

`func (o *CreatePageContentRequest) SetPageContent(v PageContent1)`

SetPageContent sets PageContent field to given value.

### HasPageContent

`func (o *CreatePageContentRequest) HasPageContent() bool`

HasPageContent returns a boolean if a field has been set.


