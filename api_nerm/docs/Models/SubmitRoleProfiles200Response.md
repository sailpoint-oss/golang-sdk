---
id: nerm-submit-role-profiles200-response
title: SubmitRoleProfiles200Response
pagination_label: SubmitRoleProfiles200Response
sidebar_label: SubmitRoleProfiles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitRoleProfiles200Response', 'NERMSubmitRoleProfiles200Response'] 
slug: /tools/sdk/go/nerm/models/submit-role-profiles200-response
tags: ['SDK', 'Software Development Kit', 'SubmitRoleProfiles200Response', 'NERMSubmitRoleProfiles200Response']
---

# SubmitRoleProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleProfiles** | Pointer to [**[]RoleProfile**](role-profile) |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubmitRoleProfiles200Response

`func NewSubmitRoleProfiles200Response() *SubmitRoleProfiles200Response`

NewSubmitRoleProfiles200Response instantiates a new SubmitRoleProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRoleProfiles200ResponseWithDefaults

`func NewSubmitRoleProfiles200ResponseWithDefaults() *SubmitRoleProfiles200Response`

NewSubmitRoleProfiles200ResponseWithDefaults instantiates a new SubmitRoleProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleProfiles

`func (o *SubmitRoleProfiles200Response) GetRoleProfiles() []RoleProfile`

GetRoleProfiles returns the RoleProfiles field if non-nil, zero value otherwise.

### GetRoleProfilesOk

`func (o *SubmitRoleProfiles200Response) GetRoleProfilesOk() (*[]RoleProfile, bool)`

GetRoleProfilesOk returns a tuple with the RoleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleProfiles

`func (o *SubmitRoleProfiles200Response) SetRoleProfiles(v []RoleProfile)`

SetRoleProfiles sets RoleProfiles field to given value.

### HasRoleProfiles

`func (o *SubmitRoleProfiles200Response) HasRoleProfiles() bool`

HasRoleProfiles returns a boolean if a field has been set.

### GetInfo

`func (o *SubmitRoleProfiles200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SubmitRoleProfiles200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SubmitRoleProfiles200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SubmitRoleProfiles200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *SubmitRoleProfiles200Response) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *SubmitRoleProfiles200Response) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *SubmitRoleProfiles200Response) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *SubmitRoleProfiles200Response) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *SubmitRoleProfiles200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitRoleProfiles200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitRoleProfiles200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmitRoleProfiles200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


