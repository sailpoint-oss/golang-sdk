---
id: nerm-delete-attribute-option-by-id200-response
title: DeleteAttributeOptionById200Response
pagination_label: DeleteAttributeOptionById200Response
sidebar_label: DeleteAttributeOptionById200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DeleteAttributeOptionById200Response', 'NERMDeleteAttributeOptionById200Response'] 
slug: /tools/sdk/go/nerm/models/delete-attribute-option-by-id200-response
tags: ['SDK', 'Software Development Kit', 'DeleteAttributeOptionById200Response', 'NERMDeleteAttributeOptionById200Response']
---

# DeleteAttributeOptionById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to **string** |  | [optional] [default to "The option {Option Value} has been deleted"]

## Methods

### NewDeleteAttributeOptionById200Response

`func NewDeleteAttributeOptionById200Response() *DeleteAttributeOptionById200Response`

NewDeleteAttributeOptionById200Response instantiates a new DeleteAttributeOptionById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAttributeOptionById200ResponseWithDefaults

`func NewDeleteAttributeOptionById200ResponseWithDefaults() *DeleteAttributeOptionById200Response`

NewDeleteAttributeOptionById200ResponseWithDefaults instantiates a new DeleteAttributeOptionById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *DeleteAttributeOptionById200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DeleteAttributeOptionById200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DeleteAttributeOptionById200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DeleteAttributeOptionById200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


