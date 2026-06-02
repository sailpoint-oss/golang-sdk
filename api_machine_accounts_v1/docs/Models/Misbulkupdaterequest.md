---
id: v1-misbulkupdaterequest
title: Misbulkupdaterequest
pagination_label: Misbulkupdaterequest
sidebar_label: Misbulkupdaterequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Misbulkupdaterequest', 'V1Misbulkupdaterequest'] 
slug: /tools/sdk/go/machineaccountsv1/models/misbulkupdaterequest
tags: ['SDK', 'Software Development Kit', 'Misbulkupdaterequest', 'V1Misbulkupdaterequest']
---

# Misbulkupdaterequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | Machine identity or machine account IDs to update. | 
**JsonPatch** | [**[]Jsonpatchoperation**](jsonpatchoperation) | JSON Patch operations to apply to each ID in the request (RFC 6902). | 

## Methods

### NewMisbulkupdaterequest

`func NewMisbulkupdaterequest(ids []string, jsonPatch []Jsonpatchoperation, ) *Misbulkupdaterequest`

NewMisbulkupdaterequest instantiates a new Misbulkupdaterequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMisbulkupdaterequestWithDefaults

`func NewMisbulkupdaterequestWithDefaults() *Misbulkupdaterequest`

NewMisbulkupdaterequestWithDefaults instantiates a new Misbulkupdaterequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *Misbulkupdaterequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *Misbulkupdaterequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *Misbulkupdaterequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetJsonPatch

`func (o *Misbulkupdaterequest) GetJsonPatch() []Jsonpatchoperation`

GetJsonPatch returns the JsonPatch field if non-nil, zero value otherwise.

### GetJsonPatchOk

`func (o *Misbulkupdaterequest) GetJsonPatchOk() (*[]Jsonpatchoperation, bool)`

GetJsonPatchOk returns a tuple with the JsonPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPatch

`func (o *Misbulkupdaterequest) SetJsonPatch(v []Jsonpatchoperation)`

SetJsonPatch sets JsonPatch field to given value.



