---
id: v1-intelrisklinks
title: Intelrisklinks
pagination_label: Intelrisklinks
sidebar_label: Intelrisklinks
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelrisklinks', 'V1Intelrisklinks'] 
slug: /tools/sdk/go/intelligencepackage/models/intelrisklinks
tags: ['SDK', 'Software Development Kit', 'Intelrisklinks', 'V1Intelrisklinks']
---

# Intelrisklinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outliers** | Pointer to [**Intelhref**](intelhref) | Link to fetch the next outlier page for the same identity. | [optional] 

## Methods

### NewIntelrisklinks

`func NewIntelrisklinks() *Intelrisklinks`

NewIntelrisklinks instantiates a new Intelrisklinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelrisklinksWithDefaults

`func NewIntelrisklinksWithDefaults() *Intelrisklinks`

NewIntelrisklinksWithDefaults instantiates a new Intelrisklinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutliers

`func (o *Intelrisklinks) GetOutliers() Intelhref`

GetOutliers returns the Outliers field if non-nil, zero value otherwise.

### GetOutliersOk

`func (o *Intelrisklinks) GetOutliersOk() (*Intelhref, bool)`

GetOutliersOk returns a tuple with the Outliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutliers

`func (o *Intelrisklinks) SetOutliers(v Intelhref)`

SetOutliers sets Outliers field to given value.

### HasOutliers

`func (o *Intelrisklinks) HasOutliers() bool`

HasOutliers returns a boolean if a field has been set.


