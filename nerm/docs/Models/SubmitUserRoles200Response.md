---
id: nerm-submit-user-roles200-response
title: SubmitUserRoles200Response
pagination_label: SubmitUserRoles200Response
sidebar_label: SubmitUserRoles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitUserRoles200Response', 'NERMSubmitUserRoles200Response'] 
slug: /tools/sdk/go/nerm/models/submit-user-roles200-response
tags: ['SDK', 'Software Development Kit', 'SubmitUserRoles200Response', 'NERMSubmitUserRoles200Response']
---

# SubmitUserRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRoles** | Pointer to [**[]UserRole**](user-role) |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubmitUserRoles200Response

`func NewSubmitUserRoles200Response() *SubmitUserRoles200Response`

NewSubmitUserRoles200Response instantiates a new SubmitUserRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitUserRoles200ResponseWithDefaults

`func NewSubmitUserRoles200ResponseWithDefaults() *SubmitUserRoles200Response`

NewSubmitUserRoles200ResponseWithDefaults instantiates a new SubmitUserRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRoles

`func (o *SubmitUserRoles200Response) GetUserRoles() []UserRole`

GetUserRoles returns the UserRoles field if non-nil, zero value otherwise.

### GetUserRolesOk

`func (o *SubmitUserRoles200Response) GetUserRolesOk() (*[]UserRole, bool)`

GetUserRolesOk returns a tuple with the UserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoles

`func (o *SubmitUserRoles200Response) SetUserRoles(v []UserRole)`

SetUserRoles sets UserRoles field to given value.

### HasUserRoles

`func (o *SubmitUserRoles200Response) HasUserRoles() bool`

HasUserRoles returns a boolean if a field has been set.

### GetInfo

`func (o *SubmitUserRoles200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SubmitUserRoles200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SubmitUserRoles200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SubmitUserRoles200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *SubmitUserRoles200Response) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *SubmitUserRoles200Response) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *SubmitUserRoles200Response) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *SubmitUserRoles200Response) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *SubmitUserRoles200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitUserRoles200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitUserRoles200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmitUserRoles200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


