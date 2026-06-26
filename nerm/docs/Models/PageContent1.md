---
id: nerm-page-content1
title: PageContent1
pagination_label: PageContent1
sidebar_label: PageContent1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageContent1', 'NERMPageContent1'] 
slug: /tools/sdk/go/nerm/models/page-content1
tags: ['SDK', 'Software Development Kit', 'PageContent1', 'NERMPageContent1']
---

# PageContent1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The user-specified identifier for the record | [optional] 
**Type** | **string** | The type of content on the page. | 
**Content** | Pointer to **string** | The text content to present in this page content record. | [optional] 

## Methods

### NewPageContent1

`func NewPageContent1(type_ string, ) *PageContent1`

NewPageContent1 instantiates a new PageContent1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageContent1WithDefaults

`func NewPageContent1WithDefaults() *PageContent1`

NewPageContent1WithDefaults instantiates a new PageContent1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *PageContent1) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PageContent1) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PageContent1) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *PageContent1) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *PageContent1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PageContent1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PageContent1) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *PageContent1) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageContent1) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageContent1) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageContent1) HasContent() bool`

HasContent returns a boolean if a field has been set.


