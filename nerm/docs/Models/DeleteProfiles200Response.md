---
id: nerm-delete-profiles200-response
title: DeleteProfiles200Response
pagination_label: DeleteProfiles200Response
sidebar_label: DeleteProfiles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DeleteProfiles200Response', 'NERMDeleteProfiles200Response'] 
slug: /tools/sdk/go/nerm/models/delete-profiles200-response
tags: ['SDK', 'Software Development Kit', 'DeleteProfiles200Response', 'NERMDeleteProfiles200Response']
---

# DeleteProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]Profile**](profile) |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeleteProfiles200Response

`func NewDeleteProfiles200Response() *DeleteProfiles200Response`

NewDeleteProfiles200Response instantiates a new DeleteProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteProfiles200ResponseWithDefaults

`func NewDeleteProfiles200ResponseWithDefaults() *DeleteProfiles200Response`

NewDeleteProfiles200ResponseWithDefaults instantiates a new DeleteProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *DeleteProfiles200Response) GetProfiles() []Profile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *DeleteProfiles200Response) GetProfilesOk() (*[]Profile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *DeleteProfiles200Response) SetProfiles(v []Profile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *DeleteProfiles200Response) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetInfo

`func (o *DeleteProfiles200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DeleteProfiles200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DeleteProfiles200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DeleteProfiles200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *DeleteProfiles200Response) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *DeleteProfiles200Response) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *DeleteProfiles200Response) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *DeleteProfiles200Response) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *DeleteProfiles200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeleteProfiles200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeleteProfiles200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeleteProfiles200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


