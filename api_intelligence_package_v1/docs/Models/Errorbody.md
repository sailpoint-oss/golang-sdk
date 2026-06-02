---
id: v1-errorbody
title: Errorbody
pagination_label: Errorbody
sidebar_label: Errorbody
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Errorbody', 'V1Errorbody'] 
slug: /tools/sdk/go/intelligencepackagev1/models/errorbody
tags: ['SDK', 'Software Development Kit', 'Errorbody', 'V1Errorbody']
---

# Errorbody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailCode** | Pointer to **string** | Machine-readable error code returned by the Intelligence Package service. | [optional] 
**Message** | Pointer to **string** | Human-readable explanation of the error suitable for client logging. | [optional] 

## Methods

### NewErrorbody

`func NewErrorbody() *Errorbody`

NewErrorbody instantiates a new Errorbody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorbodyWithDefaults

`func NewErrorbodyWithDefaults() *Errorbody`

NewErrorbodyWithDefaults instantiates a new Errorbody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailCode

`func (o *Errorbody) GetDetailCode() string`

GetDetailCode returns the DetailCode field if non-nil, zero value otherwise.

### GetDetailCodeOk

`func (o *Errorbody) GetDetailCodeOk() (*string, bool)`

GetDetailCodeOk returns a tuple with the DetailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailCode

`func (o *Errorbody) SetDetailCode(v string)`

SetDetailCode sets DetailCode field to given value.

### HasDetailCode

`func (o *Errorbody) HasDetailCode() bool`

HasDetailCode returns a boolean if a field has been set.

### GetMessage

`func (o *Errorbody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Errorbody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Errorbody) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Errorbody) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


