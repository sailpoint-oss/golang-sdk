---
id: nerm-get-profile-type-attributes-sort-parameter
title: GetProfileTypeAttributesSortParameter
pagination_label: GetProfileTypeAttributesSortParameter
sidebar_label: GetProfileTypeAttributesSortParameter
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetProfileTypeAttributesSortParameter', 'NERMGetProfileTypeAttributesSortParameter'] 
slug: /tools/sdk/go/nerm/models/get-profile-type-attributes-sort-parameter
tags: ['SDK', 'Software Development Kit', 'GetProfileTypeAttributesSortParameter', 'NERMGetProfileTypeAttributesSortParameter']
---

# GetProfileTypeAttributesSortParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 

## Methods

### NewGetProfileTypeAttributesSortParameter

`func NewGetProfileTypeAttributesSortParameter() *GetProfileTypeAttributesSortParameter`

NewGetProfileTypeAttributesSortParameter instantiates a new GetProfileTypeAttributesSortParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfileTypeAttributesSortParameterWithDefaults

`func NewGetProfileTypeAttributesSortParameterWithDefaults() *GetProfileTypeAttributesSortParameter`

NewGetProfileTypeAttributesSortParameterWithDefaults instantiates a new GetProfileTypeAttributesSortParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *GetProfileTypeAttributesSortParameter) GetAttr() string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *GetProfileTypeAttributesSortParameter) GetAttrOk() (*string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *GetProfileTypeAttributesSortParameter) SetAttr(v string)`

SetAttr sets Attr field to given value.

### HasAttr

`func (o *GetProfileTypeAttributesSortParameter) HasAttr() bool`

HasAttr returns a boolean if a field has been set.

### GetOrder

`func (o *GetProfileTypeAttributesSortParameter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetProfileTypeAttributesSortParameter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetProfileTypeAttributesSortParameter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetProfileTypeAttributesSortParameter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


