---
id: v1-misbulkresponse
title: Misbulkresponse
pagination_label: Misbulkresponse
sidebar_label: Misbulkresponse
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Misbulkresponse', 'V1Misbulkresponse'] 
slug: /tools/sdk/go/machineaccountsv1/models/misbulkresponse
tags: ['SDK', 'Software Development Kit', 'Misbulkresponse', 'V1Misbulkresponse']
---

# Misbulkresponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Machine identity or machine account ID from the request. | 
**Status** | **int32** | HTTP status code for this ID. For example, 200 indicates success, 404 indicates the resource was not found or is not accessible to the caller, and 409 indicates a conflict such as a duplicate ID in the batch. | 
**Message** | Pointer to **string** | Human-readable detail for this result. | [optional] 

## Methods

### NewMisbulkresponse

`func NewMisbulkresponse(id string, status int32, ) *Misbulkresponse`

NewMisbulkresponse instantiates a new Misbulkresponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMisbulkresponseWithDefaults

`func NewMisbulkresponseWithDefaults() *Misbulkresponse`

NewMisbulkresponseWithDefaults instantiates a new Misbulkresponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Misbulkresponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Misbulkresponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Misbulkresponse) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Misbulkresponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Misbulkresponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Misbulkresponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *Misbulkresponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Misbulkresponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Misbulkresponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Misbulkresponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


