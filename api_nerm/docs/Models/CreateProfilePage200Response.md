---
id: nerm-create-profile-page200-response
title: CreateProfilePage200Response
pagination_label: CreateProfilePage200Response
sidebar_label: CreateProfilePage200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateProfilePage200Response', 'NERMCreateProfilePage200Response'] 
slug: /tools/sdk/go/nerm/models/create-profile-page200-response
tags: ['SDK', 'Software Development Kit', 'CreateProfilePage200Response', 'NERMCreateProfilePage200Response']
---

# CreateProfilePage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**Pages**](pages) |  | [optional] 

## Methods

### NewCreateProfilePage200Response

`func NewCreateProfilePage200Response() *CreateProfilePage200Response`

NewCreateProfilePage200Response instantiates a new CreateProfilePage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfilePage200ResponseWithDefaults

`func NewCreateProfilePage200ResponseWithDefaults() *CreateProfilePage200Response`

NewCreateProfilePage200ResponseWithDefaults instantiates a new CreateProfilePage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CreateProfilePage200Response) GetPage() Pages`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CreateProfilePage200Response) GetPageOk() (*Pages, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CreateProfilePage200Response) SetPage(v Pages)`

SetPage sets Page field to given value.

### HasPage

`func (o *CreateProfilePage200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.


