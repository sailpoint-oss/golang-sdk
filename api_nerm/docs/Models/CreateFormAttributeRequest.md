---
id: nerm-create-form-attribute-request
title: CreateFormAttributeRequest
pagination_label: CreateFormAttributeRequest
sidebar_label: CreateFormAttributeRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateFormAttributeRequest', 'NERMCreateFormAttributeRequest'] 
slug: /tools/sdk/go/nerm/models/create-form-attribute-request
tags: ['SDK', 'Software Development Kit', 'CreateFormAttributeRequest', 'NERMCreateFormAttributeRequest']
---

# CreateFormAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormAttribute** | Pointer to [**FormAttributes**](form-attributes) |  | [optional] 

## Methods

### NewCreateFormAttributeRequest

`func NewCreateFormAttributeRequest() *CreateFormAttributeRequest`

NewCreateFormAttributeRequest instantiates a new CreateFormAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFormAttributeRequestWithDefaults

`func NewCreateFormAttributeRequestWithDefaults() *CreateFormAttributeRequest`

NewCreateFormAttributeRequestWithDefaults instantiates a new CreateFormAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormAttribute

`func (o *CreateFormAttributeRequest) GetFormAttribute() FormAttributes`

GetFormAttribute returns the FormAttribute field if non-nil, zero value otherwise.

### GetFormAttributeOk

`func (o *CreateFormAttributeRequest) GetFormAttributeOk() (*FormAttributes, bool)`

GetFormAttributeOk returns a tuple with the FormAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormAttribute

`func (o *CreateFormAttributeRequest) SetFormAttribute(v FormAttributes)`

SetFormAttribute sets FormAttribute field to given value.

### HasFormAttribute

`func (o *CreateFormAttributeRequest) HasFormAttribute() bool`

HasFormAttribute returns a boolean if a field has been set.


