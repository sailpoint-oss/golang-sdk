---
id: nerm-job
title: Job
pagination_label: Job
sidebar_label: Job
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Job', 'NERMJob'] 
slug: /tools/sdk/go/nerm/models/job
tags: ['SDK', 'Software Development Kit', 'Job', 'NERMJob']
---

# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *Job) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Job) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Job) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Job) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *Job) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *Job) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *Job) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *Job) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *Job) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Job) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


