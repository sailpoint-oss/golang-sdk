---
id: nerm-delegations-get200-response
title: DelegationsGet200Response
pagination_label: DelegationsGet200Response
sidebar_label: DelegationsGet200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DelegationsGet200Response', 'NERMDelegationsGet200Response'] 
slug: /tools/sdk/go/nerm/models/delegations-get200-response
tags: ['SDK', 'Software Development Kit', 'DelegationsGet200Response', 'NERMDelegationsGet200Response']
---

# DelegationsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegations** | Pointer to [**[]Delegation**](delegation) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](metadata) |  | [optional] 

## Methods

### NewDelegationsGet200Response

`func NewDelegationsGet200Response() *DelegationsGet200Response`

NewDelegationsGet200Response instantiates a new DelegationsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationsGet200ResponseWithDefaults

`func NewDelegationsGet200ResponseWithDefaults() *DelegationsGet200Response`

NewDelegationsGet200ResponseWithDefaults instantiates a new DelegationsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegations

`func (o *DelegationsGet200Response) GetDelegations() []Delegation`

GetDelegations returns the Delegations field if non-nil, zero value otherwise.

### GetDelegationsOk

`func (o *DelegationsGet200Response) GetDelegationsOk() (*[]Delegation, bool)`

GetDelegationsOk returns a tuple with the Delegations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegations

`func (o *DelegationsGet200Response) SetDelegations(v []Delegation)`

SetDelegations sets Delegations field to given value.

### HasDelegations

`func (o *DelegationsGet200Response) HasDelegations() bool`

HasDelegations returns a boolean if a field has been set.

### GetMetadata

`func (o *DelegationsGet200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DelegationsGet200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DelegationsGet200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DelegationsGet200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


