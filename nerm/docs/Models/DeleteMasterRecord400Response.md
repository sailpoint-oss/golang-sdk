---
id: nerm-delete-master-record400-response
title: DeleteMasterRecord400Response
pagination_label: DeleteMasterRecord400Response
sidebar_label: DeleteMasterRecord400Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DeleteMasterRecord400Response', 'NERMDeleteMasterRecord400Response'] 
slug: /tools/sdk/go/nerm/models/delete-master-record400-response
tags: ['SDK', 'Software Development Kit', 'DeleteMasterRecord400Response', 'NERMDeleteMasterRecord400Response']
---

# DeleteMasterRecord400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | A message describing the error that occurred | [optional] 

## Methods

### NewDeleteMasterRecord400Response

`func NewDeleteMasterRecord400Response() *DeleteMasterRecord400Response`

NewDeleteMasterRecord400Response instantiates a new DeleteMasterRecord400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMasterRecord400ResponseWithDefaults

`func NewDeleteMasterRecord400ResponseWithDefaults() *DeleteMasterRecord400Response`

NewDeleteMasterRecord400ResponseWithDefaults instantiates a new DeleteMasterRecord400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DeleteMasterRecord400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteMasterRecord400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteMasterRecord400Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DeleteMasterRecord400Response) HasError() bool`

HasError returns a boolean if a field has been set.


