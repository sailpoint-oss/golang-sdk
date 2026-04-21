---
id: nerm-submit-roles200-response
title: SubmitRoles200Response
pagination_label: SubmitRoles200Response
sidebar_label: SubmitRoles200Response
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitRoles200Response', 'NERMSubmitRoles200Response'] 
slug: /tools/sdk/go/nerm/models/submit-roles200-response
tags: ['SDK', 'Software Development Kit', 'SubmitRoles200Response', 'NERMSubmitRoles200Response']
---

# SubmitRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]Role**](role) |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**JobStatus** | Pointer to [**JobJobStatus**](job-job-status) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubmitRoles200Response

`func NewSubmitRoles200Response() *SubmitRoles200Response`

NewSubmitRoles200Response instantiates a new SubmitRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitRoles200ResponseWithDefaults

`func NewSubmitRoles200ResponseWithDefaults() *SubmitRoles200Response`

NewSubmitRoles200ResponseWithDefaults instantiates a new SubmitRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *SubmitRoles200Response) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SubmitRoles200Response) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SubmitRoles200Response) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *SubmitRoles200Response) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetInfo

`func (o *SubmitRoles200Response) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SubmitRoles200Response) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SubmitRoles200Response) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SubmitRoles200Response) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobStatus

`func (o *SubmitRoles200Response) GetJobStatus() JobJobStatus`

GetJobStatus returns the JobStatus field if non-nil, zero value otherwise.

### GetJobStatusOk

`func (o *SubmitRoles200Response) GetJobStatusOk() (*JobJobStatus, bool)`

GetJobStatusOk returns a tuple with the JobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStatus

`func (o *SubmitRoles200Response) SetJobStatus(v JobJobStatus)`

SetJobStatus sets JobStatus field to given value.

### HasJobStatus

`func (o *SubmitRoles200Response) HasJobStatus() bool`

HasJobStatus returns a boolean if a field has been set.

### GetStatus

`func (o *SubmitRoles200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitRoles200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitRoles200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmitRoles200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


