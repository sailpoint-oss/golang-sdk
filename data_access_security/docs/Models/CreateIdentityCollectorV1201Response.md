---
id: v1-create-identity-collector-v1201-response
title: CreateIdentityCollectorV1201Response
pagination_label: CreateIdentityCollectorV1201Response
sidebar_label: CreateIdentityCollectorV1201Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateIdentityCollectorV1201Response', 'V1CreateIdentityCollectorV1201Response'] 
slug: /tools/sdk/go/dataaccesssecurity/models/create-identity-collector-v1201-response
tags: ['SDK', 'Software Development Kit', 'CreateIdentityCollectorV1201Response', 'V1CreateIdentityCollectorV1201Response']
---

# CreateIdentityCollectorV1201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the created identity collector. | [optional] 
**Name** | Pointer to **string** | The display name of the created identity collector. | [optional] 

## Methods

### NewCreateIdentityCollectorV1201Response

`func NewCreateIdentityCollectorV1201Response() *CreateIdentityCollectorV1201Response`

NewCreateIdentityCollectorV1201Response instantiates a new CreateIdentityCollectorV1201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIdentityCollectorV1201ResponseWithDefaults

`func NewCreateIdentityCollectorV1201ResponseWithDefaults() *CreateIdentityCollectorV1201Response`

NewCreateIdentityCollectorV1201ResponseWithDefaults instantiates a new CreateIdentityCollectorV1201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateIdentityCollectorV1201Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateIdentityCollectorV1201Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateIdentityCollectorV1201Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateIdentityCollectorV1201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateIdentityCollectorV1201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIdentityCollectorV1201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIdentityCollectorV1201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateIdentityCollectorV1201Response) HasName() bool`

HasName returns a boolean if a field has been set.


