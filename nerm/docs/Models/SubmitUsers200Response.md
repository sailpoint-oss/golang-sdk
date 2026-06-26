---
id: nerm-submit-users200-response
title: SubmitUsers200Response
pagination_label: SubmitUsers200Response
sidebar_label: SubmitUsers200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUsers200Response', 'NERMSubmitUsers200Response'] 
slug: /tools/sdk/go/nerm/models/submit-users200-response
tags: ['SDK', 'Software Development Kit', 'SubmitUsers200Response', 'NERMSubmitUsers200Response']
---

# SubmitUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]User**](user) |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubmitUsers200Response

`func NewSubmitUsers200Response() *SubmitUsers200Response`

NewSubmitUsers200Response instantiates a new SubmitUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUsers200ResponseWithDefaults

`func NewSubmitUsers200ResponseWithDefaults() *SubmitUsers200Response`

NewSubmitUsers200ResponseWithDefaults instantiates a new SubmitUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *SubmitUsers200Response) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SubmitUsers200Response) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SubmitUsers200Response) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *SubmitUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetInfo

`func (o *SubmitUsers200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SubmitUsers200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SubmitUsers200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SubmitUsers200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *SubmitUsers200Response) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *SubmitUsers200Response) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *SubmitUsers200Response) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *SubmitUsers200Response) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *SubmitUsers200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitUsers200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitUsers200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmitUsers200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


