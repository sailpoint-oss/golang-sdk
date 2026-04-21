---
id: nerm-get-page-elements200-response
title: GetPageElements200Response
pagination_label: GetPageElements200Response
sidebar_label: GetPageElements200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetPageElements200Response', 'NERMGetPageElements200Response'] 
slug: /tools/sdk/go/nerm/models/get-page-elements200-response
tags: ['SDK', 'Software Development Kit', 'GetPageElements200Response', 'NERMGetPageElements200Response']
---

# GetPageElements200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageElement** | Pointer to [**PageElement**](page-element) |  | [optional] 

## Methods

### NewGetPageElements200Response

`func NewGetPageElements200Response() *GetPageElements200Response`

NewGetPageElements200Response instantiates a new GetPageElements200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPageElements200ResponseWithDefaults

`func NewGetPageElements200ResponseWithDefaults() *GetPageElements200Response`

NewGetPageElements200ResponseWithDefaults instantiates a new GetPageElements200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageElement

`func (o *GetPageElements200Response) GetPageElement() PageElement`

GetPageElement returns the PageElement field if non-nil, zero value otherwise.

### GetPageElementOk

`func (o *GetPageElements200Response) GetPageElementOk() (*PageElement, bool)`

GetPageElementOk returns a tuple with the PageElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageElement

`func (o *GetPageElements200Response) SetPageElement(v PageElement)`

SetPageElement sets PageElement field to given value.

### HasPageElement

`func (o *GetPageElements200Response) HasPageElement() bool`

HasPageElement returns a boolean if a field has been set.


