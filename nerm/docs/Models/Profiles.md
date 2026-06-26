---
id: nerm-profiles
title: Profiles
pagination_label: Profiles
sidebar_label: Profiles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Profiles', 'NERMProfiles'] 
slug: /tools/sdk/go/nerm/models/profiles
tags: ['SDK', 'Software Development Kit', 'Profiles', 'NERMProfiles']
---

# Profiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]Profile**](profile) |  | [optional] 

## Methods

### NewProfiles

`func NewProfiles() *Profiles`

NewProfiles instantiates a new Profiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilesWithDefaults

`func NewProfilesWithDefaults() *Profiles`

NewProfilesWithDefaults instantiates a new Profiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *Profiles) GetProfiles() []Profile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *Profiles) GetProfilesOk() (*[]Profile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *Profiles) SetProfiles(v []Profile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *Profiles) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


