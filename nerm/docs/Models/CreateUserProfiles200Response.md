---
id: nerm-create-user-profiles200-response
title: CreateUserProfiles200Response
pagination_label: CreateUserProfiles200Response
sidebar_label: CreateUserProfiles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'CreateUserProfiles200Response', 'NERMCreateUserProfiles200Response'] 
slug: /tools/sdk/go/nerm/models/create-user-profiles200-response
tags: ['SDK', 'Software Development Kit', 'CreateUserProfiles200Response', 'NERMCreateUserProfiles200Response']
---

# CreateUserProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserProfiles** | Pointer to [**[]UserProfile**](user-profile) |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateUserProfiles200Response

`func NewCreateUserProfiles200Response() *CreateUserProfiles200Response`

NewCreateUserProfiles200Response instantiates a new CreateUserProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserProfiles200ResponseWithDefaults

`func NewCreateUserProfiles200ResponseWithDefaults() *CreateUserProfiles200Response`

NewCreateUserProfiles200ResponseWithDefaults instantiates a new CreateUserProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserProfiles

`func (o *CreateUserProfiles200Response) GetUserProfiles() []UserProfile`

GetUserProfiles returns the UserProfiles field if non-nil, zero value otherwise.

### GetUserProfilesOk

`func (o *CreateUserProfiles200Response) GetUserProfilesOk() (*[]UserProfile, bool)`

GetUserProfilesOk returns a tuple with the UserProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfiles

`func (o *CreateUserProfiles200Response) SetUserProfiles(v []UserProfile)`

SetUserProfiles sets UserProfiles field to given value.

### HasUserProfiles

`func (o *CreateUserProfiles200Response) HasUserProfiles() bool`

HasUserProfiles returns a boolean if a field has been set.

### GetInfo

`func (o *CreateUserProfiles200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CreateUserProfiles200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CreateUserProfiles200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CreateUserProfiles200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *CreateUserProfiles200Response) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *CreateUserProfiles200Response) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *CreateUserProfiles200Response) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *CreateUserProfiles200Response) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *CreateUserProfiles200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUserProfiles200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUserProfiles200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateUserProfiles200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


