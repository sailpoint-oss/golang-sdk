---
id: nerm-get-schema-mapped-profiles-collection200-response
title: GetSchemaMappedProfilesCollection200Response
pagination_label: GetSchemaMappedProfilesCollection200Response
sidebar_label: GetSchemaMappedProfilesCollection200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetSchemaMappedProfilesCollection200Response', 'NERMGetSchemaMappedProfilesCollection200Response'] 
slug: /tools/sdk/go/nerm/models/get-schema-mapped-profiles-collection200-response
tags: ['SDK', 'Software Development Kit', 'GetSchemaMappedProfilesCollection200Response', 'NERMGetSchemaMappedProfilesCollection200Response']
---

# GetSchemaMappedProfilesCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]Profile**](profile) |  | [optional] 
**Metadata** | Pointer to [**MetadataWithAfterId**](metadata-with-after-id) |  | [optional] 

## Methods

### NewGetSchemaMappedProfilesCollection200Response

`func NewGetSchemaMappedProfilesCollection200Response() *GetSchemaMappedProfilesCollection200Response`

NewGetSchemaMappedProfilesCollection200Response instantiates a new GetSchemaMappedProfilesCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSchemaMappedProfilesCollection200ResponseWithDefaults

`func NewGetSchemaMappedProfilesCollection200ResponseWithDefaults() *GetSchemaMappedProfilesCollection200Response`

NewGetSchemaMappedProfilesCollection200ResponseWithDefaults instantiates a new GetSchemaMappedProfilesCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *GetSchemaMappedProfilesCollection200Response) GetProfiles() []Profile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetSchemaMappedProfilesCollection200Response) GetProfilesOk() (*[]Profile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetSchemaMappedProfilesCollection200Response) SetProfiles(v []Profile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetSchemaMappedProfilesCollection200Response) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetMetadata

`func (o *GetSchemaMappedProfilesCollection200Response) GetMetadata() MetadataWithAfterId`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetSchemaMappedProfilesCollection200Response) GetMetadataOk() (*MetadataWithAfterId, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetSchemaMappedProfilesCollection200Response) SetMetadata(v MetadataWithAfterId)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetSchemaMappedProfilesCollection200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


