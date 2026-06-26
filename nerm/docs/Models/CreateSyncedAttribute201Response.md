---
id: nerm-create-synced-attribute201-response
title: CreateSyncedAttribute201Response
pagination_label: CreateSyncedAttribute201Response
sidebar_label: CreateSyncedAttribute201Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateSyncedAttribute201Response', 'NERMCreateSyncedAttribute201Response'] 
slug: /tools/sdk/go/nerm/models/create-synced-attribute201-response
tags: ['SDK', 'Software Development Kit', 'CreateSyncedAttribute201Response', 'NERMCreateSyncedAttribute201Response']
---

# CreateSyncedAttribute201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncedAttribute** | Pointer to [**SyncedAttribute**](synced-attribute) |  | [optional] 

## Methods

### NewCreateSyncedAttribute201Response

`func NewCreateSyncedAttribute201Response() *CreateSyncedAttribute201Response`

NewCreateSyncedAttribute201Response instantiates a new CreateSyncedAttribute201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSyncedAttribute201ResponseWithDefaults

`func NewCreateSyncedAttribute201ResponseWithDefaults() *CreateSyncedAttribute201Response`

NewCreateSyncedAttribute201ResponseWithDefaults instantiates a new CreateSyncedAttribute201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncedAttribute

`func (o *CreateSyncedAttribute201Response) GetSyncedAttribute() SyncedAttribute`

GetSyncedAttribute returns the SyncedAttribute field if non-nil, zero value otherwise.

### GetSyncedAttributeOk

`func (o *CreateSyncedAttribute201Response) GetSyncedAttributeOk() (*SyncedAttribute, bool)`

GetSyncedAttributeOk returns a tuple with the SyncedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAttribute

`func (o *CreateSyncedAttribute201Response) SetSyncedAttribute(v SyncedAttribute)`

SetSyncedAttribute sets SyncedAttribute field to given value.

### HasSyncedAttribute

`func (o *CreateSyncedAttribute201Response) HasSyncedAttribute() bool`

HasSyncedAttribute returns a boolean if a field has been set.


