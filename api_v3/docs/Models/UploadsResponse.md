---
id: uploads-response
title: UploadsResponse
pagination_label: UploadsResponse
sidebar_label: UploadsResponse
sidebar_class_name: gosdk
keywords: ['go', 'golang', 'sdk', 'UploadsResponse'] 
slug: /tools/sdk/go/v3/models/uploads-response
tags: ['SDK', 'Software Development Kit', 'UploadsResponse']
---

# UploadsResponse

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

### NewUploadsResponse

`func NewUploadsResponse(jobId string, status string, type_ string, created time.Time, modified time.Time, ) *UploadsResponse`

NewUploadsResponse instantiates a new UploadsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadsResponseWithDefaults

`func NewUploadsResponseWithDefaults() *UploadsResponse`

NewUploadsResponseWithDefaults instantiates a new UploadsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *UploadsResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UploadsResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UploadsResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *UploadsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *UploadsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadsResponse) SetType(v string)`

SetType sets Type field to given value.


### GetTenant

`func (o *UploadsResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *UploadsResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *UploadsResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *UploadsResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetRequesterName

`func (o *UploadsResponse) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *UploadsResponse) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *UploadsResponse) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *UploadsResponse) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### GetCreated

`func (o *UploadsResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UploadsResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UploadsResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *UploadsResponse) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *UploadsResponse) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *UploadsResponse) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetName

`func (o *UploadsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UploadsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserCanDelete

`func (o *UploadsResponse) GetUserCanDelete() bool`

GetUserCanDelete returns the UserCanDelete field if non-nil, zero value otherwise.

### GetUserCanDeleteOk

`func (o *UploadsResponse) GetUserCanDeleteOk() (*bool, bool)`

GetUserCanDeleteOk returns a tuple with the UserCanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCanDelete

`func (o *UploadsResponse) SetUserCanDelete(v bool)`

SetUserCanDelete sets UserCanDelete field to given value.

### HasUserCanDelete

`func (o *UploadsResponse) HasUserCanDelete() bool`

HasUserCanDelete returns a boolean if a field has been set.

### GetIsPartial

`func (o *UploadsResponse) GetIsPartial() bool`

GetIsPartial returns the IsPartial field if non-nil, zero value otherwise.

### GetIsPartialOk

`func (o *UploadsResponse) GetIsPartialOk() (*bool, bool)`

GetIsPartialOk returns a tuple with the IsPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartial

`func (o *UploadsResponse) SetIsPartial(v bool)`

SetIsPartial sets IsPartial field to given value.

### HasIsPartial

`func (o *UploadsResponse) HasIsPartial() bool`

HasIsPartial returns a boolean if a field has been set.

### GetBackupType

`func (o *UploadsResponse) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *UploadsResponse) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *UploadsResponse) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *UploadsResponse) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetHydrationStatus

`func (o *UploadsResponse) GetHydrationStatus() string`

GetHydrationStatus returns the HydrationStatus field if non-nil, zero value otherwise.

### GetHydrationStatusOk

`func (o *UploadsResponse) GetHydrationStatusOk() (*string, bool)`

GetHydrationStatusOk returns a tuple with the HydrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydrationStatus

`func (o *UploadsResponse) SetHydrationStatus(v string)`

SetHydrationStatus sets HydrationStatus field to given value.

### HasHydrationStatus

`func (o *UploadsResponse) HasHydrationStatus() bool`

HasHydrationStatus returns a boolean if a field has been set.


[[Back to top]](#) 


