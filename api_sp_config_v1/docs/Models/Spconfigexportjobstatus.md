---
id: v1-spconfigexportjobstatus
title: Spconfigexportjobstatus
pagination_label: Spconfigexportjobstatus
sidebar_label: Spconfigexportjobstatus
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Spconfigexportjobstatus', 'V1Spconfigexportjobstatus'] 
slug: /tools/sdk/go/spconfigv1/models/spconfigexportjobstatus
tags: ['SDK', 'Software Development Kit', 'Spconfigexportjobstatus', 'V1Spconfigexportjobstatus']
---

# Spconfigexportjobstatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** | Unique id assigned to this job. | 
**Status** | **string** | Status of the job. | 
**Type** | **string** | Type of the job, either export or import. | 
**Expiration** | Pointer to **time.Time** | The time until which the artifacts will be available for download. | [optional] 
**Created** | **time.Time** | The time the job was started. | 
**Modified** | **time.Time** | The time of the last update to the job. | 
**Description** | Pointer to **string** | Optional user defined description/name for export job. | [optional] 
**Completed** | Pointer to **time.Time** | The time the job was completed. | [optional] 

## Methods

### NewSpconfigexportjobstatus

`func NewSpconfigexportjobstatus(jobId string, status string, type_ string, created time.Time, modified time.Time, ) *Spconfigexportjobstatus`

NewSpconfigexportjobstatus instantiates a new Spconfigexportjobstatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpconfigexportjobstatusWithDefaults

`func NewSpconfigexportjobstatusWithDefaults() *Spconfigexportjobstatus`

NewSpconfigexportjobstatusWithDefaults instantiates a new Spconfigexportjobstatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *Spconfigexportjobstatus) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Spconfigexportjobstatus) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Spconfigexportjobstatus) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *Spconfigexportjobstatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Spconfigexportjobstatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Spconfigexportjobstatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *Spconfigexportjobstatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Spconfigexportjobstatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Spconfigexportjobstatus) SetType(v string)`

SetType sets Type field to given value.


### GetExpiration

`func (o *Spconfigexportjobstatus) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Spconfigexportjobstatus) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Spconfigexportjobstatus) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Spconfigexportjobstatus) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetCreated

`func (o *Spconfigexportjobstatus) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Spconfigexportjobstatus) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Spconfigexportjobstatus) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Spconfigexportjobstatus) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Spconfigexportjobstatus) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Spconfigexportjobstatus) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetDescription

`func (o *Spconfigexportjobstatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Spconfigexportjobstatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Spconfigexportjobstatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Spconfigexportjobstatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCompleted

`func (o *Spconfigexportjobstatus) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *Spconfigexportjobstatus) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *Spconfigexportjobstatus) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *Spconfigexportjobstatus) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


