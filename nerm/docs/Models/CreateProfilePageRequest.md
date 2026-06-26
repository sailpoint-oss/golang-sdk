---
id: nerm-create-profile-page-request
title: CreateProfilePageRequest
pagination_label: CreateProfilePageRequest
sidebar_label: CreateProfilePageRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateProfilePageRequest', 'NERMCreateProfilePageRequest'] 
slug: /tools/sdk/go/nerm/models/create-profile-page-request
tags: ['SDK', 'Software Development Kit', 'CreateProfilePageRequest', 'NERMCreateProfilePageRequest']
---

# CreateProfilePageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ProfilePage**](profile-page) |  | [optional] 

## Methods

### NewCreateProfilePageRequest

`func NewCreateProfilePageRequest() *CreateProfilePageRequest`

NewCreateProfilePageRequest instantiates a new CreateProfilePageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfilePageRequestWithDefaults

`func NewCreateProfilePageRequestWithDefaults() *CreateProfilePageRequest`

NewCreateProfilePageRequestWithDefaults instantiates a new CreateProfilePageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *CreateProfilePageRequest) GetPage() ProfilePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CreateProfilePageRequest) GetPageOk() (*ProfilePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CreateProfilePageRequest) SetPage(v ProfilePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *CreateProfilePageRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.


