---
id: nerm-job-job-status
title: JobJobStatus
pagination_label: JobJobStatus
sidebar_label: JobJobStatus
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'JobJobStatus', 'NERMJobJobStatus'] 
slug: /tools/sdk/go/nerm/models/job-job-status
tags: ['SDK', 'Software Development Kit', 'JobJobStatus', 'NERMJobJobStatus']
---

# JobJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewJobJobStatus

`func NewJobJobStatus() *JobJobStatus`

NewJobJobStatus instantiates a new JobJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobJobStatusWithDefaults

`func NewJobJobStatusWithDefaults() *JobJobStatus`

NewJobJobStatusWithDefaults instantiates a new JobJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobJobStatus) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobJobStatus) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobJobStatus) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *JobJobStatus) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStatus

`func (o *JobJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobJobStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


