---
id: nerm-page-element
title: PageElement
pagination_label: PageElement
sidebar_label: PageElement
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageElement', 'NERMPageElement'] 
slug: /tools/sdk/go/nerm/models/page-element
tags: ['SDK', 'Software Development Kit', 'PageElement', 'NERMPageElement']
---

# PageElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the page element | [optional] [readonly] 
**Uid** | Pointer to **string** | The user-specified identifier for the record | [optional] 
**ElementType** | **string** | The type of content on the page. | 
**PageUid** | Pointer to **string** | The user-specified identifier of the page record this applies to; one of page_id or page_uid must be present. | [optional] 
**PageId** | Pointer to **string** | The id of the page record this applies to; one of page_id or page_uid must be present. | [optional] 
**ElementUid** | Pointer to **string** | The user-specified identifier of the record (Form or PageContent) this applies to; one of element_id or element_uid must be present. | [optional] 
**ElementId** | Pointer to **string** | The id of the record (Form or PageContent) this applies to; one of element_id or element_uid must be present. | [optional] 
**Order** | Pointer to **int32** | The position of the attribute in the ProfileType's naming | [optional] 

## Methods

### NewPageElement

`func NewPageElement(elementType string, ) *PageElement`

NewPageElement instantiates a new PageElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageElementWithDefaults

`func NewPageElementWithDefaults() *PageElement`

NewPageElementWithDefaults instantiates a new PageElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageElement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageElement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUid

`func (o *PageElement) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PageElement) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PageElement) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *PageElement) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetElementType

`func (o *PageElement) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *PageElement) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *PageElement) SetElementType(v string)`

SetElementType sets ElementType field to given value.


### GetPageUid

`func (o *PageElement) GetPageUid() string`

GetPageUid returns the PageUid field if non-nil, zero value otherwise.

### GetPageUidOk

`func (o *PageElement) GetPageUidOk() (*string, bool)`

GetPageUidOk returns a tuple with the PageUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUid

`func (o *PageElement) SetPageUid(v string)`

SetPageUid sets PageUid field to given value.

### HasPageUid

`func (o *PageElement) HasPageUid() bool`

HasPageUid returns a boolean if a field has been set.

### GetPageId

`func (o *PageElement) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PageElement) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PageElement) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *PageElement) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetElementUid

`func (o *PageElement) GetElementUid() string`

GetElementUid returns the ElementUid field if non-nil, zero value otherwise.

### GetElementUidOk

`func (o *PageElement) GetElementUidOk() (*string, bool)`

GetElementUidOk returns a tuple with the ElementUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementUid

`func (o *PageElement) SetElementUid(v string)`

SetElementUid sets ElementUid field to given value.

### HasElementUid

`func (o *PageElement) HasElementUid() bool`

HasElementUid returns a boolean if a field has been set.

### GetElementId

`func (o *PageElement) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *PageElement) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *PageElement) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *PageElement) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetOrder

`func (o *PageElement) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PageElement) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PageElement) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PageElement) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


