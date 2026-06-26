---
id: nerm-get-job-status200-response-job-data-inner
title: GetJobStatus200ResponseJobDataInner
pagination_label: GetJobStatus200ResponseJobDataInner
sidebar_label: GetJobStatus200ResponseJobDataInner
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'GetJobStatus200ResponseJobDataInner', 'NERMGetJobStatus200ResponseJobDataInner'] 
slug: /tools/sdk/go/nerm/models/get-job-status200-response-job-data-inner
tags: ['SDK', 'Software Development Kit', 'GetJobStatus200ResponseJobDataInner', 'NERMGetJobStatus200ResponseJobDataInner']
---

# GetJobStatus200ResponseJobDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**ManagerId** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetJobStatus200ResponseJobDataInner

`func NewGetJobStatus200ResponseJobDataInner() *GetJobStatus200ResponseJobDataInner`

NewGetJobStatus200ResponseJobDataInner instantiates a new GetJobStatus200ResponseJobDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobStatus200ResponseJobDataInnerWithDefaults

`func NewGetJobStatus200ResponseJobDataInnerWithDefaults() *GetJobStatus200ResponseJobDataInner`

NewGetJobStatus200ResponseJobDataInnerWithDefaults instantiates a new GetJobStatus200ResponseJobDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetJobStatus200ResponseJobDataInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetJobStatus200ResponseJobDataInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetJobStatus200ResponseJobDataInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetJobStatus200ResponseJobDataInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetManagerId

`func (o *GetJobStatus200ResponseJobDataInner) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *GetJobStatus200ResponseJobDataInner) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *GetJobStatus200ResponseJobDataInner) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *GetJobStatus200ResponseJobDataInner) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### GetErrors

`func (o *GetJobStatus200ResponseJobDataInner) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetJobStatus200ResponseJobDataInner) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetJobStatus200ResponseJobDataInner) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetJobStatus200ResponseJobDataInner) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


