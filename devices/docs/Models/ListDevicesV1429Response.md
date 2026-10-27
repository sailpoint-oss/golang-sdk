---
id: v1-list-devices-v1429-response
title: ListDevicesV1429Response
pagination_label: ListDevicesV1429Response
sidebar_label: ListDevicesV1429Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ListDevicesV1429Response', 'V1ListDevicesV1429Response'] 
slug: /tools/sdk/go/devices/models/list-devices-v1429-response
tags: ['SDK', 'Software Development Kit', 'ListDevicesV1429Response', 'V1ListDevicesV1429Response']
---

# ListDevicesV1429Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **interface{}** | A message describing the error | [optional] 

## Methods

### NewListDevicesV1429Response

`func NewListDevicesV1429Response() *ListDevicesV1429Response`

NewListDevicesV1429Response instantiates a new ListDevicesV1429Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDevicesV1429ResponseWithDefaults

`func NewListDevicesV1429ResponseWithDefaults() *ListDevicesV1429Response`

NewListDevicesV1429ResponseWithDefaults instantiates a new ListDevicesV1429Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ListDevicesV1429Response) GetMessage() interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListDevicesV1429Response) GetMessageOk() (*interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListDevicesV1429Response) SetMessage(v interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListDevicesV1429Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ListDevicesV1429Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ListDevicesV1429Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

