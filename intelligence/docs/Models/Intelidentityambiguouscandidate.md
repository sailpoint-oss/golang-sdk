---
id: v1-intelidentityambiguouscandidate
title: Intelidentityambiguouscandidate
pagination_label: Intelidentityambiguouscandidate
sidebar_label: Intelidentityambiguouscandidate
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelidentityambiguouscandidate', 'V1Intelidentityambiguouscandidate'] 
slug: /tools/sdk/go/intelligence/models/intelidentityambiguouscandidate
tags: ['SDK', 'Software Development Kit', 'Intelidentityambiguouscandidate', 'V1Intelidentityambiguouscandidate']
---

# Intelidentityambiguouscandidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identity identifier for one of the ambiguous matching identities. | 
**DisplayName** | Pointer to **string** | Display name for the ambiguous matching identity when available. | [optional] 

## Methods

### NewIntelidentityambiguouscandidate

`func NewIntelidentityambiguouscandidate(id string, ) *Intelidentityambiguouscandidate`

NewIntelidentityambiguouscandidate instantiates a new Intelidentityambiguouscandidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelidentityambiguouscandidateWithDefaults

`func NewIntelidentityambiguouscandidateWithDefaults() *Intelidentityambiguouscandidate`

NewIntelidentityambiguouscandidateWithDefaults instantiates a new Intelidentityambiguouscandidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Intelidentityambiguouscandidate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Intelidentityambiguouscandidate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Intelidentityambiguouscandidate) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *Intelidentityambiguouscandidate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Intelidentityambiguouscandidate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Intelidentityambiguouscandidate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Intelidentityambiguouscandidate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


