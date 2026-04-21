---
id: nerm-get-forms200-response
title: GetForms200Response
pagination_label: GetForms200Response
sidebar_label: GetForms200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetForms200Response', 'NERMGetForms200Response'] 
slug: /tools/sdk/go/nerm/models/get-forms200-response
tags: ['SDK', 'Software Development Kit', 'GetForms200Response', 'NERMGetForms200Response']
---

# GetForms200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forms** | Pointer to [**[]Form**](form) |  | [optional] 

## Methods

### NewGetForms200Response

`func NewGetForms200Response() *GetForms200Response`

NewGetForms200Response instantiates a new GetForms200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetForms200ResponseWithDefaults

`func NewGetForms200ResponseWithDefaults() *GetForms200Response`

NewGetForms200ResponseWithDefaults instantiates a new GetForms200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForms

`func (o *GetForms200Response) GetForms() []Form`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *GetForms200Response) GetFormsOk() (*[]Form, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *GetForms200Response) SetForms(v []Form)`

SetForms sets Forms field to given value.

### HasForms

`func (o *GetForms200Response) HasForms() bool`

HasForms returns a boolean if a field has been set.


