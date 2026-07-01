---
id: v1-intelidentityambiguousbody
title: Intelidentityambiguousbody
pagination_label: Intelidentityambiguousbody
sidebar_label: Intelidentityambiguousbody
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentityambiguousbody', 'V1Intelidentityambiguousbody'] 
slug: /tools/sdk/go/intelligence/models/intelidentityambiguousbody
tags: ['SDK', 'Software Development Kit', 'Intelidentityambiguousbody', 'V1Intelidentityambiguousbody']
---

# Intelidentityambiguousbody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailCode** | **string** | Constant detail code indicating that more than one identity matched the filter. | 
**Message** | Pointer to **string** | Optional explanatory text describing why the filter was considered ambiguous. | [optional] 
**Candidates** | [**[]Intelidentityambiguouscandidate**](intelidentityambiguouscandidate) | Collection of identities that matched the ambiguous filter expression. | 

## Methods

### NewIntelidentityambiguousbody

`func NewIntelidentityambiguousbody(detailCode string, candidates []Intelidentityambiguouscandidate, ) *Intelidentityambiguousbody`

NewIntelidentityambiguousbody instantiates a new Intelidentityambiguousbody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentityambiguousbodyWithDefaults

`func NewIntelidentityambiguousbodyWithDefaults() *Intelidentityambiguousbody`

NewIntelidentityambiguousbodyWithDefaults instantiates a new Intelidentityambiguousbody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailCode

`func (o *Intelidentityambiguousbody) GetDetailCode() string`

GetDetailCode returns the DetailCode field if non-nil, zero value otherwise.

### GetDetailCodeOk

`func (o *Intelidentityambiguousbody) GetDetailCodeOk() (*string, bool)`

GetDetailCodeOk returns a tuple with the DetailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailCode

`func (o *Intelidentityambiguousbody) SetDetailCode(v string)`

SetDetailCode sets DetailCode field to given value.


### GetMessage

`func (o *Intelidentityambiguousbody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Intelidentityambiguousbody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Intelidentityambiguousbody) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Intelidentityambiguousbody) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCandidates

`func (o *Intelidentityambiguousbody) GetCandidates() []Intelidentityambiguouscandidate`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *Intelidentityambiguousbody) GetCandidatesOk() (*[]Intelidentityambiguouscandidate, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *Intelidentityambiguousbody) SetCandidates(v []Intelidentityambiguouscandidate)`

SetCandidates sets Candidates field to given value.



