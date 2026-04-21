---
id: nerm-page-element1
title: PageElement1
pagination_label: PageElement1
sidebar_label: PageElement1
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageElement1', 'NERMPageElement1'] 
slug: /tools/sdk/go/nerm/models/page-element1
tags: ['SDK', 'Software Development Kit', 'PageElement1', 'NERMPageElement1']
---

# PageElement1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the page content | [optional] [readonly] 
**Uid** | Pointer to **string** | The user-specified identifier for the record | [optional] 
**ElementType** | **string** | The type of content on the page. | 
**PageUid** | Pointer to **string** | The user-specified identifier of the page record this applies to; one of page_id or page_uid must be present. | [optional] 
**PageId** | Pointer to **string** | The id of the page record this applies to; one of page_id or page_uid must be present. | [optional] 
**ElementUid** | Pointer to **string** | The user-specified identifier of the record (Form or PageContent) this applies to; one of element_id or element_uid must be present. | [optional] 
**ElementId** | Pointer to **string** | The id of the record (Form or PageContent) this applies to; one of element_id or element_uid must be present. | [optional] 
**Order** | Pointer to **int32** | The ordinal position of the attribute in the user-interface page. | [optional] 

## Methods

### NewPageElement1

`func NewPageElement1(elementType string, ) *PageElement1`

NewPageElement1 instantiates a new PageElement1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageElement1WithDefaults

`func NewPageElement1WithDefaults() *PageElement1`

NewPageElement1WithDefaults instantiates a new PageElement1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageElement1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageElement1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageElement1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageElement1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *PageElement1) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PageElement1) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PageElement1) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *PageElement1) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetElementType

`func (o *PageElement1) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *PageElement1) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *PageElement1) SetElementType(v string)`

SetElementType sets ElementType field to given value.


### GetPageUid

`func (o *PageElement1) GetPageUid() string`

GetPageUid returns the PageUid field if non-nil, zero value otherwise.

### GetPageUidOk

`func (o *PageElement1) GetPageUidOk() (*string, bool)`

GetPageUidOk returns a tuple with the PageUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUid

`func (o *PageElement1) SetPageUid(v string)`

SetPageUid sets PageUid field to given value.

### HasPageUid

`func (o *PageElement1) HasPageUid() bool`

HasPageUid returns a boolean if a field has been set.

### GetPageId

`func (o *PageElement1) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PageElement1) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PageElement1) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *PageElement1) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetElementUid

`func (o *PageElement1) GetElementUid() string`

GetElementUid returns the ElementUid field if non-nil, zero value otherwise.

### GetElementUidOk

`func (o *PageElement1) GetElementUidOk() (*string, bool)`

GetElementUidOk returns a tuple with the ElementUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementUid

`func (o *PageElement1) SetElementUid(v string)`

SetElementUid sets ElementUid field to given value.

### HasElementUid

`func (o *PageElement1) HasElementUid() bool`

HasElementUid returns a boolean if a field has been set.

### GetElementId

`func (o *PageElement1) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *PageElement1) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *PageElement1) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *PageElement1) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetOrder

`func (o *PageElement1) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PageElement1) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PageElement1) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PageElement1) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


