---
id: nerm-submit-user-managers200-response
title: SubmitUserManagers200Response
pagination_label: SubmitUserManagers200Response
sidebar_label: SubmitUserManagers200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserManagers200Response', 'NERMSubmitUserManagers200Response'] 
slug: /tools/sdk/go/nerm/models/submit-user-managers200-response
tags: ['SDK', 'Software Development Kit', 'SubmitUserManagers200Response', 'NERMSubmitUserManagers200Response']
---

# SubmitUserManagers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**UserManagers** | Pointer to [**[]UserManager**](user-manager) |  | [optional] 

## Methods

### NewSubmitUserManagers200Response

`func NewSubmitUserManagers200Response() *SubmitUserManagers200Response`

NewSubmitUserManagers200Response instantiates a new SubmitUserManagers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserManagers200ResponseWithDefaults

`func NewSubmitUserManagers200ResponseWithDefaults() *SubmitUserManagers200Response`

NewSubmitUserManagers200ResponseWithDefaults instantiates a new SubmitUserManagers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *SubmitUserManagers200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SubmitUserManagers200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SubmitUserManagers200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SubmitUserManagers200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *SubmitUserManagers200Response) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *SubmitUserManagers200Response) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *SubmitUserManagers200Response) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *SubmitUserManagers200Response) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *SubmitUserManagers200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitUserManagers200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitUserManagers200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmitUserManagers200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserManagers

`func (o *SubmitUserManagers200Response) GetUserManagers() []UserManager`

GetUserManagers returns the UserManagers field if non-nil, zero value otherwise.

### GetUserManagersOk

`func (o *SubmitUserManagers200Response) GetUserManagersOk() (*[]UserManager, bool)`

GetUserManagersOk returns a tuple with the UserManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagers

`func (o *SubmitUserManagers200Response) SetUserManagers(v []UserManager)`

SetUserManagers sets UserManagers field to given value.

### HasUserManagers

`func (o *SubmitUserManagers200Response) HasUserManagers() bool`

HasUserManagers returns a boolean if a field has been set.


