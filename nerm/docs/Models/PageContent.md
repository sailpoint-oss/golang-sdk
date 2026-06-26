---
id: nerm-page-content
title: PageContent
pagination_label: PageContent
sidebar_label: PageContent
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageContent', 'NERMPageContent'] 
slug: /tools/sdk/go/nerm/models/page-content
tags: ['SDK', 'Software Development Kit', 'PageContent', 'NERMPageContent']
---

# PageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the page content | [optional] [readonly] 
**Uid** | Pointer to **string** | The user-specified identifier for the record | [optional] 
**Type** | **string** | The type of content on the page. | 
**Content** | Pointer to **string** | The text content to present in this page content record. | [optional] 
**CreatedAt** | Pointer to **SailPointTime** | The date-time the record created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **SailPointTime** | The date-time the record was last updated. | [optional] [readonly] 

## Methods

### NewPageContent

`func NewPageContent(type_ string, ) *PageContent`

NewPageContent instantiates a new PageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageContentWithDefaults

`func NewPageContentWithDefaults() *PageContent`

NewPageContentWithDefaults instantiates a new PageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageContent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageContent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageContent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageContent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *PageContent) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PageContent) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PageContent) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *PageContent) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetType

`func (o *PageContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PageContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PageContent) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *PageContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageContent) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageContent) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PageContent) GetCreatedAt() SailPointTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageContent) GetCreatedAtOk() (*SailPointTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageContent) SetCreatedAt(v SailPointTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageContent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PageContent) GetUpdatedAt() SailPointTime`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PageContent) GetUpdatedAtOk() (*SailPointTime, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PageContent) SetUpdatedAt(v SailPointTime)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PageContent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


