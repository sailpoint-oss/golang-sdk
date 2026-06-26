---
id: nerm-get-job-status200-response
title: GetJobStatus200Response
pagination_label: GetJobStatus200Response
sidebar_label: GetJobStatus200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetJobStatus200Response', 'NERMGetJobStatus200Response'] 
slug: /tools/sdk/go/nerm/models/get-job-status200-response
tags: ['SDK', 'Software Development Kit', 'GetJobStatus200Response', 'NERMGetJobStatus200Response']
---

# GetJobStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**JobData** | Pointer to [**[]GetJobStatus200ResponseJobDataInner**](get-job-status200-response-job-data-inner) |  | [optional] 

## Methods

### NewGetJobStatus200Response

`func NewGetJobStatus200Response() *GetJobStatus200Response`

NewGetJobStatus200Response instantiates a new GetJobStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobStatus200ResponseWithDefaults

`func NewGetJobStatus200ResponseWithDefaults() *GetJobStatus200Response`

NewGetJobStatus200ResponseWithDefaults instantiates a new GetJobStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *GetJobStatus200Response) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetJobStatus200Response) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetJobStatus200Response) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetJobStatus200Response) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetStatus

`func (o *GetJobStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJobStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobType

`func (o *GetJobStatus200Response) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *GetJobStatus200Response) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *GetJobStatus200Response) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *GetJobStatus200Response) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetJobData

`func (o *GetJobStatus200Response) GetJobData() []GetJobStatus200ResponseJobDataInner`

GetJobData returns the JobData field if non-nil, zero value otherwise.

### GetJobDataOk

`func (o *GetJobStatus200Response) GetJobDataOk() (*[]GetJobStatus200ResponseJobDataInner, bool)`

GetJobDataOk returns a tuple with the JobData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobData

`func (o *GetJobStatus200Response) SetJobData(v []GetJobStatus200ResponseJobDataInner)`

SetJobData sets JobData field to given value.

### HasJobData

`func (o *GetJobStatus200Response) HasJobData() bool`

HasJobData returns a boolean if a field has been set.


