---
id: v2026-start-application-discovery403-response
title: StartApplicationDiscovery403Response
pagination_label: StartApplicationDiscovery403Response
sidebar_label: StartApplicationDiscovery403Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'StartApplicationDiscovery403Response', 'V2026StartApplicationDiscovery403Response'] 
slug: /tools/sdk/go/v2026/models/start-application-discovery403-response
tags: ['SDK', 'Software Development Kit', 'StartApplicationDiscovery403Response', 'V2026StartApplicationDiscovery403Response']
---

# StartApplicationDiscovery403Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailCode** | Pointer to **string** | Fine-grained error code providing more detail of the error. | [optional] 
**TrackingId** | Pointer to **string** | Unique tracking id for the error. | [optional] 
**Messages** | Pointer to [**[]ErrorMessageDto**](error-message-dto) | Generic localized reason for error | [optional] 
**Causes** | Pointer to [**[]ErrorMessageDto**](error-message-dto) | Plain-text descriptive reasons to provide additional detail to the text provided in the messages field | [optional] 
**Error** | **string** | Error message when quota is exceeded | 

## Methods

### NewStartApplicationDiscovery403Response

`func NewStartApplicationDiscovery403Response(error_ string, ) *StartApplicationDiscovery403Response`

NewStartApplicationDiscovery403Response instantiates a new StartApplicationDiscovery403Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartApplicationDiscovery403ResponseWithDefaults

`func NewStartApplicationDiscovery403ResponseWithDefaults() *StartApplicationDiscovery403Response`

NewStartApplicationDiscovery403ResponseWithDefaults instantiates a new StartApplicationDiscovery403Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailCode

`func (o *StartApplicationDiscovery403Response) GetDetailCode() string`

GetDetailCode returns the DetailCode field if non-nil, zero value otherwise.

### GetDetailCodeOk

`func (o *StartApplicationDiscovery403Response) GetDetailCodeOk() (*string, bool)`

GetDetailCodeOk returns a tuple with the DetailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailCode

`func (o *StartApplicationDiscovery403Response) SetDetailCode(v string)`

SetDetailCode sets DetailCode field to given value.

### HasDetailCode

`func (o *StartApplicationDiscovery403Response) HasDetailCode() bool`

HasDetailCode returns a boolean if a field has been set.

### GetTrackingId

`func (o *StartApplicationDiscovery403Response) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *StartApplicationDiscovery403Response) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *StartApplicationDiscovery403Response) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *StartApplicationDiscovery403Response) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### GetMessages

`func (o *StartApplicationDiscovery403Response) GetMessages() []ErrorMessageDto`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *StartApplicationDiscovery403Response) GetMessagesOk() (*[]ErrorMessageDto, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *StartApplicationDiscovery403Response) SetMessages(v []ErrorMessageDto)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *StartApplicationDiscovery403Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetCauses

`func (o *StartApplicationDiscovery403Response) GetCauses() []ErrorMessageDto`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *StartApplicationDiscovery403Response) GetCausesOk() (*[]ErrorMessageDto, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *StartApplicationDiscovery403Response) SetCauses(v []ErrorMessageDto)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *StartApplicationDiscovery403Response) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetError

`func (o *StartApplicationDiscovery403Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StartApplicationDiscovery403Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StartApplicationDiscovery403Response) SetError(v string)`

SetError sets Error field to given value.



