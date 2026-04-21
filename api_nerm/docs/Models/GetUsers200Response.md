---
id: nerm-get-users200-response
title: GetUsers200Response
pagination_label: GetUsers200Response
sidebar_label: GetUsers200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetUsers200Response', 'NERMGetUsers200Response'] 
slug: /tools/sdk/go/nerm/models/get-users200-response
tags: ['SDK', 'Software Development Kit', 'GetUsers200Response', 'NERMGetUsers200Response']
---

# GetUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]User**](user) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetUsers200Response

`func NewGetUsers200Response() *GetUsers200Response`

NewGetUsers200Response instantiates a new GetUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsers200ResponseWithDefaults

`func NewGetUsers200ResponseWithDefaults() *GetUsers200Response`

NewGetUsers200ResponseWithDefaults instantiates a new GetUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GetUsers200Response) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetUsers200Response) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetUsers200Response) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetMetadata

`func (o *GetUsers200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetUsers200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetUsers200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetUsers200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


