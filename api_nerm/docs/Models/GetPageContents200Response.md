---
id: nerm-get-page-contents200-response
title: GetPageContents200Response
pagination_label: GetPageContents200Response
sidebar_label: GetPageContents200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetPageContents200Response', 'NERMGetPageContents200Response'] 
slug: /tools/sdk/go/nerm/models/get-page-contents200-response
tags: ['SDK', 'Software Development Kit', 'GetPageContents200Response', 'NERMGetPageContents200Response']
---

# GetPageContents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PageContent**](page-content) |  | [optional] 

## Methods

### NewGetPageContents200Response

`func NewGetPageContents200Response() *GetPageContents200Response`

NewGetPageContents200Response instantiates a new GetPageContents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPageContents200ResponseWithDefaults

`func NewGetPageContents200ResponseWithDefaults() *GetPageContents200Response`

NewGetPageContents200ResponseWithDefaults instantiates a new GetPageContents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetPageContents200Response) GetPage() PageContent`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPageContents200Response) GetPageOk() (*PageContent, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPageContents200Response) SetPage(v PageContent)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPageContents200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.


