---
id: v1-intelidentityriskbody
title: Intelidentityriskbody
pagination_label: Intelidentityriskbody
sidebar_label: Intelidentityriskbody
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentityriskbody', 'V1Intelidentityriskbody'] 
slug: /tools/sdk/go/intelligencepackage/models/intelidentityriskbody
tags: ['SDK', 'Software Development Kit', 'Intelidentityriskbody', 'V1Intelidentityriskbody']
---

# Intelidentityriskbody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outliers** | [**[]Inteloutlieraccessitem**](inteloutlieraccessitem) | Page of outlier access-items associated with the resolved identity outlier. | 
**OutliersTotal** | **NullableInt64** | Total available outlier access-item count from upstream. | 
**Links** | Pointer to [**NullableIntelrisklinks**](intelrisklinks) | Continuation links map; omitted when no additional page exists. | [optional] 

## Methods

### NewIntelidentityriskbody

`func NewIntelidentityriskbody(outliers []Inteloutlieraccessitem, outliersTotal NullableInt64, ) *Intelidentityriskbody`

NewIntelidentityriskbody instantiates a new Intelidentityriskbody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentityriskbodyWithDefaults

`func NewIntelidentityriskbodyWithDefaults() *Intelidentityriskbody`

NewIntelidentityriskbodyWithDefaults instantiates a new Intelidentityriskbody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutliers

`func (o *Intelidentityriskbody) GetOutliers() []Inteloutlieraccessitem`

GetOutliers returns the Outliers field if non-nil, zero value otherwise.

### GetOutliersOk

`func (o *Intelidentityriskbody) GetOutliersOk() (*[]Inteloutlieraccessitem, bool)`

GetOutliersOk returns a tuple with the Outliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutliers

`func (o *Intelidentityriskbody) SetOutliers(v []Inteloutlieraccessitem)`

SetOutliers sets Outliers field to given value.


### GetOutliersTotal

`func (o *Intelidentityriskbody) GetOutliersTotal() int64`

GetOutliersTotal returns the OutliersTotal field if non-nil, zero value otherwise.

### GetOutliersTotalOk

`func (o *Intelidentityriskbody) GetOutliersTotalOk() (*int64, bool)`

GetOutliersTotalOk returns a tuple with the OutliersTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutliersTotal

`func (o *Intelidentityriskbody) SetOutliersTotal(v int64)`

SetOutliersTotal sets OutliersTotal field to given value.


### SetOutliersTotalNil

`func (o *Intelidentityriskbody) SetOutliersTotalNil(b bool)`

 SetOutliersTotalNil sets the value for OutliersTotal to be an explicit nil

### UnsetOutliersTotal
`func (o *Intelidentityriskbody) UnsetOutliersTotal()`

UnsetOutliersTotal ensures that no value is present for OutliersTotal, not even an explicit nil
### GetLinks

`func (o *Intelidentityriskbody) GetLinks() Intelrisklinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Intelidentityriskbody) GetLinksOk() (*Intelrisklinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Intelidentityriskbody) SetLinks(v Intelrisklinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Intelidentityriskbody) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *Intelidentityriskbody) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *Intelidentityriskbody) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

