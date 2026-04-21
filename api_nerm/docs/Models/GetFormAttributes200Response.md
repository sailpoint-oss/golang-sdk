---
id: nerm-get-form-attributes200-response
title: GetFormAttributes200Response
pagination_label: GetFormAttributes200Response
sidebar_label: GetFormAttributes200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetFormAttributes200Response', 'NERMGetFormAttributes200Response'] 
slug: /tools/sdk/go/nerm/models/get-form-attributes200-response
tags: ['SDK', 'Software Development Kit', 'GetFormAttributes200Response', 'NERMGetFormAttributes200Response']
---

# GetFormAttributes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormAttribute** | Pointer to [**FormAttribute**](form-attribute) |  | [optional] 

## Methods

### NewGetFormAttributes200Response

`func NewGetFormAttributes200Response() *GetFormAttributes200Response`

NewGetFormAttributes200Response instantiates a new GetFormAttributes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFormAttributes200ResponseWithDefaults

`func NewGetFormAttributes200ResponseWithDefaults() *GetFormAttributes200Response`

NewGetFormAttributes200ResponseWithDefaults instantiates a new GetFormAttributes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormAttribute

`func (o *GetFormAttributes200Response) GetFormAttribute() FormAttribute`

GetFormAttribute returns the FormAttribute field if non-nil, zero value otherwise.

### GetFormAttributeOk

`func (o *GetFormAttributes200Response) GetFormAttributeOk() (*FormAttribute, bool)`

GetFormAttributeOk returns a tuple with the FormAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormAttribute

`func (o *GetFormAttributes200Response) SetFormAttribute(v FormAttribute)`

SetFormAttribute sets FormAttribute field to given value.

### HasFormAttribute

`func (o *GetFormAttributes200Response) HasFormAttribute() bool`

HasFormAttribute returns a boolean if a field has been set.


