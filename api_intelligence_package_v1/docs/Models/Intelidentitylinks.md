---
id: v1-intelidentitylinks
title: Intelidentitylinks
pagination_label: Intelidentitylinks
sidebar_label: Intelidentitylinks
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentitylinks', 'V1Intelidentitylinks'] 
slug: /tools/sdk/go/intelligencepackagev1/models/intelidentitylinks
tags: ['SDK', 'Software Development Kit', 'Intelidentitylinks', 'V1Intelidentitylinks']
---

# Intelidentitylinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | [**Intelhref**](intelhref) | Hyperlink to the Intelligence Package access document for this identity. | 
**AccessHistory** | [**Intelhref**](intelhref) | Hyperlink to the Intelligence Package access history document for this identity. | 

## Methods

### NewIntelidentitylinks

`func NewIntelidentitylinks(access Intelhref, accessHistory Intelhref, ) *Intelidentitylinks`

NewIntelidentitylinks instantiates a new Intelidentitylinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentitylinksWithDefaults

`func NewIntelidentitylinksWithDefaults() *Intelidentitylinks`

NewIntelidentitylinksWithDefaults instantiates a new Intelidentitylinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *Intelidentitylinks) GetAccess() Intelhref`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Intelidentitylinks) GetAccessOk() (*Intelhref, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Intelidentitylinks) SetAccess(v Intelhref)`

SetAccess sets Access field to given value.


### GetAccessHistory

`func (o *Intelidentitylinks) GetAccessHistory() Intelhref`

GetAccessHistory returns the AccessHistory field if non-nil, zero value otherwise.

### GetAccessHistoryOk

`func (o *Intelidentitylinks) GetAccessHistoryOk() (*Intelhref, bool)`

GetAccessHistoryOk returns a tuple with the AccessHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessHistory

`func (o *Intelidentitylinks) SetAccessHistory(v Intelhref)`

SetAccessHistory sets AccessHistory field to given value.



