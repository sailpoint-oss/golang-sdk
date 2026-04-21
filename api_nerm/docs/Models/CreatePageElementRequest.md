---
id: nerm-create-page-element-request
title: CreatePageElementRequest
pagination_label: CreatePageElementRequest
sidebar_label: CreatePageElementRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreatePageElementRequest', 'NERMCreatePageElementRequest'] 
slug: /tools/sdk/go/nerm/models/create-page-element-request
tags: ['SDK', 'Software Development Kit', 'CreatePageElementRequest', 'NERMCreatePageElementRequest']
---

# CreatePageElementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageElement** | Pointer to [**PageElement1**](page-element1) |  | [optional] 

## Methods

### NewCreatePageElementRequest

`func NewCreatePageElementRequest() *CreatePageElementRequest`

NewCreatePageElementRequest instantiates a new CreatePageElementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePageElementRequestWithDefaults

`func NewCreatePageElementRequestWithDefaults() *CreatePageElementRequest`

NewCreatePageElementRequestWithDefaults instantiates a new CreatePageElementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageElement

`func (o *CreatePageElementRequest) GetPageElement() PageElement1`

GetPageElement returns the PageElement field if non-nil, zero value otherwise.

### GetPageElementOk

`func (o *CreatePageElementRequest) GetPageElementOk() (*PageElement1, bool)`

GetPageElementOk returns a tuple with the PageElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageElement

`func (o *CreatePageElementRequest) SetPageElement(v PageElement1)`

SetPageElement sets PageElement field to given value.

### HasPageElement

`func (o *CreatePageElementRequest) HasPageElement() bool`

HasPageElement returns a boolean if a field has been set.


