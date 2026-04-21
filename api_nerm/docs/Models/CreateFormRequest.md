---
id: nerm-create-form-request
title: CreateFormRequest
pagination_label: CreateFormRequest
sidebar_label: CreateFormRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateFormRequest', 'NERMCreateFormRequest'] 
slug: /tools/sdk/go/nerm/models/create-form-request
tags: ['SDK', 'Software Development Kit', 'CreateFormRequest', 'NERMCreateFormRequest']
---

# CreateFormRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forms** | Pointer to [**[]Forms**](forms) |  | [optional] 

## Methods

### NewCreateFormRequest

`func NewCreateFormRequest() *CreateFormRequest`

NewCreateFormRequest instantiates a new CreateFormRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFormRequestWithDefaults

`func NewCreateFormRequestWithDefaults() *CreateFormRequest`

NewCreateFormRequestWithDefaults instantiates a new CreateFormRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForms

`func (o *CreateFormRequest) GetForms() []Forms`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *CreateFormRequest) GetFormsOk() (*[]Forms, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *CreateFormRequest) SetForms(v []Forms)`

SetForms sets Forms field to given value.

### HasForms

`func (o *CreateFormRequest) HasForms() bool`

HasForms returns a boolean if a field has been set.


