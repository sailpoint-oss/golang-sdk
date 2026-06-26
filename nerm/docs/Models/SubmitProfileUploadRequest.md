---
id: nerm-submit-profile-upload-request
title: SubmitProfileUploadRequest
pagination_label: SubmitProfileUploadRequest
sidebar_label: SubmitProfileUploadRequest
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SubmitProfileUploadRequest', 'NERMSubmitProfileUploadRequest'] 
slug: /tools/sdk/go/nerm/models/submit-profile-upload-request
tags: ['SDK', 'Software Development Kit', 'SubmitProfileUploadRequest', 'NERMSubmitProfileUploadRequest']
---

# SubmitProfileUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewSubmitProfileUploadRequest

`func NewSubmitProfileUploadRequest() *SubmitProfileUploadRequest`

NewSubmitProfileUploadRequest instantiates a new SubmitProfileUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitProfileUploadRequestWithDefaults

`func NewSubmitProfileUploadRequestWithDefaults() *SubmitProfileUploadRequest`

NewSubmitProfileUploadRequestWithDefaults instantiates a new SubmitProfileUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *SubmitProfileUploadRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SubmitProfileUploadRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SubmitProfileUploadRequest) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *SubmitProfileUploadRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.


