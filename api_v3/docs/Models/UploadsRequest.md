---
id: uploads-request
title: UploadsRequest
pagination_label: UploadsRequest
sidebar_label: UploadsRequest
sidebar_class_name: gosdk
keywords: ['go', 'golang', 'sdk', 'UploadsRequest'] 
slug: /tools/sdk/go/v3/models/uploads-request
tags: ['SDK', 'Software Development Kit', 'UploadsRequest']
---

# UploadsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** |  **string** | Unique id assigned to this job. | 
**Status** |  **string** | Status of the job. | 
**Type** |  **string** | Type of the job, either Backup or Draft. | 
**Tenant** |  Pointer to **string** | The name of the tenant performing the upload | [optional] 
**RequesterName** |  Pointer to **string** | The name of the requester. | [optional] 
**Created** |  **time.Time** | The time the job was started. | 
**Modified** |  **time.Time** | The time of the last update to the job. | 
**Name** |  Pointer to **string** | The name assigned to the upload file in the request body. | [optional] 
**UserCanDelete** |  Pointer to **bool** | Is the job a regular backup job, if so is the user allowed to delete the backup file. Since this is an upload job it remains as false. | [optional] [default to true]
**IsPartial** |  Pointer to **bool** | Is the job a regular backup job, if so is it partial. Since this is an upload job it remains as false. | [optional] [default to false]
**BackupType** |  Pointer to **string** | What kind of backup is this being treated as. | [optional] 
**HydrationStatus** |  Pointer to **string** | have the objects contained in the upload file been hydrated. | [optional] 

## Methods

### NewUploadsRequest

`func NewUploadsRequest(jobId string, status string, type_ string, created time.Time, modified time.Time, ) *UploadsRequest`

NewUploadsRequest instantiates a new UploadsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadsRequestWithDefaults

`func NewUploadsRequestWithDefaults() *UploadsRequest`

NewUploadsRequestWithDefaults instantiates a new UploadsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *UploadsRequest) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UploadsRequest) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UploadsRequest) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *UploadsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *UploadsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTenant

`func (o *UploadsRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *UploadsRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *UploadsRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *UploadsRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetRequesterName

`func (o *UploadsRequest) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *UploadsRequest) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *UploadsRequest) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *UploadsRequest) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### GetCreated

`func (o *UploadsRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UploadsRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UploadsRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *UploadsRequest) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *UploadsRequest) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *UploadsRequest) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *UploadsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UploadsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserCanDelete

`func (o *UploadsRequest) GetUserCanDelete() bool`

GetUserCanDelete returns the UserCanDelete field if non-nil, zero value otherwise.

### GetUserCanDeleteOk

`func (o *UploadsRequest) GetUserCanDeleteOk() (*bool, bool)`

GetUserCanDeleteOk returns a tuple with the UserCanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCanDelete

`func (o *UploadsRequest) SetUserCanDelete(v bool)`

SetUserCanDelete sets UserCanDelete field to given value.

### HasUserCanDelete

`func (o *UploadsRequest) HasUserCanDelete() bool`

HasUserCanDelete returns a boolean if a field has been set.

### GetIsPartial

`func (o *UploadsRequest) GetIsPartial() bool`

GetIsPartial returns the IsPartial field if non-nil, zero value otherwise.

### GetIsPartialOk

`func (o *UploadsRequest) GetIsPartialOk() (*bool, bool)`

GetIsPartialOk returns a tuple with the IsPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartial

`func (o *UploadsRequest) SetIsPartial(v bool)`

SetIsPartial sets IsPartial field to given value.

### HasIsPartial

`func (o *UploadsRequest) HasIsPartial() bool`

HasIsPartial returns a boolean if a field has been set.

### GetBackupType

`func (o *UploadsRequest) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *UploadsRequest) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *UploadsRequest) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *UploadsRequest) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetHydrationStatus

`func (o *UploadsRequest) GetHydrationStatus() string`

GetHydrationStatus returns the HydrationStatus field if non-nil, zero value otherwise.

### GetHydrationStatusOk

`func (o *UploadsRequest) GetHydrationStatusOk() (*string, bool)`

GetHydrationStatusOk returns a tuple with the HydrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydrationStatus

`func (o *UploadsRequest) SetHydrationStatus(v string)`

SetHydrationStatus sets HydrationStatus field to given value.

### HasHydrationStatus

`func (o *UploadsRequest) HasHydrationStatus() bool`

HasHydrationStatus returns a boolean if a field has been set.


[[Back to top]](#) 


