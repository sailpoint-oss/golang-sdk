---
id: nerm-get-identity-proofing-results200-response
title: GetIdentityProofingResults200Response
pagination_label: GetIdentityProofingResults200Response
sidebar_label: GetIdentityProofingResults200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetIdentityProofingResults200Response', 'NERMGetIdentityProofingResults200Response'] 
slug: /tools/sdk/go/nerm/models/get-identity-proofing-results200-response
tags: ['SDK', 'Software Development Kit', 'GetIdentityProofingResults200Response', 'NERMGetIdentityProofingResults200Response']
---

# GetIdentityProofingResults200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProofingResults** | Pointer to [**[]IdentityProofingResult**](identity-proofing-result) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewGetIdentityProofingResults200Response

`func NewGetIdentityProofingResults200Response() *GetIdentityProofingResults200Response`

NewGetIdentityProofingResults200Response instantiates a new GetIdentityProofingResults200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIdentityProofingResults200ResponseWithDefaults

`func NewGetIdentityProofingResults200ResponseWithDefaults() *GetIdentityProofingResults200Response`

NewGetIdentityProofingResults200ResponseWithDefaults instantiates a new GetIdentityProofingResults200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProofingResults

`func (o *GetIdentityProofingResults200Response) GetIdentityProofingResults() []IdentityProofingResult`

GetIdentityProofingResults returns the IdentityProofingResults field if non-nil, zero value otherwise.

### GetIdentityProofingResultsOk

`func (o *GetIdentityProofingResults200Response) GetIdentityProofingResultsOk() (*[]IdentityProofingResult, bool)`

GetIdentityProofingResultsOk returns a tuple with the IdentityProofingResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProofingResults

`func (o *GetIdentityProofingResults200Response) SetIdentityProofingResults(v []IdentityProofingResult)`

SetIdentityProofingResults sets IdentityProofingResults field to given value.

### HasIdentityProofingResults

`func (o *GetIdentityProofingResults200Response) HasIdentityProofingResults() bool`

HasIdentityProofingResults returns a boolean if a field has been set.

### GetMetadata

`func (o *GetIdentityProofingResults200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetIdentityProofingResults200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetIdentityProofingResults200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetIdentityProofingResults200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


