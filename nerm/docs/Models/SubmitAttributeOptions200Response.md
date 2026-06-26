---
id: nerm-submit-attribute-options200-response
title: SubmitAttributeOptions200Response
pagination_label: SubmitAttributeOptions200Response
sidebar_label: SubmitAttributeOptions200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitAttributeOptions200Response', 'NERMSubmitAttributeOptions200Response'] 
slug: /tools/sdk/go/nerm/models/submit-attribute-options200-response
tags: ['SDK', 'Software Development Kit', 'SubmitAttributeOptions200Response', 'NERMSubmitAttributeOptions200Response']
---

# SubmitAttributeOptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeAttributeOptions** | Pointer to [**[]AttributeOption**](attribute-option) |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubmitAttributeOptions200Response

`func NewSubmitAttributeOptions200Response() *SubmitAttributeOptions200Response`

NewSubmitAttributeOptions200Response instantiates a new SubmitAttributeOptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitAttributeOptions200ResponseWithDefaults

`func NewSubmitAttributeOptions200ResponseWithDefaults() *SubmitAttributeOptions200Response`

NewSubmitAttributeOptions200ResponseWithDefaults instantiates a new SubmitAttributeOptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeAttributeOptions

`func (o *SubmitAttributeOptions200Response) GetNeAttributeOptions() []AttributeOption`

GetNeAttributeOptions returns the NeAttributeOptions field if non-nil, zero value otherwise.

### GetNeAttributeOptionsOk

`func (o *SubmitAttributeOptions200Response) GetNeAttributeOptionsOk() (*[]AttributeOption, bool)`

GetNeAttributeOptionsOk returns a tuple with the NeAttributeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeAttributeOptions

`func (o *SubmitAttributeOptions200Response) SetNeAttributeOptions(v []AttributeOption)`

SetNeAttributeOptions sets NeAttributeOptions field to given value.

### HasNeAttributeOptions

`func (o *SubmitAttributeOptions200Response) HasNeAttributeOptions() bool`

HasNeAttributeOptions returns a boolean if a field has been set.

### GetInfo

`func (o *SubmitAttributeOptions200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SubmitAttributeOptions200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SubmitAttributeOptions200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SubmitAttributeOptions200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *SubmitAttributeOptions200Response) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *SubmitAttributeOptions200Response) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *SubmitAttributeOptions200Response) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *SubmitAttributeOptions200Response) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *SubmitAttributeOptions200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitAttributeOptions200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitAttributeOptions200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmitAttributeOptions200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


